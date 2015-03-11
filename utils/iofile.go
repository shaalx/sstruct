package utils

import (
	"bufio"
	// "fmt"
	"github.com/shaalx/sstruct/service/log"
	"io/ioutil"
	"os"
	"strings"
)

func deleteEnter(filename string) {
	// 读文件
	rfile, err := os.Open(filename)
	defer rfile.Close()
	if log.IsError("open file error", err) {
		return
	}
	reader := bufio.NewReader(rfile)

	// 写文件
	wfile, err := os.OpenFile("w_"+filename, os.O_CREATE, os.ModeDevice)
	defer wfile.Close()
	if log.IsError("open file error", err) {
		return
	}
	writer := bufio.NewWriter(wfile)

	// 边读边写
	for {
		str, err := reader.ReadString('\n')
		if log.IsError("{read file error}", err) {
			return
		}
		bs := []byte(str)
		bs = bs[:len(bs)-2]
		_, err = writer.WriteString(string(bs))
		writer.Flush()
		if log.IsError("{write to file error}", err) {

		}
	}
}

func ReadAll(filename string) chan string {
	deleteEnter(filename)
	stringChan := make(chan string, 10)
	b, err := ioutil.ReadFile("w_" + filename)
	if err != nil {
		return nil
	}
	filestring := string(b)
	filestrings := strings.Split(filestring, "。")
	go func([]string) {
		for _, substring := range filestrings {
			substrings := strings.Split(substring, "，")
			for _, it := range substrings {
				if 0 >= len(it) {
					continue
				}
				stringChan <- it + "。"
				// fmt.Println(it + "。")
				log.LOGS.Debug("%s%s\n", it, "。")
			}
		}
	}(filestrings)
	return stringChan
}

func SaveString(stringChan chan string) {
	file, err := os.OpenFile("result.txt", os.O_CREATE, os.ModeDevice)
	if nil != err {
		return
	}
	for {
		str := <-stringChan
		file.WriteString(str + "\t")
		str2 := <-stringChan
		file.WriteString(str2 + "\n")
	}
	defer file.Close()
}
