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
	show := parsers.ParseShow(getBody(url))
	info.Name = show.Name

	prev := make(chan *parsers.Episode)
	next := make(chan *parsers.Episode)
	go makeEpisodeRequest(show.PrevURL, prev)
	go makeEpisodeRequest(show.NextURL, next)
	info.Prev = <-prev
	info.Next = <-next

	wg.Done()
}

func makeEpisodeRequest(url string, c chan<- *parsers.Episode) {
	c <- parsers.ParseEpisode(getBody(url))
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
