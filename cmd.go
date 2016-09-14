// Copyright 2016 mikan.

package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aosn/activity-push/push"
)

func main() {
	fmt.Printf("hoge")
	location, _ := time.LoadLocation("Asia/Tokyo")
	var authors, translators []string
	authors = append(authors, "Cay S. Horstmann")
	translators = append(translators, "柴田 芳樹")
	var attends []push.Participant
	attends = append(attends, push.Participant{"mikan", "Japan", "Kanagawa", true})
	record := push.Record{
		push.Workshop{
			"Java SE 8 実践プログラミング",
			"978-4-8443-3667-9",
			264,
			authors,
			translators,
			"インプレス",
			push.SimpleDate{time.Date(2014, 9, 22, 0, 0, 0, 0, location)},
			"A",
		},
		push.SimpleDate{time.Date(2014, 11, 23, 0, 0, 0, 0, location)},
		attends,
		211,
		214,
		2,
	}
	bytes, _ := json.Marshal(record)
	fmt.Printf("%s", bytes)
}
