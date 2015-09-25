package main

import (
	"fmt"
	"github.com/shaalx/sstruct/service/log"
	"github.com/shaalx/sstruct/vars"
	"io/ioutil"
	"os"
)

var (
	files = vars.Files
	items = []string{"economy", "agriculture", "art", "history", "environment", "politics", "space", "sports", "philosophy", "computer"}
)

func main() {
	v2()
}

func v2() {
	leng1 := len(files)
	var count int32
	count = int32(0)
	content := ""
	for i := 0; i < leng1; i++ {
		leng2 := len(files[i])
		// stat this item
		for ii := 0; ii < leng2; ii++ {
			c := FileWrodsCount("./origin/" + items[i] + "/" + files[i][ii])
			count += c
			content += fmt.Sprintf("%s :%v\n", files[i][ii], c)
			fmt.Printf("%s :%v\n", files[i][ii], c)
		}
	}

	content += fmt.Sprintf("total:%v\n avg:%v\n", count, count/100)
	WriteFile("avg_words_stats.txt", content)

}

func FileWrodsCount(filename string) int32 {
	// ioutil.ReadFile(filename)
	b, _ := ioutil.ReadFile(filename)
	return int32(len(b) / 3)
}

// 追加
func WriteFile(filename string, content string) {
	file, err := os.OpenFile(filename, os.O_CREATE, os.ModeDevice)
	if log.IsError("{Write to file error}", err) {
		return
	}
	file.WriteString(content + "\n")
	defer file.Close()
}
