// Copyright 2016 mikan.

package push

import "time"

type Workshop struct {
	Title       string     `json:title`
	ISBN        string     `json:isbn`
	Pages       int        `json:pages`
	Authors     []string   `json:authors`
	Translators []string   `json:translators,omitempty`
	Publisher   string     `json:publisher`
	Published   SimpleDate `json:published`
	Part        string     `json:part`
}

type Participant struct {
	GitHubID string `json:github_id`
	Country  string `json:country`
	State    string `json:city`
	Attend   bool   `json:attend`
}

type Record struct {
	Workshop    Workshop      `json:workshop`
	Date        SimpleDate    `json:date`
	Attends     []Participant `json:attends`
	Begin       int           `json:begin`
	End         int           `json:begin`
	TimeMinutes int           `json:time_minutes`
}

type SimpleDate struct {
	time.Time
}

func (d SimpleDate) format() string {
	return d.Time.Format("2016-12-31")
}

func (d SimpleDate) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.format() + `"`), nil
}
