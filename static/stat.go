package main

import (
	"bufio"
	"fmt"
	"github.com/shaalx/sstruct/service/log"
	"github.com/shaalx/sstruct/vars"
	"os"
	"strings"
)

var (
	files = vars.Files
	stats = []string{"total_", "precise_", "recall_", "f-measure_"}
)

func main() {
	v2()
}

func v2() {
	leng1 := len(files)
	for i := 0; i < leng1; i++ {
		leng2 := len(files[i])
		// stat this item
		result := make([][]string, 0, 20)
		for ii := 0; ii < leng2; ii++ {
			data := Read4Lines("./stat/" + files[i][ii])
			result = append(result, data)
		}
		length := len(result)
		for j := 0; j < 4; j++ {
			content := ""
			for jj := 0; jj < length; jj++ {
				content += result[jj][j]
				fmt.Println(result[jj][j])
			}
			content = stats[j] + "\n" + content
			WriteFile(fmt.Sprintf("%s", files[i][0]), content)
		}

	}

}

func v1() {
	d, _ := os.Open("./stat/")
	f, _ := d.Readdir(0)
	result := make([][]string, 0, 100)
	for _, it := range f {
		data := Read4Lines("./stat/" + it.Name())
		result = append(result, data)
		// fmt.Println(i, data)
	}
	fmt.Println(result)
	length := len(result)
	content := ""
	for j := 0; j < 4; j++ {
		content = ""
		for i := 0; i < length; i++ {
			content += result[i][j]
			fmt.Println(result[i][j])
		}
		WriteFile(fmt.Sprintf("stat%d.txt", j), content)
	}
}

func Read4Lines(filename string) []string {
	rfile, err := os.Open(filename)
	defer rfile.Close()
	if log.IsError("open file error", err) {
		return nil
	}
	reader := bufio.NewReader(rfile)
	ret := make([]string, 0, 4)
	for i := 0; i < 4; i++ {
		str, err := reader.ReadString('\n')
		if log.IsError("{read file error}", err) {
			return nil
		}
		ss := strings.Split(str, ":")
		ret = append(ret, ss[1])
	}
	return ret
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
