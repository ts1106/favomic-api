package main

import (
	"context"
	"log"
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/mmcdole/gofeed"
	"github.com/ts1106/favomic-api/gen/ent/proto/entpb"
	api "github.com/ts1106/favomic-api/gen/ent/proto/entpb/entpbconnect"
)

func main() {
	ac := api.NewAuthorServiceClient(http.DefaultClient, "http://localhost:8080")
	// ec := entpbconnect.NewEpisodeServiceClient(http.DefaultClient, "http://localhost:8080")
	// cc := entpbconnect.NewComicServiceClient(http.DefaultClient, "http://localhost:8080")

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://shonenjumpplus.com/atom")
	for _, item := range feed.Items {
		// fmt.Println(item.Title)             // episode title
		// fmt.Println(item.Content)           // magazine title
		// fmt.Println(item.Link)              // episode link
		// fmt.Println(item.Updated)           // updated at
		// fmt.Println(item.Authors[0].Name)   // author
		// fmt.Println(item.Enclosures[0].URL) // thumbnail
		registAuthor(ac, item.Authors[0])
		break
	}
}

func registAuthor(c api.AuthorServiceClient, a *gofeed.Person) {
	res, err := c.Create(
		context.Background(),
		connect.NewRequest(&entpb.CreateAuthorRequest{Author: &entpb.Author{Name: a.Name}}),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v", res.Msg)
}
