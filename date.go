// Copyright 2018 mikan.

package chartgen

import "time"

type SimpleDate struct {
	time.Time
}

func (d SimpleDate) format() string {
	return d.Time.Format("2006-01-02")
}

func (d SimpleDate) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.format() + `"`), nil
}
