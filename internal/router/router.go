package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ts1106/favomic-api/ent"
)

func NewRouter(c *ent.Client) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Mount("/api.AuthorService", authorRouter(c))
	r.Mount("/api.ComicService", comicRouter(c))
	r.Mount("/api.EpisodeService", episodeRouter(c))
	r.Mount("/api.MagazineService", magazineRouter(c))
	r.Mount("/api.TagService", tagRouter(c))
	return r
}
