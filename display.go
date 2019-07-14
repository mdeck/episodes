// display

package main

import (
	"fmt"
	"math"
	"time"
)

func (res *Results) DisplayInit() {
	fmt.Printf("Episodes:  ")
	res.start = time.Now()
}

func (res *Results) DisplayFinal() {
	fmt.Printf("[loaded in %s]\n\n", time.Since(res.start))

	now := GetMidnight(time.Now())
	for _, info := range res.infos {
		info.Print(&info.Prev, now, false)
		info.Print(&info.Next, now, true)
		println("")
	}
}

func (info *ShowInfo) Print(ep *Episode, now time.Time, isNext bool) {
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
