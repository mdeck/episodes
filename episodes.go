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
	"sync"
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
	results := make([]*ShowInfo, len(imdbs))
	doneShow := make(chan int)
	fmt.Printf("Loading.. ")
	for idx, imdb := range imdbs {
		go makeShowRequest(imdb, &results[idx], doneShow)
	}
	for idx, _ := range imdbs {
		fmt.Printf("%v.. ", len(results)-idx)
		<-doneShow
	}
	fmt.Printf("\n\n")
	return results
}

func sortResults(results []*ShowInfo) {
	sort.Slice(results, func(i, j int) bool {
		return results[i].DeltaDays() < results[j].DeltaDays()
	})
}

func displayResults(results []*ShowInfo) {
	// display
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

func makeShowRequest(imdb string, ptrInfo **ShowInfo, doneShow chan int) {
	info := new(ShowInfo)
	url := "http://api.tvmaze.com/lookup/shows?imdb=" + imdb
	show := parsers.ParseShow(getBody(url))
	info.Name = show.Name

	var wg sync.WaitGroup
	wg.Add(2)
	go makeEpisodeRequest(show.PrevURL, &info.Prev, &wg)
	go makeEpisodeRequest(show.NextURL, &info.Next, &wg)
	wg.Wait()

	*ptrInfo = info
	doneShow <- 1
}

func makeEpisodeRequest(url string, ptrEp **parsers.Episode, wg *sync.WaitGroup) {
	*ptrEp = parsers.ParseEpisode(getBody(url))
	wg.Done()
}

func getBody(url string) string {
	//println("http request", url)
	if len(url) == 0 {
		return ""
	}
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
