package crawlerFactory

import (
	"log"
	"net/http"
	"os"
	"song-chords-crawler/models"

	"github.com/PuerkitoBio/goquery"
)

type HavCrawler struct{}

func (havCrawler *HavCrawler) CrawlSong(res *http.Response) models.Song {
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var song models.Song
	song.Title = doc.Find("h1").Text()
	log.Println(doc.Find("#lyric").Html())
	os.Exit(1)
	// Find the review items
	doc.Find("#lyric").Each(func(i int, s *goquery.Selection) {
		song.Content, err = s.Html()
	})

	return song
}
