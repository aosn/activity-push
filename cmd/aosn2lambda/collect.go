// Copyright 2018 mikan.

package main

import (
	"encoding/json"
	"github.com/aosn/chartgen"
	"log"
)

// CollectAndPush generates chart.js datasets and locate to S3.
func CollectAndPush() {
	// fetch index and iterate it
	for _, entry := range chartgen.ParseIndex(chartgen.FetchIndex()) {
		// parse entry
		dataset, err := json.Marshal(Convert(chartgen.ParseEntry(chartgen.FetchEntry(entry))))
		if err != nil {
			log.Fatal(err)
		}
		Upload(entry, dataset)
	}
}
