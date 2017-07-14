// Copyright 2016-2017 mikan.

package main

import (
	"encoding/json"

	"flag"

	"github.com/aosn/activity-push/lib"
	"github.com/aosn/activity-push/stats"
)

var target = flag.String("t", "", "target record (e.g. 1-java8)")
var host = flag.String("h", "localhost", "host name of your Elasticsearch")
var port = flag.String("p", "9200", "port number of your Elasticsearch")
var stat = flag.Bool("stat", false, "print statistics mode (don't push Elasticsearch)")

func main() {
	flag.Parse()
	if *target == "" || *host == "" || *port == "" {
		flag.Usage()
		return
	}
	records := lib.Parse(lib.Fetch(*target))
	if *stat {
		stats.PrintStat(records)
		return
	}
	for _, record := range records {
		data, _ := json.Marshal(record)
		lib.Push("http://"+*host+":"+*port+"/aosn/record/", data)
	}
}
