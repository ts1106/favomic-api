package service

import (
	"context"
	"encoding/base64"
	"fmt"

	"entgo.io/contrib/entproto"
	"entgo.io/contrib/entproto/runtime"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"github.com/ts1106/favomic-api/ent"
	"github.com/ts1106/favomic-api/ent/episode"
	api "github.com/ts1106/favomic-api/gen/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// EpisodeService implements EpisodeServiceServer
type EpisodeService struct {
	client *ent.Client
}

// NewEpisodeService returns a new EpisodeService
func NewEpisodeService(client *ent.Client) *EpisodeService {
	return &EpisodeService{
		client: client,
	}
}

// toProtoEpisode transforms the ent type to the pb type
func toProtoEpisode(e *ent.Episode) (*api.Episode, error) {
	v := &api.Episode{}
	comic, err := e.ComicID.MarshalBinary()
	if err != nil {
		return nil, err
	}
	v.ComicId = comic
	id, err := e.ID.MarshalBinary()
	if err != nil {
		return nil, err
	}
	v.Id = id
	_Thumbnail := e.Thumbnail
	v.Thumbnail = _Thumbnail
	title := e.Title
	v.Title = title
	updated_at := timestamppb.New(e.UpdatedAt)
	v.UpdatedAt = updated_at
	url := e.URL
	v.Url = url
	if edg := e.Edges.Comic; edg != nil {
		v.Comic = &api.Comic{
			Title: edg.Title,
		}
	}
	return v, nil
}

// toProtoEpisodeList transforms a list of ent type to a list of pb type
func toProtoEpisodeList(e []*ent.Episode) ([]*api.Episode, error) {
	var pbList []*api.Episode
	for _, entEntity := range e {
		pbEntity, err := toProtoEpisode(entEntity)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		pbList = append(pbList, pbEntity)
	}
	return pbList, nil
}

// Create implements EpisodeServiceServer.Create
func (svc *EpisodeService) Create(ctx context.Context, req *connect.Request[api.CreateEpisodeRequest]) (*connect.Response[api.Episode], error) {
	episode := req.Msg.GetEpisode()
	m, err := svc.createBuilder(episode)
	if err != nil {
		return nil, err
	}
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoEpisode(res)
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

// Get implements EpisodeServiceServer.Get
func (svc *EpisodeService) Get(ctx context.Context, req *connect.Request[api.GetEpisodeRequest]) (*connect.Response[api.Episode], error) {
	var (
		err error
		get *ent.Episode
	)
	var id uuid.UUID
	if err := (&id).UnmarshalBinary(req.Msg.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	switch req.Msg.GetView() {
	case api.GetEpisodeRequest_BASIC:
		get, err = svc.client.Episode.Get(ctx, id)
	case api.GetEpisodeRequest_WITH_EDGES:
		get, err = svc.client.Episode.Query().
			Where(episode.ID(id)).
			WithComic().
			Only(ctx)
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid argument: unknown view")
	}
	switch {
	case err == nil:
		proto, err := toProtoEpisode(get)
		return connect.NewResponse(proto), err
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Update implements EpisodeServiceServer.Update
func (svc *EpisodeService) Update(ctx context.Context, req *connect.Request[api.UpdateEpisodeRequest]) (*connect.Response[api.Episode], error) {
	episode := req.Msg.GetEpisode()
	var episodeID uuid.UUID
	if err := (&episodeID).UnmarshalBinary(episode.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	m := svc.client.Episode.UpdateOneID(episodeID)
	var episodeComicID uuid.UUID
	if err := (&episodeComicID).UnmarshalBinary(episode.GetComicId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	m.SetComicID(episodeComicID)
	episodeThumbnail := episode.GetThumbnail()
	m.SetThumbnail(episodeThumbnail)
	episodeTitle := episode.GetTitle()
	m.SetTitle(episodeTitle)
	episodeUpdatedAt := runtime.ExtractTime(episode.GetUpdatedAt())
	m.SetUpdatedAt(episodeUpdatedAt)
	episodeURL := episode.GetUrl()
	m.SetURL(episodeURL)
	if episode.GetComic() != nil {
		var episodeComic uuid.UUID
		if err := (&episodeComic).UnmarshalBinary(episode.GetComic().GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
		}
		m.SetComicID(episodeComic)
	}

	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoEpisode(res)
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

// Delete implements EpisodeServiceServer.Delete
func (svc *EpisodeService) Delete(ctx context.Context, req *connect.Request[api.DeleteEpisodeRequest]) (*connect.Response[emptypb.Empty], error) {
	var err error
	var id uuid.UUID
	if err := (&id).UnmarshalBinary(req.Msg.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	err = svc.client.Episode.DeleteOneID(id).Exec(ctx)
	switch {
	case err == nil:
		return connect.NewResponse(&emptypb.Empty{}), nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// List implements EpisodeServiceServer.List
func (svc *EpisodeService) List(ctx context.Context, req *connect.Request[api.ListEpisodeRequest]) (*connect.Response[api.ListEpisodeResponse], error) {
	var (
		err      error
		entList  []*ent.Episode
		pageSize int
	)
	pageSize = int(req.Msg.GetPageSize())
	switch {
	case pageSize < 0:
		return nil, status.Errorf(codes.InvalidArgument, "page size cannot be less than zero")
	case pageSize == 0 || pageSize > entproto.MaxPageSize:
		pageSize = entproto.MaxPageSize
	}
	listQuery := svc.client.Episode.Query().
		Order(ent.Desc(episode.FieldID)).
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
			Where(episode.IDLTE(pageToken))
	}
	switch req.Msg.GetView() {
	case api.ListEpisodeRequest_BASIC:
		entList, err = listQuery.All(ctx)
	case api.ListEpisodeRequest_WITH_EDGES:
		entList, err = listQuery.
			WithComic().
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
		protoList, err := toProtoEpisodeList(entList)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return connect.NewResponse(&api.ListEpisodeResponse{
			EpisodeList:   protoList,
			NextPageToken: nextPageToken,
		}), nil
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// BatchCreate implements EpisodeServiceServer.BatchCreate
func (svc *EpisodeService) BatchCreate(ctx context.Context, req *connect.Request[api.BatchCreateEpisodesRequest]) (*connect.Response[api.BatchCreateEpisodesResponse], error) {
	requests := req.Msg.GetRequests()
	if len(requests) > entproto.MaxBatchCreateSize {
		return nil, status.Errorf(codes.InvalidArgument, "batch size cannot be greater than %d", entproto.MaxBatchCreateSize)
	}
	bulk := make([]*ent.EpisodeCreate, len(requests))
	for i, req := range requests {
		episode := req.GetEpisode()
		var err error
		bulk[i], err = svc.createBuilder(episode)
		if err != nil {
			return nil, err
		}
	}
	res, err := svc.client.Episode.CreateBulk(bulk...).Save(ctx)
	switch {
	case err == nil:
		protoList, err := toProtoEpisodeList(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return connect.NewResponse(&api.BatchCreateEpisodesResponse{
			Episodes: protoList,
		}), nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

func (svc *EpisodeService) Upsert(ctx context.Context, req *connect.Request[api.UpsertEpisodeRequest]) (*connect.Response[api.Episode], error) {
	e := req.Msg.GetEpisode()
	m, err := svc.createBuilder(e)
	if err != nil {
		return nil, err
	}
	err = m.OnConflict().UpdateURL().UpdateTitle().UpdateThumbnail().UpdateUpdatedAt().Exec(ctx)
	switch {
	case err == nil:
		res, err := svc.client.Episode.Query().Where(episode.URL(e.GetUrl())).Only(ctx)
		switch {
		case err == nil:
			proto, err := toProtoEpisode(res)
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

func (svc *EpisodeService) createBuilder(episode *api.Episode) (*ent.EpisodeCreate, error) {
	m := svc.client.Episode.Create()
	var episodeComicID uuid.UUID
	if err := (&episodeComicID).UnmarshalBinary(episode.GetComicId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	m.SetComicID(episodeComicID)
	episodeThumbnail := episode.GetThumbnail()
	m.SetThumbnail(episodeThumbnail)
	episodeTitle := episode.GetTitle()
	m.SetTitle(episodeTitle)
	episodeUpdatedAt := runtime.ExtractTime(episode.GetUpdatedAt())
	m.SetUpdatedAt(episodeUpdatedAt)
	episodeURL := episode.GetUrl()
	m.SetURL(episodeURL)
	if episode.GetComic() != nil {
		var episodeComic uuid.UUID
		if err := (&episodeComic).UnmarshalBinary(episode.GetComic().GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
		}
		m.SetComicID(episodeComic)
	}
	return m, nil
}
