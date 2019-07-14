// populate

package main

import (
	"./parsers"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

func (res *Results) Populate() {
	var wg sync.WaitGroup
	for _, info := range res.infos {
		wg.Add(1)
		go info.Populate(&wg)
	}
	wg.Wait()
}

func (info *ShowInfo) Populate(wg *sync.WaitGroup) {
	url := "http://api.tvmaze.com/lookup/shows?imdb=" + info.Imdb
	j := parsers.ParseShow(getBody(url))

	prev := make(chan *Episode)
	next := make(chan *Episode)
	go makeEpisodeRequest(j.Links.PreviousEpisode.Href, prev)
	go makeEpisodeRequest(j.Links.NextEpisode.Href, next)

	info.Name = j.Name
	info.Prev = <-prev
	info.Next = <-next

	wg.Done()
}

func makeEpisodeRequest(url string, c chan<- *Episode) {
	j := parsers.ParseEpisode(getBody(url))
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
