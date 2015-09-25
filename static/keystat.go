package main

import (
	"fmt"
	"github.com/shaalx/sstruct/service/log"
	"github.com/shaalx/sstruct/vars"
	"io/ioutil"
	"os"
	"strings"
)

var (
	files = vars.Files
)

func main() {
	v2()
}

func v2() {
	leng1 := len(files)
	single := 0
	double := 0
	for i := 0; i < leng1; i++ {
		leng2 := len(files[i])
		// stat this item
		for ii := 0; ii < leng2; ii++ {
			s1, d2 := Read1Lines("./key/" + files[i][ii])
			single += s1
			double += d2
		}
	}

	content := fmt.Sprintf("single:%d\n double:%d\n", single, double)
	WriteFile("key_stats.txt", content)

}

func Read1Lines(filename string) (int, int) {
	// ioutil.ReadFile(filename)
	b, _ := ioutil.ReadFile(filename)
	keys := string(b)
	single := 0
	key_slice := strings.Split(keys, ",")
	length := len(key_slice) - 2
	key_slice = key_slice[1 : length+1]
	for _, it := range key_slice {
		if len(it) <= 9 {
			single += 1
		}
	}
	fmt.Println(key_slice, length)
	return single, len(key_slice) - single
}

// 追加
func WriteFile(filename string, content string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND, os.ModeDevice)
	if log.IsError("{Write to file error}", err) {
		return
	}
	file.WriteString(content + "\n")
	defer file.Close()
}
