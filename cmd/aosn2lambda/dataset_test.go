// Copyright 2018 mikan.

package main

import (
	"testing"

	"github.com/aosn/chartgen"
)

func TestConvert(t *testing.T) {
	record1 := chartgen.Record{
		Num:      1,
		Workshop: chartgen.Workshop{},
		Date:     chartgen.SimpleDate{},
		Attends: []string{"user1", "user2"},
		NotAttends: []string{"user3", "user4"},
		AttendsTotal: 2,
		Begin: 100,
		End: 101}
	record2 := chartgen.Record{
		Num:      2,
		Workshop: chartgen.Workshop{},
		Date:     chartgen.SimpleDate{},
		Attends: []string{"user1", "user2", "user3", "user4"},
		NotAttends: []string{},
		AttendsTotal: 4,
		Begin: 102,
		End: 103}
	records := []chartgen.Record{record1, record2}
	dataSet := Convert(records)
	if len(dataSet.ByTimes.Labels) != 2 {
		t.Fatalf("expected %d, actual %d", 2, len(dataSet.ByTimes.Labels))
	}
	if len(dataSet.ByTimes.Data) != 2 {
		t.Fatalf("expected %d, actual %d", 2, len(dataSet.ByTimes.Data))
	}
	if len(dataSet.ByAttendees.Labels) != 4 {
		t.Fatalf("expected %d, actual %d", 4, len(dataSet.ByTimes.Labels))
	}
	if len(dataSet.ByAttendees.Data) != 4 {
		t.Fatalf("expected %d, actual %d", 4, len(dataSet.ByTimes.Data))
	}
}
