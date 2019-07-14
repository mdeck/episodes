// test3.go

package main

import (
	"./parsers"
	"time"
)

type Results struct {
	infos []*ShowInfo
	start time.Time
}

type ShowInfo struct {
	Imdb string
	Name string
	Prev *parsers.Episode
	Next *parsers.Episode
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
