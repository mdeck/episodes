// parsers

package parsers

import (
	"encoding/json"
	"errors"
	"time"
)

type EpisodeJSON struct {
	Name     string
	Season   int
	Number   int
	Airstamp time.Time
}

type ShowJSON struct {
	Name  string
	Links struct {
		PreviousEpisode struct {
			Href string
		}
		NextEpisode struct {
			Href string
		}
	} `json:"_links"`
}

var ErrEmpty = errors.New("body is empty")

func ParseEpisode(body string) (*EpisodeJSON, error) {
	j := EpisodeJSON{}

	if len(body) == 0 {
		return &j, ErrEmpty
	} else {
		err := json.Unmarshal([]byte(body), &j)
		return &j, err
	}
}

func ParseShow(body string) (*ShowJSON, error) {
	j := ShowJSON{}

	if len(body) == 0 {
		return &j, ErrEmpty
	} else {
		err := json.Unmarshal([]byte(body), &j)
		return &j, err
	}
}
