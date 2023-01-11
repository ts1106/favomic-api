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
	"github.com/ts1106/favomic-api/ent/magazine"
	api "github.com/ts1106/favomic-api/gen/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// MagazineService implements MagazineServiceServer
type MagazineService struct {
	client *ent.Client
}

// NewMagazineService returns a new MagazineService
func NewMagazineService(client *ent.Client) *MagazineService {
	return &MagazineService{
		client: client,
	}
}

// toProtoMagazine transforms the ent type to the pb type
func toProtoMagazine(e *ent.Magazine) (*api.Magazine, error) {
	v := &api.Magazine{}
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
			Author: &api.Author{
				Name: edg.Edges.Author.Name,
			},
			Magazine: &api.Magazine{
				Name: edg.Edges.Magazine.Name,
			},
		})
	}
	return v, nil
}

// toProtoMagazineList transforms a list of ent type to a list of pb type
func toProtoMagazineList(e []*ent.Magazine) ([]*api.Magazine, error) {
	var pbList []*api.Magazine
	for _, entEntity := range e {
		pbEntity, err := toProtoMagazine(entEntity)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		pbList = append(pbList, pbEntity)
	}
	return pbList, nil
}

// Create implements MagazineServiceServer.Create
func (svc *MagazineService) Create(ctx context.Context, req *connect.Request[api.CreateMagazineRequest]) (*connect.Response[api.Magazine], error) {
	magazine := req.Msg.GetMagazine()
	m, err := svc.createBuilder(magazine)
	if err != nil {
		return nil, err
	}
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoMagazine(res)
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

// Get implements MagazineServiceServer.Get
func (svc *MagazineService) Get(ctx context.Context, req *connect.Request[api.GetMagazineRequest]) (*connect.Response[api.Magazine], error) {
	var (
		err error
		get *ent.Magazine
	)
	var id uuid.UUID
	if err := (&id).UnmarshalBinary(req.Msg.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	switch req.Msg.GetView() {
	case api.GetMagazineRequest_BASIC:
		get, err = svc.client.Magazine.Get(ctx, id)
	case api.GetMagazineRequest_WITH_EDGES:
		get, err = svc.client.Magazine.Query().
			Where(magazine.ID(id)).
			WithComics(func(query *ent.ComicQuery) {
				query.WithAuthor().WithMagazine()
			}).
			Only(ctx)
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid argument: unknown view")
	}
	switch {
	case err == nil:
		proto, err := toProtoMagazine(get)
		return connect.NewResponse(proto), err
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Update implements MagazineServiceServer.Update
func (svc *MagazineService) Update(ctx context.Context, req *connect.Request[api.UpdateMagazineRequest]) (*connect.Response[api.Magazine], error) {
	magazine := req.Msg.GetMagazine()
	var magazineID uuid.UUID
	if err := (&magazineID).UnmarshalBinary(magazine.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	m := svc.client.Magazine.UpdateOneID(magazineID)
	magazineName := magazine.GetName()
	m.SetName(magazineName)
	for _, item := range magazine.GetComics() {
		var comics uuid.UUID
		if err := (&comics).UnmarshalBinary(item.GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
		}
		m.AddComicIDs(comics)
	}

	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoMagazine(res)
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

// Delete implements MagazineServiceServer.Delete
func (svc *MagazineService) Delete(ctx context.Context, req *connect.Request[api.DeleteMagazineRequest]) (*connect.Response[emptypb.Empty], error) {
	var err error
	var id uuid.UUID
	if err := (&id).UnmarshalBinary(req.Msg.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	err = svc.client.Magazine.DeleteOneID(id).Exec(ctx)
	switch {
	case err == nil:
		return connect.NewResponse(&emptypb.Empty{}), nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// List implements MagazineServiceServer.List
func (svc *MagazineService) List(ctx context.Context, req *connect.Request[api.ListMagazineRequest]) (*connect.Response[api.ListMagazineResponse], error) {
	var (
		err      error
		entList  []*ent.Magazine
		pageSize int
	)
	pageSize = int(req.Msg.GetPageSize())
	switch {
	case pageSize < 0:
		return nil, status.Errorf(codes.InvalidArgument, "page size cannot be less than zero")
	case pageSize == 0 || pageSize > entproto.MaxPageSize:
		pageSize = entproto.MaxPageSize
	}
	listQuery := svc.client.Magazine.Query().
		Order(ent.Desc(magazine.FieldID)).
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
			Where(magazine.IDLTE(pageToken))
	}
	switch req.Msg.GetView() {
	case api.ListMagazineRequest_BASIC:
		entList, err = listQuery.All(ctx)
	case api.ListMagazineRequest_WITH_EDGES:
		entList, err = listQuery.
			WithComics(func(query *ent.ComicQuery) {
				query.WithAuthor().WithMagazine()
			}).
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
		protoList, err := toProtoMagazineList(entList)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return connect.NewResponse(&api.ListMagazineResponse{
			MagazineList:  protoList,
			NextPageToken: nextPageToken,
		}), nil
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// BatchCreate implements MagazineServiceServer.BatchCreate
func (svc *MagazineService) BatchCreate(ctx context.Context, req *connect.Request[api.BatchCreateMagazinesRequest]) (*connect.Response[api.BatchCreateMagazinesResponse], error) {
	requests := req.Msg.GetRequests()
	if len(requests) > entproto.MaxBatchCreateSize {
		return nil, status.Errorf(codes.InvalidArgument, "batch size cannot be greater than %d", entproto.MaxBatchCreateSize)
	}
	bulk := make([]*ent.MagazineCreate, len(requests))
	for i, req := range requests {
		magazine := req.GetMagazine()
		var err error
		bulk[i], err = svc.createBuilder(magazine)
		if err != nil {
			return nil, err
		}
	}
	res, err := svc.client.Magazine.CreateBulk(bulk...).Save(ctx)
	switch {
	case err == nil:
		protoList, err := toProtoMagazineList(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return connect.NewResponse(&api.BatchCreateMagazinesResponse{
			Magazines: protoList,
		}), nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

func (svc *MagazineService) Upsert(ctx context.Context, req *connect.Request[api.UpsertMagazineRequest]) (*connect.Response[api.Magazine], error) {
	mg := req.Msg.GetMagazine()
	m, err := svc.createBuilder(mg)
	if err != nil {
		return nil, err
	}
	err = m.OnConflict().UpdateName().Exec(ctx)
	switch {
	case err == nil:
		res, err := svc.client.Magazine.Query().Where(magazine.Name(mg.GetName())).Only(ctx)
		switch {
		case err == nil:
			proto, err := toProtoMagazine(res)
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

func (svc *MagazineService) createBuilder(magazine *api.Magazine) (*ent.MagazineCreate, error) {
	m := svc.client.Magazine.Create()
	magazineName := magazine.GetName()
	m.SetName(magazineName)
	for _, item := range magazine.GetComics() {
		var comics uuid.UUID
		if err := (&comics).UnmarshalBinary(item.GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
		}
		m.AddComicIDs(comics)
	}
	return m, nil
}
