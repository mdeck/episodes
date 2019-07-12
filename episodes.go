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
	println("Episodes:")
	println("")

	imdbs := []string{
		"tt2191671", // Elementary
		"tt5296406", // Designated Survivor
		"tt1595859", // Blue Bloods
		"tt5691552", // The Orville
		"tt1721666", // Sister Wives
		"tt3502248", // Bosch
		"tt5171438", // Star Trek Discovery
		"tt4955642", // The Good Place
		"tt5770786", // GLOW
		"tt0475784", // Westworld
		"tt4687880", // Goliath
		"tt2575988", // Silicon Valley
		"tt4158110", // Mr. Robot
		"tt4574334", // Stranger Things
		"tt1475582", // Sherlock
	}

	// for each imdb, get ShowInfo & add to slice
	is := make([]*ShowInfo, len(imdbs))
	fmt.Printf("Loading.. ")
	for idx, imdb := range imdbs {
		fmt.Printf("%v.. ", len(is)-idx)
		is[idx] = getInfo(imdb)
	}
	fmt.Printf("\n\n")

	sort.Slice(is, func(i, j int) bool {
		return is[i].DeltaDays() < is[j].DeltaDays()
	})

	// display
	for _, i := range is {
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

func getInfo(imdb string) *ShowInfo {
	var i ShowInfo
	url := "http://api.tvmaze.com/lookup/shows?imdb=" + imdb
	s := parsers.ParseShow(getBody(url))
	i.Name = s.Name
	i.Prev = parsers.ParseEpisode(getBody(s.PrevURL))
	i.Next = parsers.ParseEpisode(getBody(s.NextURL))
	return &i
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
