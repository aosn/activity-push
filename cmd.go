// Copyright 2016 mikan.

package main

import (
	"encoding/json"

	"flag"

	"github.com/aosn/activity-push/lib"
)

var target = flag.String("t", "", "target record (e.g. 1-java8)")
var host = flag.String("h", "localhost", "host for Elasticsearch")
var port = flag.String("p", "9200", "port for Elasticsearch")

func main() {
	flag.Parse()
	if *target == "" || *host == "" || *port == "" {
		flag.Usage()
		return
	}
	records := lib.Parse(lib.Fetch(*target))
	for _, record := range records {
		data, _ := json.Marshal(record)
		lib.Push("http://"+*host+":"+*port+"/aosn/record/", data)
	}
}
