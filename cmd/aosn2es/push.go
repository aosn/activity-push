// Copyright 2016-2018 mikan.

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Push(url string, body []byte) {
	fmt.Printf("%s", string(body))
	resp, _ := http.Post(url, "application/json", strings.NewReader(string(body)))
	defer resp.Body.Close()
	responseBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("----------")
	fmt.Println(resp.Status)
	fmt.Println(string(responseBody))
	if resp.StatusCode >= 300 {
		panic("ERROR " + resp.Status)
	}
}
