// Copyright 2018 mikan.

package chartgen

import (
	"bufio"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

const indexTablePrefix = "| #  | "
const indexHeaderPrefix = "|---:|"
const (
	indexFinding = iota
	entryFinding
)

func FetchIndex() string {
	url := "https://raw.githubusercontent.com/aosn/aosn.github.io/master/index.md"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	return string(byteArray)
}

func ParseIndex(markdown string) []string {
	var entries []string
	state := indexFinding // initial parser state
	entryMatcher := regexp.MustCompile(`\(/workshop/([a-z0-9-]+)\)`)
	scanner := bufio.NewScanner(strings.NewReader(markdown))
	for scanner.Scan() {
		line := scanner.Text()
		switch state {
		case indexFinding:
			if strings.HasPrefix(line, indexTablePrefix) {
				state = entryFinding
			}
		case entryFinding:
			if len(strings.TrimSpace(line)) == 0 {
				return entries
			}
			if strings.HasPrefix(line, indexHeaderPrefix) {
				continue
			}
			matched := entryMatcher.FindStringSubmatch(line)
			if len(matched) > 0 {
				entries = append(entries, matched[1]) // 0: all, 1: first match
			}
		default:
			continue // continue finding
		}
	}
	return entries
}
