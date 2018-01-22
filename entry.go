// Copyright 2016-2018 mikan.

package chartgen

import (
	"bufio"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

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
	End          int        `json:"end,omitempty"`
}

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
	ending
)

func FetchEntry(key string) string {
	url := "https://raw.githubusercontent.com/aosn/aosn.github.io/master/workshop/" + key + ".md"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	return string(byteArray)
}

func ParseEntry(markdown string) []Record {
	var workshop Workshop
	var records []Record
	var users []string
	state := titleFinding // initial parser state
	scanner := bufio.NewScanner(strings.NewReader(markdown))
	for scanner.Scan() {
		line := scanner.Text()
		switch state {
		case titleFinding:
			if strings.HasPrefix(line, titlePrefix) {
				state = publisherFinding
				workshop.Title = line[len(titlePrefix):]
			}
		case publisherFinding:
			if strings.HasPrefix(line, publisherPrefix) {
				state = publishedFinding
				workshop.Publisher = strings.Replace(line[len(publisherPrefix):], " (訳書)", "", 1)
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
			}
		case pagesFinding:
			if strings.HasPrefix(line, pagesPrefix) {
				state = isbnFinding
				pages, err := strconv.Atoi(line[len(pagesPrefix):])
				if err != nil {
					panic(err) // invalid number format
				}
				workshop.Pages = pages
			}
		case isbnFinding:
			if strings.HasPrefix(line, isbnPrefix) {
				state = participantsFinding
				workshop.ISBN = line[len(isbnPrefix):]
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
			}
		case recordFinding:
			if strings.HasPrefix(line, recordPrefix) {
				var record Record
				var attends []string
				var notAttends []string
				record.Workshop = workshop
				columns := strings.Split(line, "|")

				// 1: #
				num, err := strconv.Atoi(strings.Trim(columns[1], " "))
				if err != nil {
					state = ending // not a parse target
					break
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
				total, err := strconv.Atoi(strings.Trim(columns[3], " "))
				if err != nil {
					state = ending
					break
				} else {
					record.AttendsTotal = total
				}

				// 4: :o: / :x:
				for n, ox := range strings.Replace(strings.Trim(columns[4], " "), ":", "", -1) {
					switch ox {
					case 'o':
						attends = append(attends, users[n])
					case 'x':
						notAttends = append(notAttends, users[n])
					}
				}

				// 5: p - p (or empty)

				// 6: ex (unused)

				record.Attends = attends
				record.NotAttends = notAttends
				records = append(records, record)
			} else {
				break
			}
		case ending:
			break
		default:
			continue // continue finding
		}
	}
	return records
}
