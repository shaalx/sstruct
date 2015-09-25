package main

import (
	"bufio"
	"fmt"
	"github.com/shaalx/sstruct/service/log"
	"github.com/shaalx/sstruct/vars"
	"os"
	"strconv"
	"strings"
)

var (
	files       = vars.Files
	stats       = []string{"total_", "precise_", "recall_", "f-measure_", "key", "key", "ATT-ATT-ATT-", "ATT-ATT-SBV-", "ATT-SBV-", "ATT-VOB-", "ATT-HED-", "ATT-POB-", "ADV-ATT-", "FOB-VOB-", "SBV-SBV-", "SBV-ATT-", "SBV-VOB-", "SBV-COO-", "SBV-HED-"}
	rels        = []string{"ATT-ATT-ATT-", "ATT-ATT-SBV-", "ATT-SBV-", "ATT-VOB-", "ATT-HED-", "ATT-POB-", "ADV-ATT-", "FOB-VOB-", "SBV-SBV-", "SBV-ATT-", "SBV-VOB-", "SBV-COO-", "SBV-HED-"}
	res_statMap = make(map[string]map[int]float64)
)

func init() {
	res_statMap["C19"] = make(map[int]float64)
	res_statMap["C34"] = make(map[int]float64)
	res_statMap["C32"] = make(map[int]float64)
	res_statMap["C38"] = make(map[int]float64)
	res_statMap["C6-"] = make(map[int]float64)
	res_statMap["C3-"] = make(map[int]float64)
	res_statMap["C7-"] = make(map[int]float64)
	res_statMap["C31"] = make(map[int]float64)
	res_statMap["C11"] = make(map[int]float64)
	res_statMap["C39"] = make(map[int]float64)
}

func main() {
	v2()
}

func v2() {

	file, err := os.OpenFile("rel_stat.txt", os.O_APPEND, 0644)
	defer file.Close()
	if log.IsError("{Write to file error}", err) {
		return
	}

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
			filenameType := (files[i][ii])[:3]
			for iii, rateContent := range data {
				f, err := strconv.ParseFloat(rateContent[:7], 10)
				if err != nil {
					fmt.Println(err)
				}
				res_statMap[filenameType][iii] += f
				// fmt.Println(res_statMap)
			}
			// fmt.Print(data)
			result[ii] = append(result[ii], data...)
		}

		// for j, iti := range result {
		// 	content := ""
		// 	for _, itj := range iti {
		// 		content += itj
		// 		fmt.Print(j, "|", len(itj), "\t")
		// 	}
		// 	content = stats[j] + "\n" + content
		// 	file.WriteString(files[i][0] + "\n")
		// 	WriteFile(file, content)
		// }

	}

	for k, v := range res_statMap {
		fmt.Println(k, v)
		file.WriteString("\n\n" + k + "\n")
		for ti, it := range rels {
			fmt.Println(it)
			file.WriteString(fmt.Sprintf("%v\n", v[ti]*10))
		}
	}

}

func Read19Lines(filename string) []string {
	rfile, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND, 0644)
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
		if i < 6 {
			continue
		}
		ss := strings.Split(str, ":")
		ret = append(ret, ss[1])
	}
	return ret
}

// 追加
func WriteFile(file *os.File, content string) {
	file.WriteString(content + "\n")
	// defer file.Close()
}
