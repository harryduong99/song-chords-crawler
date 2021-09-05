package crawlerFactory

import (
	"net/http"
	"song-chords-crawler/config"
	"song-chords-crawler/models"
)

type ICrawlerFactory interface {
	CrawlSong(*http.Response) models.Song
}

func GetCrawlerFactory(host string) ICrawlerFactory {
	switch host {
	case config.HOPAMVIET:
		return &HavCrawler{}
	}
	return nil
}
