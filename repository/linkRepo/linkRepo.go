package linkRepo

import (
	"song-chords-crawler/config"
	"song-chords-crawler/models"
)

type LinkRepo interface {
	GetSongUrls() []models.Link
}

func GetLinkRepo(host string) LinkRepo {
	switch host {
	case config.HOPAMVIET:
		return &HavLinkRepo{}
	}
	return nil
}
