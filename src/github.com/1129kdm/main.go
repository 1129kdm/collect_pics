package main

import (
	"github.com/1129kdm/services"
	"github.com/1129kdm/twitter"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"sync"
)

const (
	GET_TWEET_LOOP_NUM = 16
)

func main() {
	log.Printf("info: Start get tweet images program")

	twitterApi := twitter.AuthTwitterApi()
	var wg sync.WaitGroup

	for _, screenName := range util.TwitterNames() {
		wg.Add(1)
		go func(sn string) {
			defer wg.Done()
			log.Printf("info: get user timeline of " + sn)

			var tweetsParams = util.CreateTweetsParams(sn)
			initTweets, err := twitterApi.GetUserTimeline(tweetsParams)
			if err != nil {
				log.Fatal(err)
			}

			// initial max_id
			var maxId = initTweets[0].Id
			var imgUrls []string

			// 遡れる上限が3200tweet (16 x 200)
			for i := 1; i <= GET_TWEET_LOOP_NUM; i++ {
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
						if util.ImgExist(util.ExtractImgNameFromUrl(tweet.Entities.Media[0].Media_url), sn) == true {
							log.Printf("info: " + sn + "'s img exist at " + strconv.Itoa(i) + " loop.")
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
				log.Printf("warn: " + sn + "'s imgUrls is nil. Finished get tweet images")
				runtime.Goexit()
			} else {
				log.Printf("info: " + sn + "'s imgUrls counts=" + strconv.Itoa(len(imgUrls)))
			}

			mkDirExistErr := util.MakeImgSaveDirectory(sn)
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

				var imgPath = util.CreateSaveImgPath(util.ExtractImgNameFromUrl(url), sn)
				file, err := os.Create(imgPath)
				if err != nil {
					log.Printf("error: %d", err.Error())
				}
				defer file.Close()

				io.Copy(file, response.Body)
			}

			log.Printf("info: Finished " + sn + " get tweet images")
		}(screenName)
	}
	wg.Wait()
	defer twitterApi.Close()
	log.Printf("info: Finished get tweet images program")
}
