// populate

package main

import (
	"./parsers"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func (res *Results) Populate() {
	res.infos = make([]*ShowInfo, len(imdbs))
	c := make(chan *ShowInfo)
	for idx := range res.infos {
		url := "http://api.tvmaze.com/lookup/shows?imdb=" + imdbs[idx]
		go MakeShowInfo(url, c)
	}
	for idx := range res.infos {
		res.infos[idx] = <-c
	}
}

func MakeShowInfo(url string, c chan<- *ShowInfo) {
	j, err := parsers.ParseShow(getBody(url))
	if err != nil && err != parsers.ErrEmpty {
		log.Fatalln(err)
	}

	prev, next := make(chan *Episode), make(chan *Episode)
	go MakeEpisode(j.Links.PreviousEpisode.Href, prev)
	go MakeEpisode(j.Links.NextEpisode.Href, next)

	c <- &ShowInfo{j.Name, <-prev, <-next}
}

func MakeEpisode(url string, c chan<- *Episode) {
	j, err := parsers.ParseEpisode(getBody(url))
	if err != nil {
		if err != parsers.ErrEmpty {
			log.Fatalln(err)
		}
		// Default airstamp shouldn't influence sort order
		j.Airstamp = time.Now().Add(time.Hour * 24 * 365 * 50)
	}

	airdate := GetMidnight(j.Airstamp.Local())
	c <- &Episode{j.Name, j.Season, j.Number, airdate}
}

func getBody(url string) string {
	if len(url) == 0 {
		return ""
	}
	url = strings.Replace(url, "http://", "https://", 1)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	resp.Body.Close()
	return string(body)
}
