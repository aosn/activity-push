// Copyright 2018 mikan.

package chartgen

import (
	"math"
	"sort"
)

type Entry struct {
	Name  string
	Value int
}
type OrderedEntries []Entry

func (l OrderedEntries) Len() int {
	return len(l)
}

func (l OrderedEntries) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l OrderedEntries) Less(i, j int) bool {
	if l[i].Value == l[j].Value {
		return l[i].Name < l[j].Name
	} else {
		return l[i].Value > l[j].Value // desc
	}
}

type Ranking struct {
	List  []Entry
	Max   int
	Min   int
	Total int
}

func AttendeeRanking(records []Record) Ranking {
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
	if totalAttends == 0 {
		return Ranking{[]Entry{}, 0, 0, 0}
	}
	return Ranking{desc(attendees), maxAttends, minAttends, totalAttends}
}

func desc(attendees map[string]int) []Entry {
	ranking := OrderedEntries{}
	for k, v := range attendees {
		e := Entry{k, v}
		ranking = append(ranking, e)
	}
	sort.Sort(ranking)
	return ranking
}
