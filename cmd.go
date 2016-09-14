// Copyright 2016 mikan.

package main

import (
	"encoding/json"

	"flag"

	"github.com/aosn/activity-push/lib"
)

var target = flag.String("target", "", "target (e.g. 1-java8)")

func main() {
	flag.Parse()
	if *target == "" {
		flag.Usage()
		return
	}
	records := lib.Parse(lib.Fetch(*target))
	for _, record := range records {
		data, _ := json.Marshal(record)
		lib.Push("http://192.168.1.207:9200/aosn/record/", data)
	}
}
