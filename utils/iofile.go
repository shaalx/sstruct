package utils

import (
	"bufio"
	// "fmt"
	"github.com/shaalx/sstruct/service/log"
	"io/ioutil"
	"os"
	"strings"
)

// 删除文件中的enter键 ：[10 13]，最后将文件内容保存至w_filename.txt中
func DeleteEnter(filename string) {
	// 读文件
	rfile, err := os.Open(filename)
	defer rfile.Close()
	if log.IsError("open file error", err) {
		return
	}
	reader := bufio.NewReader(rfile)

	// 写文件
	wfile, err := os.OpenFile("w_"+filename, os.O_CREATE|os.O_WRONLY, os.ModeDevice)
	defer wfile.Close()
	if log.IsError("open file error", err) {
		return
	}
	writer := bufio.NewWriter(wfile)

	// 边读边写
	for {
		str, err := reader.ReadString('\n')
		if log.IsError("{read file error}", err) {
			writer.Flush()
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

// 读取文件 1，删除enter键； 2，读取处理结果到chan中，返回chan
func ReadAll(filename string) chan string {
	DeleteEnter(filename)
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
		stringChan <- "end"
	}(filestrings)
	return stringChan
}

// 保存处理结果，结果从chan中读取
func SaveString(stringChan chan string, filename string) {
	file, err := os.OpenFile(filename, os.O_CREATE, os.ModeDevice)
	if nil != err {
		return
	}
	for {
		str := <-stringChan
		file.WriteString(str)
	}
	defer file.Close()
}

// 追加
func AppendFile(filename string, content string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND, os.ModeDevice)
	if log.IsError("{append to file error}", err) {
		return
	}
	file.WriteString(content + "\n")
	defer file.Close()
}
