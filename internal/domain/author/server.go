package author

import (
	"context"

	"github.com/ts1106/favomic-api/ent"
)

type Service interface {
	Create(context.Context, CreateRequest) (*ent.Author, error)
}

type Server struct {
	Service
	client *ent.Client
}

func NewServer(c *ent.Client) *Server {
	return &Server{client: c}
}

func (s *Server) Create(ctx context.Context, r CreateRequest) (*ent.Author, error) {
	q := s.client.Author.Create()
	q.SetName(r.Name)
	q.AddComicIDs(r.Comics...)

	e, err := q.Save(ctx)
	if err != nil {
		return nil, err
	}
	return e, nil
}
