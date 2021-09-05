package crawler

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
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
	for _, url := range songUrls {
		log.Println(url.Url)
		res, err := netClient.Get("https://hopamviet.vn/chord/song/nhan-toi-khoang-troi-em/W8IU77F6.html")

		if err != nil {
			log.Fatal(err)
		}

		defer res.Body.Close()

		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}

		song := crawlerFactory.CrawlSong(res)
		songRepo.StoreSong(song)
		os.Exit(1)
	}

}
