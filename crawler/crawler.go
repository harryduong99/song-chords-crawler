package crawler

import (
	"crypto/tls"
	"log"
	"net/http"
	"song-chords-crawler/crawler/crawlerFactory"
	"song-chords-crawler/repository/linkRepo"
	"song-chords-crawler/repository/songRepo"
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

	crawlerFactory := crawlerFactory.GetCrawlerFactory(host)
	for i, url := range songUrls {
		log.Println(i)
		log.Println(url.Url)

		res, err := netClient.Get(url.Url)

		if err != nil {
			log.Fatal(err)
		}

		defer res.Body.Close()

		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}

		song := crawlerFactory.CrawlSong(res)
		result := songRepo.StoreSong(song)
		if result {
			url.Crawled = true
			linkRepo.UpdateLink(url.ID, url)
		}
	}

}
