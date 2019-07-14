// test3.go

package main

import (
	"time"
)

type Results struct {
	infos []*ShowInfo
	start time.Time
}

type ShowInfo struct {
	Imdb string
	Name string
	Prev Episode
	Next Episode
}

type Episode struct {
	Name    string
	Season  int
	Number  int
	Airdate time.Time
}

func main() {
	results := InitResults()
	results.DisplayInit()
	results.Populate()
	results.Sort()
	results.DisplayFinal()
}

func InitResults() (r Results) {
	r.infos = make([]*ShowInfo, len(imdbs))
	for idx := range r.infos {
		r.infos[idx] = &ShowInfo{Imdb: imdbs[idx]}
	}
	return
}

func GetMidnight(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}
