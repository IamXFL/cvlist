package names

import (
	"cvlist/xlog"
	"strings"

	"github.com/mozillazg/go-pinyin"
)

var pinyinArgs = pinyin.NewArgs()

var output chan string

func init() {
	output = make(chan string, 100000)
}

// ToPinYin  将一个或多个汉字转换为拼音序列
func ToPinYin(word string) string {
	s := pinyin.LazyPinyin(word, pinyinArgs)
	res := strings.Join(s, "")
	return res
}

func Start() {
	xlog.Info("start name generation ...")
	go func() {
		defer close(output)
		// f1 + l
		F1LName()
		// f2 + l
		F2LName()
		// l + l
		LLName()

	}()
}

func Output() chan string {
	return output
}

func F1LName() {
	cnt := 0
	defer func() {
		xlog.InfoF("F1LName cnt: %v", cnt)
	}()

	for i := 0; i < len(firstNameL1); i++ {
		for j := 0; j < len(hanziUnicode); j++ {
			output <- ToPinYin(firstNameL1[i] + convertUnicodeToHanzi(hanziUnicode[j]))
			cnt++
		}
	}
}

func F2LName() {
	cnt := 0
	defer xlog.InfoF("F2LName cnt: %v", cnt)
	for i := 0; i < len(firstNameL2); i++ {
		for j := 0; j < len(hanziUnicode); j++ {
			output <- ToPinYin(firstNameL2[i] + convertUnicodeToHanzi(hanziUnicode[j]))
			cnt++
		}
	}
}

func LLName() {
	cnt := 0
	defer xlog.InfoF("LLName cnt: %v", cnt)
	for i := 0; i < len(hanziUnicode); i++ {
		for j := 0; j < len(hanziUnicode); j++ {
			output <- ToPinYin(convertUnicodeToHanzi(hanziUnicode[i]) + convertUnicodeToHanzi(hanziUnicode[j]))
			cnt++
		}
	}
}
