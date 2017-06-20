package main

import (
	"fmt"
	"github.com/1129kdm/twitter"
)

func main() {
	twitterApi := twitter.AuthTwitterApi()
	fmt.Println(twitterApi)

	text := "test ichi chan"
	tweet, err := twitterApi.PostTweet(text, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(tweet.Text)
}
