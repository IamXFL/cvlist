package crawler

import (
	"cvlist/xlog"

	"github.com/antchfx/htmlquery"
)

func findSubTask(url string) []string {
	// url = "https//baidu.com"
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		xlog.Error(err.Error())
	}

	nodes := htmlquery.Find(doc, "//a")
	subHrefs := make([]string, 0, 256)
	for _, node := range nodes {
		href := htmlquery.SelectAttr(node, "href")
		// fmt.Println(href)
		subHrefs = append(subHrefs, href)
	}
	return subHrefs
}
