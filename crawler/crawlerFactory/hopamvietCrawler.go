package crawlerFactory

import (
	"log"
	"net/http"
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
	song.Content, err = doc.Find("#lyric").Html()
	if err != nil {
		log.Printf("Get content failed!")
	}

	song.Author = doc.Find(".fa-leaf").Next().Text()
	song.Category = doc.Find(".fa-book").Next().Text()

	// Find the review items
	// doc.Find("#lyric").Each(func(i int, s *goquery.Selection) {
	// 	song.Content, err = s.Html()
	// })

	return song
}
