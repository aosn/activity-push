// Copyright 2016 mikan.

package main

import (
	"encoding/json"

	"flag"

	"fmt"
	"math"
	"sort"

	"github.com/aosn/activity-push/lib"
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
		printStat(records)
		return
	}
	for _, record := range records {
		data, _ := json.Marshal(record)
		lib.Push("http://"+*host+":"+*port+"/aosn/record/", data)
	}
}

func printStat(records []lib.Record) {
	maxAttends := math.MinInt64
	minAttends := math.MaxInt64
	totalAttends := 0
	attendees := make(map[string]int)
	for _, record := range records {
		totalAttends += record.AttendsTotal
		if record.AttendsTotal > maxAttends {
			maxAttends = record.AttendsTotal
		}
		if record.AttendsTotal < minAttends {
			minAttends = record.AttendsTotal
		}
		for _, attendee := range record.Attends {
			attendees[attendee]++
		}
	}
	fmt.Println("## 統計")
	fmt.Println()
	fmt.Println("### 同時参加数")
	fmt.Println()
	fmt.Println("| 項目 | 値 |")
	fmt.Println("|:----:|:--:|")
	fmt.Printf("| 最大同時参加数 | %d |\n", maxAttends)
	fmt.Printf("| 最小同時参加数 | %d |\n", minAttends)
	fmt.Printf("| 平均同時参加数 | %2.2f |\n", float64(totalAttends)/float64(len(records)))
	fmt.Println()
	printRanking(attendees)
}

func printRanking(attendees map[string]int) {
	ranking := List{}
	for k, v := range attendees {
		e := Entry{k, v}
		ranking = append(ranking, e)
	}
	sort.Sort(ranking)
	fmt.Println("### 参加回数ランキング")
	fmt.Println()
	fmt.Println("| 順位 | 参加者 | 参加回数 |")
	fmt.Println("|:---:|:-------|:--------:|")
	for i, e := range ranking {
		fmt.Printf("| %d | ![](/images/users/%s_16.png) [%s](https://github.com/%s) | %d |\n",
			i+1, e.name, e.name, e.name, e.value)
	}
}

type Entry struct {
	name  string
	value int
}
type List []Entry

func (l List) Len() int {
	return len(l)
}

func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l List) Less(i, j int) bool {
	if l[i].value == l[j].value {
		return (l[i].name < l[j].name)
	} else {
		return (l[i].value > l[j].value) // desc
	}
}
