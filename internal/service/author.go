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
	"github.com/ts1106/favomic-api/ent/author"
	api "github.com/ts1106/favomic-api/gen/api/v1"
	apiconnect "github.com/ts1106/favomic-api/gen/api/v1/v1connect"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthorService struct {
	client *ent.Client
	apiconnect.UnimplementedAuthorServiceHandler
}

func NewAuthorService(client *ent.Client) *AuthorService {
	return &AuthorService{
		client: client,
	}
}

// Create implements AuthorServiceServer.Create
func (svc *AuthorService) Create(ctx context.Context, req *connect.Request[api.CreateAuthorRequest]) (*connect.Response[api.Author], error) {
	author := req.Msg.GetAuthor()
	m, err := svc.createBuilder(author)
	if err != nil {
		return nil, err
	}
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoAuthor(res)
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

// Get implements AuthorServiceServer.Get
func (svc *AuthorService) Get(ctx context.Context, req *connect.Request[api.GetAuthorRequest]) (*connect.Response[api.Author], error) {
	var (
		err error
		get *ent.Author
	)
	var id uuid.UUID
	if err := (&id).UnmarshalBinary(req.Msg.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	switch req.Msg.GetView() {
	case api.GetAuthorRequest_BASIC:
		get, err = svc.client.Author.Get(ctx, id)
	case api.GetAuthorRequest_WITH_EDGES:
		get, err = svc.client.Author.Query().
			Where(author.ID(id)).
			WithComics().
			Only(ctx)
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid argument: unknown view")
	}
	switch {
	case err == nil:
		proto, err := toProtoAuthor(get)
		return connect.NewResponse(proto), err
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Update implements AuthorServiceServer.Update
func (svc *AuthorService) Update(ctx context.Context, req *connect.Request[api.UpdateAuthorRequest]) (*connect.Response[api.Author], error) {
	author := req.Msg.GetAuthor()
	var authorID uuid.UUID
	if err := (&authorID).UnmarshalBinary(author.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	m := svc.client.Author.UpdateOneID(authorID)
	authorName := author.GetName()
	m.SetName(authorName)
	for _, item := range author.GetComics() {
		var comics uuid.UUID
		if err := (&comics).UnmarshalBinary(item.GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
		}
		m.AddComicIDs(comics)
	}

	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoAuthor(res)
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

// Delete implements AuthorServiceServer.Delete
func (svc *AuthorService) Delete(ctx context.Context, req *connect.Request[api.DeleteAuthorRequest]) (*connect.Response[emptypb.Empty], error) {
	var err error
	var id uuid.UUID
	if err := (&id).UnmarshalBinary(req.Msg.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	err = svc.client.Author.DeleteOneID(id).Exec(ctx)
	switch {
	case err == nil:
		return connect.NewResponse(&emptypb.Empty{}), nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// List implements AuthorServiceServer.List
func (svc *AuthorService) List(ctx context.Context, req *connect.Request[api.ListAuthorRequest]) (*connect.Response[api.ListAuthorResponse], error) {
	var (
		err      error
		entList  []*ent.Author
		pageSize int
	)
	pageSize = int(req.Msg.GetPageSize())
	switch {
	case pageSize < 0:
		return nil, status.Errorf(codes.InvalidArgument, "page size cannot be less than zero")
	case pageSize == 0 || pageSize > entproto.MaxPageSize:
		pageSize = entproto.MaxPageSize
	}
	listQuery := svc.client.Author.Query().
		Order(ent.Desc(author.FieldID)).
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
			Where(author.IDLTE(pageToken))
	}
	switch req.Msg.GetView() {
	case api.ListAuthorRequest_BASIC:
		entList, err = listQuery.All(ctx)
	case api.ListAuthorRequest_WITH_EDGES:
		entList, err = listQuery.
			WithComics().
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
		protoList, err := toProtoAuthorList(entList)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return connect.NewResponse(&api.ListAuthorResponse{
			AuthorList:    protoList,
			NextPageToken: nextPageToken,
		}), nil
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// BatchCreate implements AuthorServiceServer.BatchCreate
func (svc *AuthorService) BatchCreate(ctx context.Context, req *connect.Request[api.BatchCreateAuthorsRequest]) (*connect.Response[api.BatchCreateAuthorsResponse], error) {
	requests := req.Msg.GetRequests()
	if len(requests) > entproto.MaxBatchCreateSize {
		return nil, status.Errorf(codes.InvalidArgument, "batch size cannot be greater than %d", entproto.MaxBatchCreateSize)
	}
	bulk := make([]*ent.AuthorCreate, len(requests))
	for i, req := range requests {
		author := req.GetAuthor()
		var err error
		bulk[i], err = svc.createBuilder(author)
		if err != nil {
			return nil, err
		}
	}
	res, err := svc.client.Author.CreateBulk(bulk...).Save(ctx)
	switch {
	case err == nil:
		protoList, err := toProtoAuthorList(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return connect.NewResponse(&api.BatchCreateAuthorsResponse{
			Authors: protoList,
		}), nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

func (svc *AuthorService) Upsert(ctx context.Context, req *connect.Request[api.UpsertAuthorRequest]) (*connect.Response[api.Author], error) {
	a := req.Msg.GetAuthor()
	m, err := svc.createBuilder(a)
	if err != nil {
		return nil, err
	}
	err = m.OnConflict().UpdateName().Exec(ctx)
	switch {
	case err == nil:
		res, err := svc.client.Author.Query().Where(author.Name(a.GetName())).Only(ctx)
		switch {
		case err == nil:
			proto, err := toProtoAuthor(res)
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

func (svc *AuthorService) Search(ctx context.Context, req *connect.Request[api.SearchAuthorRequest]) (*connect.Response[api.Author], error) {
	query := svc.client.Author.Query().
		Where(author.Name(req.Msg.GetName()))

	switch req.Msg.GetView() {
	case api.SearchAuthorRequest_BASIC:
	case api.SearchAuthorRequest_WITH_EDGES:
		query.WithComics()
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid argument: unknown view")
	}

	get, err := query.Only(ctx)

	switch {
	case err == nil:
		proto, err := toProtoAuthor(get)
		return connect.NewResponse(proto), err
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}
}

func (svc *AuthorService) createBuilder(author *api.Author) (*ent.AuthorCreate, error) {
	m := svc.client.Author.Create()
	authorName := author.GetName()
	m.SetName(authorName)
	for _, item := range author.GetComics() {
		var comics uuid.UUID
		if err := (&comics).UnmarshalBinary(item.GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
		}
		m.AddComicIDs(comics)
	}
	return m, nil
}

func toProtoAuthor(e *ent.Author) (*api.Author, error) {
	v := &api.Author{}
	id, err := e.ID.MarshalBinary()
	if err != nil {
		return nil, err
	}
	v.Id = id
	name := e.Name
	v.Name = name
	for _, edg := range e.Edges.Comics {
		id, err := edg.ID.MarshalBinary()
		if err != nil {
			return nil, err
		}
		v.Comics = append(v.Comics, &api.Comic{
			Id:    id,
			Title: edg.Title,
		})
	}
	return v, nil
}

func toProtoAuthorList(e []*ent.Author) ([]*api.Author, error) {
	var pbList []*api.Author
	for _, entEntity := range e {
		pbEntity, err := toProtoAuthor(entEntity)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		pbList = append(pbList, pbEntity)
	}
	return pbList, nil
}
