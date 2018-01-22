// Copyright 2016-2018 mikan.

package main

import (
	"flag"
	"fmt"

	"github.com/aosn/chartgen"
)

var target = flag.String("t", "", "target record (e.g. 1-java8)")
var all = flag.Bool("a", false, "execute all")

func main() {
	flag.Parse()
	if *all {
		for _, l := range chartgen.ParseIndex(chartgen.FetchIndex()) {
			fmt.Printf("==========%s==========\n", l)
			records := chartgen.ParseEntry(chartgen.FetchEntry(l))
			PrintStat(records)
		}
		return
	}
	if *target == "" {
		for _, l := range chartgen.ParseIndex(chartgen.FetchIndex()) {
			fmt.Println(l)
		}
		return
	}
	records := chartgen.ParseEntry(chartgen.FetchEntry(*target))
	PrintStat(records)
	return
}
