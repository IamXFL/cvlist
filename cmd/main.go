package main

import (
	"cvlist/crawler"
	"cvlist/model"
	"cvlist/names"
	"cvlist/xlog"
	"fmt"
	"runtime"
	"sort"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)
	go Monitor()

	xlog.Info("start job ...")
	names.Start()
	crawler.Start()
	output := crawler.Output()
	uis := make([]*model.UrlItem, 0, 200)
	patch := 1
	for {
		item, ok := <-output
		if !ok {
			xlog.InfoF("flush last patch: %v", patch)
			flush(uis)
			break
		}
		uis = append(uis, item)
		if len(uis) >= 200 {
			xlog.InfoF("flush patch: %v", patch)
			flush(uis)
			uis = uis[0:0] // clear
			patch++
		}
	}
	fmt.Println("Task all complete!")
}

func Monitor() {
	for {
		time.Sleep(time.Second * 1)
		fmt.Printf("goroutine num : %v \n", runtime.NumGoroutine())
	}
}

func flush(uis []*model.UrlItem) {
	sort.Slice(uis, func(i, j int) bool {
		return uis[i].Score > uis[j].Score
	})
	for i := 0; i < len(uis); i++ {
		if uis[i].Score >= 2 {
			xlog.InfoF("url: %v, score: %v", uis[i].Url, uis[i].Score)
		}
	}
}
