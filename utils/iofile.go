package utils

import (
	"bufio"
	// "fmt"
	"github.com/shaalx/sstruct/service/log"
	. "github.com/shaalx/sstruct/vars"
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
	wfile, err := os.OpenFile(TEMP_FILENAME, os.O_CREATE|os.O_WRONLY, os.ModeDevice)
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
		// fmt.Println(bs)
		bs = bs[:len(bs)-2]
		// fmt.Println(bs)
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
	b, err := ioutil.ReadFile(TEMP_FILENAME)
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
	// 删除临时文件
	defer func(filename string) {
		os.Remove(filename)
	}(TEMP_FILENAME)
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

// 保存处理结果，结果从chan中读取
func SaveStringAppend(stringChan chan string, filename string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND, os.ModeDevice)
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

// 追加
func WriteFile(filename string, content string) {
	file, err := os.OpenFile(filename, os.O_CREATE, os.ModeDevice)
	if log.IsError("{Write to file error}", err) {
		return
	}
	file.WriteString(content + "\n")
	defer file.Close()
}

// 文件目录
func ReadDir(dir string) []string {
	files, _ := ioutil.ReadDir(dir)
	filenames := make([]string, 0)
	for _, it := range files {
		if !it.IsDir() {
			filenames = append(filenames, it.Name())
		}
	}
	return filenames
}

// 读取key文件
func ReadKey(filename string) (string, error) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		file, _ := os.OpenFile(filename, os.O_CREATE, 0664)
		defer file.Close()
		bs = []byte(",1,2,3,4,5,6,7,8,9,10,")
		WriteFile(filename, ",1,2,3,4,5,6,7,8,9,10,")
		err = nil
	}
	return string(bs), err
}
