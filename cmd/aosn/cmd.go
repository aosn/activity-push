// Copyright 2016-2018 mikan.

package main

import (
	"flag"

	"github.com/aosn/chartgen"
)

var target = flag.String("t", "", "target record (e.g. 1-java8)")

func main() {
	flag.Parse()
	if *target == "" {
		flag.Usage()
		return
	}
	records := chartgen.Parse(chartgen.Fetch(*target))
	PrintStat(records)
	return
}
