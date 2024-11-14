package crawler

import (
	"cvlist/client"
	"cvlist/filter"
	"cvlist/model"
	"cvlist/names"
	"cvlist/xlog"
	"fmt"
	"strings"
	"sync"
)

const MaxConcurrentNum = 100

var input chan string // names
var output chan *model.UrlItem

func init() {
	output = make(chan *model.UrlItem, 10000)
}

func Start() {
	xlog.Info("start crawl website ...")
	go func() {
		defer close(output)
		do()
	}()
}

func Output() chan *model.UrlItem {
	return output
}

func do() {
	maxConcurrenControl := make(chan struct{}, MaxConcurrentNum)
	input = names.Output()
	wg := &sync.WaitGroup{}
	mm := make(map[string]struct{})

	for {
		name, ok := <-input
		if !ok {
			break
		}
		// 去重
		if _, ok := mm[name]; ok {
			continue
		}
		mm[name] = struct{}{}

		wg.Add(1)
		maxConcurrenControl <- struct{}{}
		go func(innerName string) {
			defer func() {
				<-maxConcurrenControl
				wg.Done()
			}()

			process(innerName)
		}(name)
	}
	wg.Wait()
}

func process(name string) {
	url := "https://" + name + ".github.io"
	body, err := client.GET(url)
	if err != nil || body == "" {
		return
	}
	body = strings.ToLower(body)
	// findSubTask(url) // todo
	ui := &model.UrlItem{
		Url:   url,
		Score: 0.0,
	}
	for _, fiterFunc := range filter.FuncMap {
		if fiterFunc(&body) {
			ui.Score++
		}
	}
	if ui.Score >= 3 {
		fmt.Printf("url: %v, score: %v \n", ui.Url, ui.Score)
	}
	output <- ui
}
