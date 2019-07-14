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
	"sync"
	"time"
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

func (res *Results) Populate() {
	fmt.Printf("Loading.. ")
	start := time.Now()

	var wg sync.WaitGroup
	for _, info := range res.infos {
		wg.Add(1)
		go info.Populate(&wg)
	}
	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("\nLoaded in %s\n\n", elapsed)
}

func (res *Results) Sort() {
	infos := res.infos
	sort.Slice(infos, func(i, j int) bool {
		return infos[i].DeltaDays() < infos[j].DeltaDays()
	})
}

func (res *Results) Display() {
	now := parsers.GetMidnight(time.Now())
	for _, info := range res.infos {
		info.Print(info.Prev, now, false)
		info.Print(info.Next, now, true)
		println("")
	}
}

func (info *ShowInfo) Print(ep *parsers.Episode, now time.Time, isNext bool) {
	if ep.Season == 0 && ep.Number == 0 {
		fmt.Printf(" %20s  -- Next episode unknown\n", "")
		return
	}

	showName := info.Name
	if isNext {
		showName = ""
	}
	days := math.Round(ep.Airdate.Sub(now).Hours() / 24)
	fmt.Printf(" %20s  %s %+4vd  [s%02v e%02v]  %s\n",
		showName, ep.Airdate.Format("2006.Jan.02"),
		int(days), ep.Season, ep.Number, ep.Name)
}

func (info *ShowInfo) DeltaDays() float64 {
	now := parsers.GetMidnight(time.Now())
	prev := math.Abs(now.Sub(info.Prev.Airdate).Hours())
	next := math.Abs(now.Sub(info.Next.Airdate).Hours())
	return math.Min(prev, next)
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
