package main

import (
	"log"
	"song-chords-crawler/config"
	"song-chords-crawler/crawler"
	"song-chords-crawler/databaseDriver"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	crawler.Crawl(config.HOPAMVIET)
}

func loadDatabase() {
	databaseDriver.ConnectDatabase()
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
