package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ts1106/favomic-api/ent"
)

func NewRouter(c *ent.Client) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Mount("/entpb.AuthorService", authorRouter(c))
	r.Mount("/entpb.ComicService", comicRouter(c))
	r.Mount("/entpb.EpisodeService", episodeRouter(c))
	r.Mount("/entpb.MagazineService", magazineRouter(c))
	r.Mount("/entpb.TagService", tagRouter(c))
	return r
}
