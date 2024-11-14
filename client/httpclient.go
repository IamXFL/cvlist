package client

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var HttpClient *http.Client

const TIMEOUT_SECOND = 10

func init() {
	tr := &http.Transport{
		MaxIdleConns:        200,
		IdleConnTimeout:     30 * time.Second,
		DisableKeepAlives:   false,
		TLSHandshakeTimeout: 10 * time.Second,
	}
	HttpClient = &http.Client{
		Transport: tr,
	}
}

//.github.io  github常用
// .com： 通用顶级域名，最常见，易于记忆，适合各种类型的个人网站。
//.net： 适用于网络相关的工作，如程序员、前端开发等。
//.dev： 专为开发者设计的，非常适合程序员、技术人员。
//.me： 表示“我”，强调个人品牌，适合个人博客或简历网站。
//.io： 近年来很流行，常用于科技类网站。
//.xyz： 新兴域名，价格相对较低，适合个人项目。
//.org: 通常用于非营利组织、社区或开源项目。

// 获取body string
func GET(url string) (string, error) {
	return findMeValid(url)
}

func findMeValid(url string) (string, error) {

	resp, err := HttpClient.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func findBlogValid(str string) (string, error) {
	// var url = "https://" + str + ".github.io"
	var url = "https://" + str + ".blog"
	resp, err := HttpClient.Get(url)
	if err != nil {
		return url, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return url, err
	}

	if resp.StatusCode == 200 {
		bstr := string(body)
		if strings.Contains(strings.ToLower(bstr), "电话") && strings.Contains(strings.ToLower(bstr), "邮箱") {
			fmt.Println(string(url))
			return url, nil
		} else if strings.Contains(strings.ToLower(bstr), "tel") {
			fmt.Println(string(url))
			return url, nil
		} else if strings.Contains(strings.ToLower(bstr), "姓名") {
			fmt.Println(string(url))
			return url, nil
		} else if strings.Contains(strings.ToLower(bstr), "邮箱") {
			fmt.Println(string(url))
			return url, nil
		}
	}
	return url, errors.New("status not 200")
}
