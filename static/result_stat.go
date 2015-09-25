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
	stats = []string{"total_", "precise_", "recall_", "f-measure_", "key", "key", "ATT-ATT-ATT-", "ATT-ATT-SBV-", "ATT-SBV-", "ATT-VOB-", "ATT-HED-", "ATT-POB-", "ADV-ATT-", "FOB-VOB-", "SBV-SBV-", "SBV-ATT-", "SBV-VOB-", "SBV-COO-", "SBV-HED-"}
)

func main() {
	v2()
}

func v2() {
	leng1 := len(files)
	for i := 0; i < leng1; i++ {
		leng2 := len(files[i])
		// stat this item
		result := make([][]string, leng2, 18)
		for k := 0; k < leng2; k++ {
			result[i] = make([]string, 0, 21)
		}
		for ii := 0; ii < leng2; ii++ {
			data := Read19Lines("./stat/" + files[i][ii])
			// fmt.Print(data)
			result[ii] = append(result[ii], data...)
		}
		// result[10-1][19-1]
		// for m := 0; m < 10; m++ {
		// 	for n := 0; n < 19; n++ {
		// 		fmt.Println(result[m][n], "---")
		// 	}
		// }
		// length := len(result)
		// fmt.Println(result)
		for j, iti := range result {
			content := ""
			for _, itj := range iti {
				content += itj
				fmt.Print(j, "|", len(itj), "\t")
			}
			content = stats[j] + "\n" + content
			WriteFile(fmt.Sprintf("%s", files[i][0]), content)
		}
		// for j := 0; j < 19; j++ {
		// 	content := ""
		// 	for jj := 0; jj < length; jj++ {
		// 		content += result[jj][j]
		// 		// fmt.Println(result[jj][j])
		// 	}
		// 	content = stats[j] + "\n" + content
		// 	WriteFile(fmt.Sprintf("%s", files[i][0]), content)
		// }

	}

}

func v1() {
	// ioutil.
	// d, _ := os.Open("./stat/")
	// f, _ := d.Readdir(0)
	// result := make([][]string, 0, 100)
	// for _, it := range f {
	// 	data := Read19Lines("./stat/" + it.Name())
	// 	result = append(result, data)
	// 	// fmt.Println(i, data)
	// }
}

func Read19Lines(filename string) []string {
	rfile, err := os.Open(filename)
	defer rfile.Close()
	if log.IsError("open file error", err) {
		return nil
	}
	reader := bufio.NewReader(rfile)
	ret := make([]string, 0, 21)
	for i := 0; i < 19; i++ {
		str, err := reader.ReadString('\n')
		if log.IsError("{read file error}", err) {
			panic(err.Error())
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
