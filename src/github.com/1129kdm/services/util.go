package util

import (
	"net/url"
	"os"
	"strings"
)

const (
	SAVE_IMG_PATH = "/tmp/"
	GET_TWEET_NUM = "200"
	URL_SEPARATE  = "/"
)

func TwitterNames() (m []string) {
	return []string{"kei_azm", "nanako_taya", "sena_akasaka"}
}

func MakeImgSaveDirectory(screenName string) error {
	return os.Mkdir(SAVE_IMG_PATH+screenName, 0777)
}

func CreateTweetsParams(screenName string) url.Values {
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

func ImgExist(imgName string, screenName string) bool {
	_, err := os.Stat(SAVE_IMG_PATH + screenName + "/" + imgName)
	return err == nil
}
