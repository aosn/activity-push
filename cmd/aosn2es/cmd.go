// Copyright 2016-2018 mikan.

package main

import (
	"encoding/json"
	"flag"

	"github.com/aosn/chartgen"
)

var target = flag.String("t", "", "target record (e.g. 1-java8)")
var host = flag.String("h", "localhost", "host name of your Elasticsearch")
var port = flag.String("p", "9200", "port number of your Elasticsearch")

func main() {
	flag.Parse()
	if *target == "" || *host == "" || *port == "" {
		flag.Usage()
		return
	}
	records := chartgen.Parse(chartgen.Fetch(*target))
	for _, record := range records {
		data, _ := json.Marshal(record)
		Push("http://"+*host+":"+*port+"/aosn/record/", data)
	}
}
