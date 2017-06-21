package main

import (
	"fmt"
	"github.com/1129kdm/haraeki"
	"github.com/1129kdm/twitter"
	"net/url"
)

func main() {
	twitterApi := twitter.AuthTwitterApi()
	v := url.Values{}
	v.Set("screen_name", "nakano_aimi")
	v.Set("count", "1")

	tweets, err := twitterApi.GetUserTimeline(v)
	if err != nil {
		panic(err)
	}

	for _, tweet := range tweets {
		if tweet.Entities.Media != nil {
			imgUrl := tweet.Entities.Media[0].Media_url
			fmt.Println("tweet: ", imgUrl)
		}
	}

	fmt.Println(haraeki.MemberTwitterNames())
	twitterApi.Close()
}
