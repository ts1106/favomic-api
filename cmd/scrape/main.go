package main

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

func main() {
	ac := apiconnect.NewAuthorServiceClient(http.DefaultClient, "http://localhost:8080")
	ec := apiconnect.NewEpisodeServiceClient(http.DefaultClient, "http://localhost:8080")
	cc := apiconnect.NewComicServiceClient(http.DefaultClient, "http://localhost:8080")
	mc := apiconnect.NewMagazineServiceClient(http.DefaultClient, "http://localhost:8080")

	magazine := registMagazine(mc, "少年ジャンプ＋")

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://shonenjumpplus.com/atom")
	for _, item := range feed.Items {
		author := registAuthor(ac, item.Authors[0])
		comic := registComic(cc, item, author.GetId(), magazine.GetId())
		_ = registEpisode(ec, item, comic.GetId())
	}
}

func registMagazine(c apiconnect.MagazineServiceClient, name string) *api.Magazine {
	res, err := c.Upsert(
		context.Background(),
		connect.NewRequest(&api.UpsertMagazineRequest{Magazine: &api.Magazine{Name: name}}),
	)
	if err != nil {
		log.Fatal(err)
	}
	return res.Msg
}

func registAuthor(c apiconnect.AuthorServiceClient, a *gofeed.Person) *api.Author {
	res, err := c.Upsert(
		context.Background(),
		connect.NewRequest(&api.UpsertAuthorRequest{Author: &api.Author{Name: a.Name}}),
	)
	if err != nil {
		log.Fatal(err)
	}
	return res.Msg
}

func registEpisode(c apiconnect.EpisodeServiceClient, item *gofeed.Item, comicId []byte) *api.Episode {
	res, err := c.Upsert(
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

func registComic(c apiconnect.ComicServiceClient, item *gofeed.Item, authorId []byte, magazineId []byte) *api.Comic {
	res, err := c.Upsert(
		context.Background(),
		connect.NewRequest(&api.UpsertComicRequest{Comic: &api.Comic{Title: item.Content, AuthorId: authorId, MagazineId: magazineId}}),
	)
	if err != nil {
		log.Fatal(err)
	}
	return res.Msg
}
