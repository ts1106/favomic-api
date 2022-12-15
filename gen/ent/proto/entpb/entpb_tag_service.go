// Code generated by protoc-gen-entgrpc. DO NOT EDIT.
package entpb

import (
	context "context"
	base64 "encoding/base64"
	entproto "entgo.io/contrib/entproto"
	sqlgraph "entgo.io/ent/dialect/sql/sqlgraph"
	fmt "fmt"
	connect_go "github.com/bufbuild/connect-go"
	uuid "github.com/google/uuid"
	ent "github.com/ts1106/favomic-api/ent"
	comic "github.com/ts1106/favomic-api/ent/comic"
	tag "github.com/ts1106/favomic-api/ent/tag"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// TagService implements TagServiceServer
type TagService struct {
	client *ent.Client
}

// NewTagService returns a new TagService
func NewTagService(client *ent.Client) *TagService {
	return &TagService{
		client: client,
	}
}

// toProtoTag transforms the ent type to the pb type
func toProtoTag(e *ent.Tag) (*Tag, error) {
	v := &Tag{}
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
		v.Comics = append(v.Comics, &Comic{
			Id: id,
		})
	}
	return v, nil
}

// toProtoTagList transforms a list of ent type to a list of pb type
func toProtoTagList(e []*ent.Tag) ([]*Tag, error) {
	var pbList []*Tag
	for _, entEntity := range e {
		pbEntity, err := toProtoTag(entEntity)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		pbList = append(pbList, pbEntity)
	}
	return pbList, nil
}

// Create implements TagServiceServer.Create
func (svc *TagService) Create(ctx context.Context, req *connect_go.Request[CreateTagRequest]) (*connect_go.Response[Tag], error) {
	tag := req.Msg.GetTag()
	m, err := svc.createBuilder(tag)
	if err != nil {
		return nil, err
	}
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoTag(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return connect_go.NewResponse(proto), nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Get implements TagServiceServer.Get
func (svc *TagService) Get(ctx context.Context, req *connect_go.Request[GetTagRequest]) (*connect_go.Response[Tag], error) {
	var (
		err error
		get *ent.Tag
	)
	var id uuid.UUID
	if err := (&id).UnmarshalBinary(req.Msg.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	switch req.Msg.GetView() {
	case GetTagRequest_VIEW_UNSPECIFIED, GetTagRequest_BASIC:
		get, err = svc.client.Tag.Get(ctx, id)
	case GetTagRequest_WITH_EDGE_IDS:
		get, err = svc.client.Tag.Query().
			Where(tag.ID(id)).
			WithComics(func(query *ent.ComicQuery) {
				query.Select(comic.FieldID)
			}).
			Only(ctx)
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid argument: unknown view")
	}
	switch {
	case err == nil:
		proto, err := toProtoTag(get)
		return connect_go.NewResponse(proto), err
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Update implements TagServiceServer.Update
func (svc *TagService) Update(ctx context.Context, req *connect_go.Request[UpdateTagRequest]) (*connect_go.Response[Tag], error) {
	tag := req.Msg.GetTag()
	var tagID uuid.UUID
	if err := (&tagID).UnmarshalBinary(tag.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	m := svc.client.Tag.UpdateOneID(tagID)
	tagName := tag.GetName()
	m.SetName(tagName)
	for _, item := range tag.GetComics() {
		var comics uuid.UUID
		if err := (&comics).UnmarshalBinary(item.GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
		}
		m.AddComicIDs(comics)
	}

	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoTag(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return connect_go.NewResponse(proto), nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Delete implements TagServiceServer.Delete
func (svc *TagService) Delete(ctx context.Context, req *connect_go.Request[DeleteTagRequest]) (*connect_go.Response[emptypb.Empty], error) {
	var err error
	var id uuid.UUID
	if err := (&id).UnmarshalBinary(req.Msg.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	err = svc.client.Tag.DeleteOneID(id).Exec(ctx)
	switch {
	case err == nil:
		return connect_go.NewResponse(&emptypb.Empty{}), nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// List implements TagServiceServer.List
func (svc *TagService) List(ctx context.Context, req *connect_go.Request[ListTagRequest]) (*connect_go.Response[ListTagResponse], error) {
	var (
		err      error
		entList  []*ent.Tag
		pageSize int
	)
	pageSize = int(req.Msg.GetPageSize())
	switch {
	case pageSize < 0:
		return nil, status.Errorf(codes.InvalidArgument, "page size cannot be less than zero")
	case pageSize == 0 || pageSize > entproto.MaxPageSize:
		pageSize = entproto.MaxPageSize
	}
	listQuery := svc.client.Tag.Query().
		Order(ent.Desc(tag.FieldID)).
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
			Where(tag.IDLTE(pageToken))
	}
	switch req.Msg.GetView() {
	case ListTagRequest_VIEW_UNSPECIFIED, ListTagRequest_BASIC:
		entList, err = listQuery.All(ctx)
	case ListTagRequest_WITH_EDGE_IDS:
		entList, err = listQuery.
			WithComics(func(query *ent.ComicQuery) {
				query.Select(comic.FieldID)
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
		protoList, err := toProtoTagList(entList)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return connect_go.NewResponse(&ListTagResponse{
			TagList:       protoList,
			NextPageToken: nextPageToken,
		}), nil
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// BatchCreate implements TagServiceServer.BatchCreate
func (svc *TagService) BatchCreate(ctx context.Context, req *connect_go.Request[BatchCreateTagsRequest]) (*connect_go.Response[BatchCreateTagsResponse], error) {
	requests := req.Msg.GetRequests()
	if len(requests) > entproto.MaxBatchCreateSize {
		return nil, status.Errorf(codes.InvalidArgument, "batch size cannot be greater than %d", entproto.MaxBatchCreateSize)
	}
	bulk := make([]*ent.TagCreate, len(requests))
	for i, req := range requests {
		tag := req.GetTag()
		var err error
		bulk[i], err = svc.createBuilder(tag)
		if err != nil {
			return nil, err
		}
	}
	res, err := svc.client.Tag.CreateBulk(bulk...).Save(ctx)
	switch {
	case err == nil:
		protoList, err := toProtoTagList(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return connect_go.NewResponse(&BatchCreateTagsResponse{
			Tags: protoList,
		}), nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

func (svc *TagService) createBuilder(tag *Tag) (*ent.TagCreate, error) {
	m := svc.client.Tag.Create()
	tagName := tag.GetName()
	m.SetName(tagName)
	for _, item := range tag.GetComics() {
		var comics uuid.UUID
		if err := (&comics).UnmarshalBinary(item.GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
		}
		m.AddComicIDs(comics)
	}
	return m, nil
}