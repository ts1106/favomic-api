package service

import (
	"context"
	"encoding/base64"
	"fmt"

	"entgo.io/contrib/entproto"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"github.com/ts1106/favomic-api/ent"
	"github.com/ts1106/favomic-api/ent/comic"
	api "github.com/ts1106/favomic-api/gen/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ComicService implements ComicServiceServer
type ComicService struct {
	client *ent.Client
}

// NewComicService returns a new ComicService
func NewComicService(client *ent.Client) *ComicService {
	return &ComicService{
		client: client,
	}
}

// toProtoComic transforms the ent type to the pb type
func toProtoComic(e *ent.Comic) (*api.Comic, error) {
	v := &api.Comic{}
	author, err := e.AuthorID.MarshalBinary()
	if err != nil {
		return nil, err
	}
	v.AuthorId = author
	id, err := e.ID.MarshalBinary()
	if err != nil {
		return nil, err
	}
	v.Id = id
	magazine, err := e.MagazineID.MarshalBinary()
	if err != nil {
		return nil, err
	}
	v.MagazineId = magazine
	title := e.Title
	v.Title = title
	if edg := e.Edges.Author; edg != nil {
		v.Author = &api.Author{
			Name: edg.Name,
		}
	}
	for _, edg := range e.Edges.Episodes {
		id, err := edg.ID.MarshalBinary()
		if err != nil {
			return nil, err
		}
		v.Episodes = append(v.Episodes, &api.Episode{
			Id:        id,
			Title:     edg.Title,
			Url:       edg.URL,
			Thumbnail: edg.Thumbnail,
			UpdatedAt: timestamppb.New(edg.UpdatedAt),
		})
	}
	if edg := e.Edges.Magazine; edg != nil {
		v.Magazine = &api.Magazine{
			Name: edg.Name,
		}
	}
	for _, edg := range e.Edges.Tags {
		id, err := edg.ID.MarshalBinary()
		if err != nil {
			return nil, err
		}
		v.Tags = append(v.Tags, &api.Tag{
			Id:   id,
			Name: edg.Name,
		})
	}
	return v, nil
}

// toProtoComicList transforms a list of ent type to a list of pb type
func toProtoComicList(e []*ent.Comic) ([]*api.Comic, error) {
	var pbList []*api.Comic
	for _, entEntity := range e {
		pbEntity, err := toProtoComic(entEntity)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		pbList = append(pbList, pbEntity)
	}
	return pbList, nil
}

// Create implements ComicServiceServer.Create
func (svc *ComicService) Create(ctx context.Context, req *connect.Request[api.CreateComicRequest]) (*connect.Response[api.Comic], error) {
	comic := req.Msg.GetComic()
	m, err := svc.createBuilder(comic)
	if err != nil {
		return nil, err
	}
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoComic(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return connect.NewResponse(proto), nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Get implements ComicServiceServer.Get
func (svc *ComicService) Get(ctx context.Context, req *connect.Request[api.GetComicRequest]) (*connect.Response[api.Comic], error) {
	var (
		err error
		get *ent.Comic
	)
	var id uuid.UUID
	if err := (&id).UnmarshalBinary(req.Msg.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	switch req.Msg.GetView() {
	case api.GetComicRequest_BASIC:
		get, err = svc.client.Comic.Get(ctx, id)
	case api.GetComicRequest_WITH_EDGES:
		get, err = svc.client.Comic.Query().
			Where(comic.ID(id)).
			WithAuthor().
			WithEpisodes().
			WithMagazine().
			WithTags().
			Only(ctx)
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid argument: unknown view")
	}
	switch {
	case err == nil:
		proto, err := toProtoComic(get)
		return connect.NewResponse(proto), err
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Update implements ComicServiceServer.Update
func (svc *ComicService) Update(ctx context.Context, req *connect.Request[api.UpdateComicRequest]) (*connect.Response[api.Comic], error) {
	comic := req.Msg.GetComic()
	var comicID uuid.UUID
	if err := (&comicID).UnmarshalBinary(comic.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	m := svc.client.Comic.UpdateOneID(comicID)
	var comicAuthorID uuid.UUID
	if err := (&comicAuthorID).UnmarshalBinary(comic.GetAuthorId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	m.SetAuthorID(comicAuthorID)
	var comicMagazineID uuid.UUID
	if err := (&comicMagazineID).UnmarshalBinary(comic.GetMagazineId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	m.SetMagazineID(comicMagazineID)
	comicTitle := comic.GetTitle()
	m.SetTitle(comicTitle)
	if comic.GetAuthor() != nil {
		var comicAuthor uuid.UUID
		if err := (&comicAuthor).UnmarshalBinary(comic.GetAuthor().GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
		}
		m.SetAuthorID(comicAuthor)
	}
	for _, item := range comic.GetEpisodes() {
		var episodes uuid.UUID
		if err := (&episodes).UnmarshalBinary(item.GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
		}
		m.AddEpisodeIDs(episodes)
	}
	if comic.GetMagazine() != nil {
		var comicMagazine uuid.UUID
		if err := (&comicMagazine).UnmarshalBinary(comic.GetMagazine().GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
		}
		m.SetMagazineID(comicMagazine)
	}
	for _, item := range comic.GetTags() {
		var tags uuid.UUID
		if err := (&tags).UnmarshalBinary(item.GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
		}
		m.AddTagIDs(tags)
	}

	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoComic(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return connect.NewResponse(proto), nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Delete implements ComicServiceServer.Delete
func (svc *ComicService) Delete(ctx context.Context, req *connect.Request[api.DeleteComicRequest]) (*connect.Response[emptypb.Empty], error) {
	var err error
	var id uuid.UUID
	if err := (&id).UnmarshalBinary(req.Msg.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	err = svc.client.Comic.DeleteOneID(id).Exec(ctx)
	switch {
	case err == nil:
		return connect.NewResponse(&emptypb.Empty{}), nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// List implements ComicServiceServer.List
func (svc *ComicService) List(ctx context.Context, req *connect.Request[api.ListComicRequest]) (*connect.Response[api.ListComicResponse], error) {
	var (
		err      error
		entList  []*ent.Comic
		pageSize int
	)
	pageSize = int(req.Msg.GetPageSize())
	switch {
	case pageSize < 0:
		return nil, status.Errorf(codes.InvalidArgument, "page size cannot be less than zero")
	case pageSize == 0 || pageSize > entproto.MaxPageSize:
		pageSize = entproto.MaxPageSize
	}
	listQuery := svc.client.Comic.Query().
		Order(ent.Desc(comic.FieldID)).
		Limit(pageSize + 1)
	if req.Msg.GetPageToken() != "" {
		bytes, err := base64.StdEncoding.DecodeString(req.Msg.PageToken)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "page token is invalid")
		}
		pageToken, err := uuid.ParseBytes(bytes)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "page token is invalid")
		}
		listQuery = listQuery.
			Where(comic.IDLTE(pageToken))
	}
	switch req.Msg.GetView() {
	case api.ListComicRequest_BASIC:
		entList, err = listQuery.All(ctx)
	case api.ListComicRequest_WITH_EDGES:
		entList, err = listQuery.
			WithAuthor().
			WithEpisodes().
			WithMagazine().
			WithTags().
			All(ctx)
	}
	switch {
	case err == nil:
		var nextPageToken string
		if len(entList) == pageSize+1 {
			nextPageToken = base64.StdEncoding.EncodeToString(
				[]byte(fmt.Sprintf("%v", entList[len(entList)-1].ID)))
			entList = entList[:len(entList)-1]
		}
		protoList, err := toProtoComicList(entList)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return connect.NewResponse(&api.ListComicResponse{
			ComicList:     protoList,
			NextPageToken: nextPageToken,
		}), nil
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// BatchCreate implements ComicServiceServer.BatchCreate
func (svc *ComicService) BatchCreate(ctx context.Context, req *connect.Request[api.BatchCreateComicsRequest]) (*connect.Response[api.BatchCreateComicsResponse], error) {
	requests := req.Msg.GetRequests()
	if len(requests) > entproto.MaxBatchCreateSize {
		return nil, status.Errorf(codes.InvalidArgument, "batch size cannot be greater than %d", entproto.MaxBatchCreateSize)
	}
	bulk := make([]*ent.ComicCreate, len(requests))
	for i, req := range requests {
		comic := req.GetComic()
		var err error
		bulk[i], err = svc.createBuilder(comic)
		if err != nil {
			return nil, err
		}
	}
	res, err := svc.client.Comic.CreateBulk(bulk...).Save(ctx)
	switch {
	case err == nil:
		protoList, err := toProtoComicList(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return connect.NewResponse(&api.BatchCreateComicsResponse{
			Comics: protoList,
		}), nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

func (svc *ComicService) Upsert(ctx context.Context, req *connect.Request[api.UpsertComicRequest]) (*connect.Response[api.Comic], error) {
	c := req.Msg.GetComic()
	m, err := svc.createBuilder(c)
	if err != nil {
		return nil, err
	}
	err = m.OnConflict().UpdateTitle().Exec(ctx)
	switch {
	case err == nil:
		res, err := svc.client.Comic.Query().Where(comic.Title(c.GetTitle())).Only(ctx)
		switch {
		case err == nil:
			proto, err := toProtoComic(res)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "internal error: %s", err)
			}
			return connect.NewResponse(proto), nil
		case sqlgraph.IsUniqueConstraintError(err):
			return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
		case ent.IsConstraintError(err):
			return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
		default:
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

func (svc *ComicService) createBuilder(comic *api.Comic) (*ent.ComicCreate, error) {
	m := svc.client.Comic.Create()
	var comicAuthorID uuid.UUID
	if err := (&comicAuthorID).UnmarshalBinary(comic.GetAuthorId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	m.SetAuthorID(comicAuthorID)
	var comicMagazineID uuid.UUID
	if err := (&comicMagazineID).UnmarshalBinary(comic.GetMagazineId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	m.SetMagazineID(comicMagazineID)
	comicTitle := comic.GetTitle()
	m.SetTitle(comicTitle)
	if comic.GetAuthor() != nil {
		var comicAuthor uuid.UUID
		if err := (&comicAuthor).UnmarshalBinary(comic.GetAuthor().GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
		}
		m.SetAuthorID(comicAuthor)
	}
	for _, item := range comic.GetEpisodes() {
		var episodes uuid.UUID
		if err := (&episodes).UnmarshalBinary(item.GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
		}
		m.AddEpisodeIDs(episodes)
	}
	if comic.GetMagazine() != nil {
		var comicMagazine uuid.UUID
		if err := (&comicMagazine).UnmarshalBinary(comic.GetMagazine().GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
		}
		m.SetMagazineID(comicMagazine)
	}
	for _, item := range comic.GetTags() {
		var tags uuid.UUID
		if err := (&tags).UnmarshalBinary(item.GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
		}
		m.AddTagIDs(tags)
	}
	return m, nil
}
