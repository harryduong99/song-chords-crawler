package linkRepo

import (
	"context"
	"log"
	"song-chords-crawler/config"
	"song-chords-crawler/databaseDriver"
	"song-chords-crawler/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func UpdateLink(objID primitive.ObjectID, link models.Link) bool {
	log.Println(link.Crawled)
	collection := databaseDriver.Mongo.ConnectCollection(config.DB_NAME, config.COL_LINKS)

	filter := bson.M{"_id": bson.M{"$eq": objID}}
	update := bson.M{"$set": bson.M{"crawled": link.Crawled, "url": link.Url, "domain": link.Domain}}

	rs, err := collection.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	log.Printf("adsf")
	log.Println(rs)
	return err == nil

}
