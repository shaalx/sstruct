package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func ReadAll(filename string) chan string {
	stringChan := make(chan string, 10)
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil
	}
	filestring := string(b)
	// newf := strings.Replace(filestring, "\n", "", -1)
	// fmt.Println(newf)
	reg := regexp.MustCompile(`\n+`)
	repstr := reg.ReplaceAllString(filestring, "")
	fmt.Println(repstr)
	filestrings := strings.Split(filestring, "。")
	go func([]string) {
		for _, it := range filestrings {
			stringChan <- it + "。"
			fmt.Println(it + "。")
		}
	}(filestrings)
	return stringChan
}

func SaveString(stringChan chan string) {
	file, err := os.OpenFile("result.txt", os.O_APPEND, 0764)
	if nil != err {
		return
	}
	for {
		str := <-stringChan
		file.WriteString(str + "\n")
	}
	defer file.Close()
}
