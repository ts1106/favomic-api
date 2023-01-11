package router

import (
	"github.com/bufbuild/connect-go"
	"github.com/go-chi/chi/v5"
	"github.com/ts1106/favomic-api/ent"
	"github.com/ts1106/favomic-api/internal/service"
)

func episodeRouter(c *ent.Client) chi.Router {
	r := chi.NewRouter()
	svc := service.NewEpisodeService(c)
	r.Handle("/Create", connect.NewUnaryHandler(
		"/Create",
		svc.Create,
	))
	r.Handle("/Get", connect.NewUnaryHandler(
		"/Get",
		svc.Get,
	))
	r.Handle("/Update", connect.NewUnaryHandler(
		"/Update",
		svc.Update,
	))
	r.Handle("/Delete", connect.NewUnaryHandler(
		"/Delete",
		svc.Delete,
	))
	r.Handle("/List", connect.NewUnaryHandler(
		"/List",
		svc.List,
	))
	r.Handle("/BatchCreate", connect.NewUnaryHandler(
		"/BatchCreate",
		svc.BatchCreate,
	))
	r.Handle("/Upsert", connect.NewUnaryHandler(
		"/Upsert",
		svc.Upsert,
	))
	return r
}
