// sort

package main

import (
	"./parsers"
	"math"
	"sort"
	"time"
)

func (res *Results) Sort() {
	infos := res.infos
	sort.Slice(infos, func(i, j int) bool {
		return infos[i].DeltaDays() < infos[j].DeltaDays()
	})
}

func (info *ShowInfo) DeltaDays() float64 {
	now := parsers.GetMidnight(time.Now())
	prev := math.Abs(now.Sub(info.Prev.Airdate).Hours())
	next := math.Abs(now.Sub(info.Next.Airdate).Hours())
	return math.Min(prev, next)
}
