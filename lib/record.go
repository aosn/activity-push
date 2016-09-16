// Copyright 2016 mikan.

package lib

import "time"

type Workshop struct {
	Title     string     `json:"title"`
	ISBN      string     `json:"isbn"`
	Pages     int        `json:"pages"`
	Publisher string     `json:"publisher"`
	Published SimpleDate `json:"published"`
}

type Record struct {
	Num          int        `json:"num"`
	Workshop     Workshop   `json:"workshop"`
	Date         SimpleDate `json:"date"`
	Attends      []string   `json:"attends"`
	NotAttends   []string   `json:"not-attends"`
	AttendsTotal int        `json:"attends-total"`
	Begin        int        `json:"begin,omitempty"`
	End          int        `json:"begin,omitempty"`
}

type SimpleDate struct {
	time.Time
}

func (d SimpleDate) format() string {
	return d.Time.Format("2006-01-02")
}

func (d SimpleDate) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.format() + `"`), nil
}
