// Copyright 2018 mikan.

package main

import "github.com/aosn/chartgen"

type DataSet struct {
	ByTimes     IntLabeledData    `json:"by-times"`
	ByAttendees StringLabeledData `json:"by-attendees"`
}

type IntLabeledData struct {
	Labels []int `json:"labels"`
	Data   []int `json:"data"`
}

type StringLabeledData struct {
	Labels []string `json:"labels"`
	Data   []int    `json:"data"`
}

func Convert(records []chartgen.Record) DataSet {
	// by times
	var byTimesLabels []int
	var byTimesRecords []int
	for _, record := range records {
		byTimesLabels = append(byTimesLabels, record.Num)
		byTimesRecords = append(byTimesRecords, record.AttendsTotal)
	}

	// by attendees
	var byAttendeesLabels []string
	var byAttendeesRecords []int
	for _, entry := range chartgen.AttendeeRanking(records).List {
		byAttendeesLabels = append(byAttendeesLabels, entry.Name)
		byAttendeesRecords = append(byAttendeesRecords, entry.Value)
	}

	return DataSet{IntLabeledData{byTimesLabels, byTimesRecords}, StringLabeledData{byAttendeesLabels, byAttendeesRecords}}
}
