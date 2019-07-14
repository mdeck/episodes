// parsers

package parsers

import (
	"encoding/json"
	"log"
	"time"
)

type EpisodeJSON struct {
	Name     string
	Season   int
	Number   int
	Airstamp time.Time
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

func ParseEpisode(body string) *EpisodeJSON {
	j := EpisodeJSON{}

	if len(body) == 0 {
		// Default airstamp shouldn't influence sort order
		j.Airstamp = time.Now().Add(time.Hour * 24 * 365 * 50)
		return &j
	}

	err := json.Unmarshal([]byte(body), &j)
	if err != nil {
		log.Fatalln(err)
	}

	return &j
}

func ParseShow(body string) *ShowJSON {
	j := ShowJSON{}

	if err := json.Unmarshal([]byte(body), &j); err != nil {
		log.Fatalln(err)
	}

	return &j
}
