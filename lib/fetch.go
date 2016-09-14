// Copyright 2016 mikan.

package lib

import (
	"io/ioutil"
	"net/http"
)

func Fetch(key string) string {
	url := "https://raw.githubusercontent.com/aosn/aosn.github.io/master/workshop/" + key + ".md"
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	return string(byteArray)
}
