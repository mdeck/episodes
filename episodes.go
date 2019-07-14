// test3.go

package main

import (
	"./parsers"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"sort"
	"strings"
	"time"
)

type ShowInfo struct {
	Name string
	Prev *parsers.Episode
	Next *parsers.Episode
}

func (i *ShowInfo) DeltaDays() float64 {
	now := parsers.GetMidnight(time.Now())
	prev := math.Abs(now.Sub(i.Prev.Airdate).Hours())
	next := math.Abs(now.Sub(i.Next.Airdate).Hours())
	return math.Min(prev, next)
}

func main() {
	println("Episodes:\n")
	results := retrieveShowEpisodeInfo()
	sortResults(results)
	displayResults(results)
}

func retrieveShowEpisodeInfo() []*ShowInfo {
	fmt.Printf("Loading.. ")
	start := time.Now()

	results := make([]*ShowInfo, len(imdbs))
	c := make(chan *ShowInfo)
	for _, imdb := range imdbs {
		go makeShowRequest(imdb, c)
	}
	for idx, _ := range imdbs {
		fmt.Printf("%v.. ", len(results)-idx)
		results[idx] = <-c
	}

	elapsed := time.Since(start)
	fmt.Printf("\nLoaded in %s\n\n", elapsed)
	return results
}

func sortResults(results []*ShowInfo) {
	sort.Slice(results, func(i, j int) bool {
		return results[i].DeltaDays() < results[j].DeltaDays()
	})
}

func displayResults(results []*ShowInfo) {
	for _, i := range results {
		now := parsers.GetMidnight(time.Now())

		days := math.Round(i.Prev.Airdate.Sub(now).Hours() / 24)
		fmt.Printf(" %20s  %s %+4vd  [s%02v e%02v]  %s\n",
			i.Name, i.Prev.Airdate.Format("2006.Jan.02"),
			int(days), i.Prev.Season, i.Prev.Number, i.Prev.Name)

		if i.Next.Season == 0 && i.Next.Number == 0 {
			fmt.Printf(" %20s  -- Next episode unknown\n", "")
		} else {
			days = math.Round(i.Next.Airdate.Sub(now).Hours() / 24)
			fmt.Printf(" %20s  %s %+4vd  [s%02v e%02v]  %s\n",
				"", i.Next.Airdate.Format("2006.Jan.02"),
				int(days), i.Next.Season, i.Next.Number, i.Next.Name)
		}

		println("")
	}
}

func makeShowRequest(imdb string, c chan<- *ShowInfo) {
	info := new(ShowInfo)
	url := "http://api.tvmaze.com/lookup/shows?imdb=" + imdb
	show := parsers.ParseShow(getBody(url))
	info.Name = show.Name

	prev := make(chan *parsers.Episode)
	next := make(chan *parsers.Episode)
	go makeEpisodeRequest(show.PrevURL, prev)
	go makeEpisodeRequest(show.NextURL, next)
	info.Prev = <-prev
	info.Next = <-next

	c <- info
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
