package twitter

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/joho/godotenv"
	"log"
	. "os"
)

func envLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("load .env error")
	}
}

func AuthTwitterApi() *anaconda.TwitterApi {
	envLoad()
	anaconda.SetConsumerKey(Getenv("TWITTER_CONSUMER_KEY"))
	anaconda.SetConsumerSecret(Getenv("TWITTER_CONSUMER_SECRET"))
	api := anaconda.NewTwitterApi(Getenv("TWITTER_ACCESS_TOKEN"), Getenv("TWITTER_ACCESS_TOKEN_SECRET"))
	return api
}
