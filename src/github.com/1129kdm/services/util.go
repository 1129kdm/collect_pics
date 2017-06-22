package util

import (
	"net/url"
	"os"
	"strings"
)

const (
	SAVE_IMG_PATH = "/tmp/"
	GET_TWEET_NUM = "10"
	URL_SEPARATE  = "/"
)

func TwitterNames() (m []string) {
	return []string{"nakano_aimi", "kei_azm"}
}

func MakeImgSaveDirectory(screenName string) error {
	return os.Mkdir(SAVE_IMG_PATH+screenName, 0777)
}

func CreateUrlValues(screenName string) url.Values {
	v := url.Values{}
	v.Set("screen_name", screenName)
	v.Set("count", GET_TWEET_NUM)
	return v
}

func ExtractImgNameFromUrl(url string) string {
	var splitUrl = strings.Split(url, URL_SEPARATE)
	return splitUrl[len(splitUrl)-1]
}

func CreateSaveImgPath(imgName string, screenName string) string {
	return SAVE_IMG_PATH + screenName + "/" + imgName
}
