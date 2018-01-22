// Copyright 2016-2018 mikan.

package main

import (
	"fmt"

	"github.com/aosn/chartgen"
)

func PrintStat(records []chartgen.Record) {
	ranking := chartgen.AttendeeRanking(records)
	avg := 0.0
	if ranking.Total > 0 {
		avg = float64(ranking.Total) / float64(len(records))
	}
	fmt.Println("## 統計")
	fmt.Println()
	fmt.Println("### 同時参加数")
	fmt.Println()
	fmt.Println("| 項目 | 値 |")
	fmt.Println("|:----:|:--:|")
	fmt.Printf("| 最大同時参加数 | %d |\n", ranking.Max)
	fmt.Printf("| 最小同時参加数 | %d |\n", ranking.Min)
	fmt.Printf("| 平均同時参加数 | %2.2f |\n", avg)
	fmt.Println()
	fmt.Println("### 参加回数ランキング")
	fmt.Println()
	fmt.Println("| 順位 | 参加者 | 参加回数 |")
	fmt.Println("|:---:|:-------|:--------:|")
	for i, e := range ranking.List {
		fmt.Printf("| %d | ![](/images/users/%s_16.png) [%s](https://github.com/%s) | %d |\n",
			i+1, e.Name, e.Name, e.Name, e.Value)
	}
}
