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

type HavLinkRepo struct{}

func (linkRepo *HavLinkRepo) GetSongUrls() []models.Link {
	var links []models.Link
	collection := databaseDriver.Mongo.ConnectCollection(config.DB_NAME, config.COL_LINKS)
	filter := bson.D{{"url", primitive.Regex{Pattern: "/chord/song/", Options: ""}}}

	cur, err := collection.Find(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var link models.Link
		// & character returns the memory address of the following variable.
		err := cur.Decode(&link) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		links = append(links, link)
	}

	return links
}
