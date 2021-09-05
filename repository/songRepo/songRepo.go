package songRepo

import (
	"context"
	"song-chords-crawler/config"
	"song-chords-crawler/databaseDriver"
	"song-chords-crawler/models"

	"go.mongodb.org/mongo-driver/bson"
)

func StoreSong(song models.Song) bool {
	collection := databaseDriver.Mongo.ConnectCollection(config.DB_NAME, config.COL_SONGS)

	bbytes, _ := bson.Marshal(song)
	_, err := collection.InsertOne(context.Background(), bbytes)

	return err == nil
}
