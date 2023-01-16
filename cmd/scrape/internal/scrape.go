package scrape

import (
	"context"
	"log"
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/mmcdole/gofeed"
	api "github.com/ts1106/favomic-api/gen/api/v1"
	apiconnect "github.com/ts1106/favomic-api/gen/api/v1/v1connect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Scrape struct {
	ac apiconnect.AuthorServiceClient
	ec apiconnect.EpisodeServiceClient
	cc apiconnect.ComicServiceClient
	mc apiconnect.MagazineServiceClient
}

func NewScrape(url string) *Scrape {
	return &Scrape{
		ac: apiconnect.NewAuthorServiceClient(http.DefaultClient, url),
		ec: apiconnect.NewEpisodeServiceClient(http.DefaultClient, url),
		cc: apiconnect.NewComicServiceClient(http.DefaultClient, url),
		mc: apiconnect.NewMagazineServiceClient(http.DefaultClient, url),
	}
}

func (s *Scrape) RegistMagazine(name string) *api.Magazine {
	res, err := s.mc.Upsert(
		context.Background(),
		connect.NewRequest(&api.UpsertMagazineRequest{Magazine: &api.Magazine{Name: name}}),
	)
	if err != nil {
		log.Fatal(err)
	}
	return res.Msg
}

func (s *Scrape) RegistAuthor(a *gofeed.Person) *api.Author {
	res, err := s.ac.Upsert(
		context.Background(),
		connect.NewRequest(&api.UpsertAuthorRequest{Author: &api.Author{Name: a.Name}}),
	)
	if err != nil {
		log.Fatal(err)
	}
	return res.Msg
}

func (s *Scrape) RegistEpisode(item *gofeed.Item, comicId []byte) *api.Episode {
	res, err := s.ec.Upsert(
		context.Background(),
		connect.NewRequest(
			&api.UpsertEpisodeRequest{Episode: &api.Episode{
				Title:     item.Title,
				Url:       item.Link,
				Thumbnail: item.Enclosures[0].URL,
				UpdatedAt: timestamppb.New(*item.UpdatedParsed),
				ComicId:   comicId,
			}}),
	)
	if err != nil {
		log.Fatal(err)
	}
	return res.Msg
}

func (s *Scrape) RegistComic(item *gofeed.Item, authorId []byte, magazineId []byte) *api.Comic {
	res, err := s.cc.Upsert(
		context.Background(),
		connect.NewRequest(&api.UpsertComicRequest{Comic: &api.Comic{Title: item.Content, AuthorId: authorId, MagazineId: magazineId}}),
	)
	if err != nil {
		log.Fatal(err)
	}
	return res.Msg
}
