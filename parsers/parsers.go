// parsers

package parsers

import (
	"encoding/json"
	"log"
	"time"
)

type Episode struct {
	Name    string
	Season  int
	Number  int
	Airdate time.Time `json:"-"`
	//AirdateString string `json:"airdate"`
	Airstamp time.Time
}

type Show struct {
	Name    string
	PrevURL string
	NextURL string
}

type ShowJSON struct {
	Name  string `json:"name"`
	Links struct {
		PreviousEpisode struct {
			Href string `json:"href"`
		} `json:"previousepisode"`
		NextEpisode struct {
			Href string `json:"href"`
		} `json:"nextepisode"`
	} `json:"_links"`
}

func GetMidnight(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}

func ParseEpisode(body string) *Episode {
	e := new(Episode)

	if len(body) == 0 {
		e.Airdate = time.Now().Add(time.Hour * 24 * 365 * 50)
		return e
	}

	err := json.Unmarshal([]byte(body), e)
	if err != nil {
		log.Fatalln(err)
	}

	e.Airdate = GetMidnight(e.Airstamp.Local())
	return e
}

func ParseShow(body string) *Show {
	s := new(Show)

	var sj ShowJSON
	if err := json.Unmarshal([]byte(body), &sj); err != nil {
		log.Fatalln(err)
	}

	s.Name = sj.Name
	s.PrevURL = sj.Links.PreviousEpisode.Href
	s.NextURL = sj.Links.NextEpisode.Href
	return s
}
