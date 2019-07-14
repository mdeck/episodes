// test3.go

package main

import (
	"./parsers"
)

type Results struct {
	infos []*ShowInfo
}

type ShowInfo struct {
	Imdb string
	Name string
	Prev *parsers.Episode
	Next *parsers.Episode
}

func main() {
	println("Episodes:\n")
	results := InitResults()
	results.Populate()
	results.Sort()
	results.Display()
}

func InitResults() (r Results) {
	r.infos = make([]*ShowInfo, len(imdbs))
	for idx := range r.infos {
		r.infos[idx] = &ShowInfo{Imdb: imdbs[idx]}
	}
	return
}
