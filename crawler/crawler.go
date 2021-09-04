package crawler

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"song-chords-crawler/repository/linkRepo"

	"github.com/PuerkitoBio/goquery"
)

var (
	config = &tls.Config{
		InsecureSkipVerify: true,
	}
	transport = &http.Transport{
		TLSClientConfig: config,
	}
	netClient = &http.Client{
		Transport: transport,
	}
	// queue          = make(chan string)
	// hasVisited     = make(map[string]bool)
	// chunksToInsert = []models.Link{}
	// w              = sync.WaitGroup{}
)

func Crawl(host string) {
	songUrls := linkRepo.GetLinkRepo(host).GetSongUrls()
	log.Println(len(songUrls))
	for _, url := range songUrls {
		log.Println(url.Url)
		res, err := netClient.Get(url.Url)

		if err != nil {
			log.Fatal(err)
		}

		defer res.Body.Close()

		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}

		// Load the HTML document
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Find the review items
		doc.Find("h1").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the title
			title := s.Text()
			fmt.Printf("Review %d: %s\n", i, title)
		})
	}

}
