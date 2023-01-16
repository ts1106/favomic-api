package main

import (
	"github.com/mmcdole/gofeed"
	scrape "github.com/ts1106/favomic-api/cmd/scrape/internal"
)

const (
	magazineName = "少年ジャンプ＋"
)

func main() {
	s := scrape.NewScrape("http://localhost:8080")

	magazine := s.RegistMagazine(magazineName)

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://shonenjumpplus.com/atom")
	for _, item := range feed.Items {
		author := s.RegistAuthor(item.Authors[0])
		comic := s.RegistComic(item, author.GetId(), magazine.GetId())
		_ = s.RegistEpisode(item, comic.GetId())
	}
}
