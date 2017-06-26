package main

import (
	"github.com/1129kdm/services"
	"github.com/1129kdm/twitter"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	twitterApi := twitter.AuthTwitterApi()

	for _, screenName := range util.TwitterNames() {
		log.Printf("info: get user timeline of " + screenName)

		var tweetsParams = util.CreateTweetsParams(screenName)
		initTweets, err := twitterApi.GetUserTimeline(tweetsParams)
		if err != nil {
			log.Fatal(err)
		}

		// initial max_id
		var maxId = initTweets[0].Id
		var imgUrls []string

		// 遡れる上限が3200tweet (16 x 200)
		for i := 1; i <= 16; i++ {
			// 指定したmax_id以下のtweetを取得する
			tweetsParams.Set("max_id", strconv.FormatInt(maxId, 10))
			tweets, err := twitterApi.GetUserTimeline(tweetsParams)
			if err != nil {
				log.Fatal(err)
			}

			if len(tweets) == 0 {
				continue
			}

			for _, tweet := range tweets {
				if tweet.Entities.Media != nil {
					// ローカルに存在している画像が見つかったらツイート取得ループを抜ける
					if util.ImgExist(util.ExtractImgNameFromUrl(tweet.Entities.Media[0].Media_url), screenName) == true {
						log.Printf("info: " + screenName + "'s img exist at " + strconv.Itoa(i) + " loop.")
						goto URL_CHECK
					}
					imgUrls = append(imgUrls, tweet.Entities.Media[0].Media_url)
				}
			}
			// 最後のId - 1 をしないと取得するツイートの初めが重複する
			maxId = tweets[len(tweets)-1].Id - 1
		}
	URL_CHECK:
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
