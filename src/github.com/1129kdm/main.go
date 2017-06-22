package main

import (
	"github.com/1129kdm/services"
	"github.com/1129kdm/twitter"
	"io"
	"log"
	"net/http"
	"os"
	//"net/url"
)

func main() {
	twitterApi := twitter.AuthTwitterApi()

	for _, screenName := range util.TwitterNames() {
		tweets, err := twitterApi.GetUserTimeline(util.CreateUrlValues(screenName))
		if err != nil {
			log.Fatal(err)
		}

		var imgUrls []string
		for _, tweet := range tweets {
			if tweet.Entities.Media != nil {
				imgUrls = append(imgUrls, tweet.Entities.Media[0].Media_url)
			}
		}
		if imgUrls == nil {
			log.Printf("warn: " + screenName + "'s imgUrls is nil")
			continue
		}

		mkDirExistErr := util.MakeImgSaveDirectory(screenName)
		if mkDirExistErr != nil {
			log.Printf("warn: " + mkDirExistErr.Error())
		}

		for _, url := range imgUrls {
			response, err := http.Get(url)
			if err != nil {
				log.Printf("error: " + err.Error())
				continue
			}
			defer response.Body.Close()

			var imgPath = util.CreateSaveImgPath(util.ExtractImgNameFromUrl(url), screenName)
			file, err := os.Create(imgPath)
			if err != nil {
				log.Printf("error: %d", err.Error())
			}
			defer file.Close()

			io.Copy(file, response.Body)
		}
	}

	defer twitterApi.Close()
}
