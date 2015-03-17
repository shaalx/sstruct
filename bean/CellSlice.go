package bean

import (
	"fmt"
	"github.com/shaalx/sstruct/utils"
	"regexp"
)

type Stats map[string]int32

type Cell struct {
	Word string
	Freq int32
}

type CellSlice []*Cell

var sentence string
var threshold = int32(3)
var filter []string = []string{
	"的", "在", "和", "了", "也", "上", "还", "是", "年", "有", "，", "。", " ", "都", "而", "我们", "我", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
}

func (s Stats) Map2Slice() CellSlice {
	cellSlice := make(CellSlice, 0)
	for k, v := range s {
		filtered := false
		for _, it := range filter {
			if it == k {
				filtered = true
			}
		}
		if filtered {
			continue
		}
		if "\n\t" == k {
			continue
		}
		// 排除单个汉字 或 非汉字
		if 3 >= len(k) {
			continue
		}
		// 非汉字
		rege := regexp.MustCompile(`[\P{Han}]+`)
		index := rege.FindIndex([]byte(k))
		if 0 < len(index) {
			continue
		}
		r := []rune(k)
		if 13 == r[0] && 10 == r[1] {
			continue
		}
		cell := Cell{k, v}
		cellSlice = append(cellSlice, &cell)
	}
	return cellSlice
}

func (c CellSlice) Len() int {
	return len(c)
}

func (c CellSlice) Less(i, j int) bool {
	return c[i].Freq < c[j].Freq
}

func (c CellSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c CellSlice) String() {
	for i, v := range c {
		if threshold >= v.Freq {
			continue
		}
		fmt.Printf("%d\t %v\n", i, v)
	}
}

func (c CellSlice) OutFreqAndWeight() {
	stringSaveChan := make(chan string, 5)
	go utils.SaveString(stringSaveChan, "stat.txt")
	topicmap2 := make(map[string]*Topic, 1)
	for _, v := range TopicSet {
		topicmap2[v.Const] = v
	}
	for _, v := range c {
		if threshold >= v.Freq {
			continue
		}
		strStat := fmt.Sprintf("%v\t%v\t%v\t%v", v.Freq, topicmap2[v.Word].Weight, topicmap2[v.Word].Const, topicmap2[v.Word].Relate)
		stringSaveChan <- strStat
		fmt.Println(strStat)
	}
	fmt.Println("the end.")
}
