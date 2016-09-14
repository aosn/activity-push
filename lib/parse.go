// Copyright 2016 mikan.

package lib

import (
	"bufio"
	"strconv"
	"strings"
	"time"
)

const titlePrefix = "title: "
const publisherPrefix = "* 出版: "
const publishedPrefix = "* 発売: "
const pagesPrefix = "* 頁数: "
const isbnPrefix = "* ISBN: "
const participantsPrefix = "| # | Date"
const participantPrefix = "![](/images/users/"
const recordPrefix = "| "
const (
	titleFinding = iota
	publisherFinding
	publishedFinding
	pagesFinding
	isbnFinding
	participantsFinding
	recordFinding
)

func Parse(markdown string) []Record {
	scanner := bufio.NewScanner(strings.NewReader(markdown))
	var workshop Workshop
	var records []Record
	var users []string
	state := titleFinding // initial parser state
	for scanner.Scan() {
		line := scanner.Text()
		switch state {
		case titleFinding:
			if strings.HasPrefix(line, titlePrefix) {
				state = publisherFinding
				workshop.Title = line[len(titlePrefix):]
			} else {
				continue
			}
		case publisherFinding:
			if strings.HasPrefix(line, publisherPrefix) {
				state = publishedFinding
				workshop.Publisher = strings.Replace(line[len(publisherPrefix):], " (訳書)", "", 1)
			} else {
				continue
			}
		case publishedFinding:
			if strings.HasPrefix(line, publishedPrefix) {
				state = pagesFinding
				format := "2006/01/02"
				t, err := time.Parse(format, line[len(publishedPrefix):])
				if err != nil {
					panic(err) // invalid date format
				}
				workshop.Published = SimpleDate{t}
			} else {
				continue
			}
		case pagesFinding:
			if strings.HasPrefix(line, pagesPrefix) {
				state = isbnFinding
				pages, err := strconv.Atoi(line[len(pagesPrefix):])
				if err != nil {
					panic(err) // invalid number format
				}
				workshop.Pages = pages
			} else {
				continue
			}
		case isbnFinding:
			if strings.HasPrefix(line, isbnPrefix) {
				state = participantsFinding
				workshop.ISBN = line[len(isbnPrefix):]
			} else {
				continue
			}
		case participantsFinding:
			if strings.HasPrefix(line, participantsPrefix) {
				for _, column := range strings.Split(line, "|") {
					column := strings.Trim(column, " ")
					if strings.HasPrefix(column, participantPrefix) {
						for _, element := range strings.Split(column, " ") {
							if strings.HasPrefix(element, participantPrefix) {
								end := strings.Index(element, "_")
								name := element[len(participantPrefix):end]
								users = append(users, name)
							}
						}
						break
					} else {
						continue
					}
				}
				state = recordFinding
			} else {
				continue
			}
		case recordFinding:
			if strings.HasPrefix(line, recordPrefix) {
				var record Record
				var attends []Participant
				record.Workshop = workshop
				columns := strings.Split(line, "|")

				// 1: #
				num, err := strconv.Atoi(strings.Trim(columns[1], " "))
				if err != nil {
					panic(err) // invalid number format
				}
				record.Num = num

				// 2: Date & Time
				format := "2006-01-02"
				t, err := time.Parse(format, strings.Trim(columns[2], " ")[0:10])
				if err != nil {
					panic(err) // invalid date format
				}
				record.Date = SimpleDate{t}

				// 3: A (unused)

				// 4: :o: / :x:
				for n, ox := range strings.Replace(strings.Trim(columns[4], " "), ":", "", -1) {
					var p Participant
					p.GitHubID = users[n]
					switch ox {
					case 'o':
						p.Attend = true
					default:
						p.Attend = false
					}
					attends = append(attends, p)
				}

				// 5: p - p (or empty)

				// 6: ex (unused)

				record.Attends = attends
				records = append(records, record)
			} else {
				break
			}
		default:
			continue // continue finding
		}
	}
	return records
}
