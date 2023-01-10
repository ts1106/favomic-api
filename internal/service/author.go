package service

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/ts1106/favomic-api/ent"
	"github.com/ts1106/favomic-api/ent/author"
	api "github.com/ts1106/favomic-api/gen/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthorService struct {
	client *ent.Client
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
			Id: id,
		})
	}
	return v, nil
}

func (svc *AuthorService) Search(ctx context.Context, req *connect.Request[api.SearchAuthorRequest]) (*connect.Response[api.Author], error) {
	var (
		err error
		get *ent.Author
	)
	switch req.Msg.GetView() {
	case api.SearchAuthorRequest_BASIC:
		get, err = svc.client.Author.Query().
			Where(author.Name(req.Msg.GetName())).
			Only(ctx)
	case api.SearchAuthorRequest_WITH_EDGES:
		get, err = svc.client.Author.Query().
			Where(author.Name(req.Msg.GetName())).
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
