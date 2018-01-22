// Copyright 2018 mikan.

package chartgen

import (
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	original := "2018-01-01"
	parsed, err := time.Parse("2006-01-02", original)
	if err != nil {
		t.Error(err) // invalid date format
	}
	date := SimpleDate{parsed}
	if date.format() != original {
		t.Errorf("expected %s, actual %s", original, date.format())
	}
}
