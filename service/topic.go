package service

import (
	"fmt"
	"github.com/shaalx/sstruct/persistence"
	"github.com/shaalx/sstruct/persistence/mgodb"
	"github.com/shaalx/sstruct/service/fetch"
	// "strings"
	// "github.com/shaalx/sstruct/service/log"
	// . "github.com/shaalx/sstruct/bean"
	"github.com/shaalx/sstruct/service/search"
	"github.com/shaalx/sstruct/utils"
	"sort"
	// "time"
	"regexp"
	// "strings"
)

type TopicAction struct {
	persis persistence.MgoPersistence
}

var (
	TopicServer    = []string{"", "sstruct", "topic"}
	stringSaveChan chan string
	TopicSet       TopicSlice
)

func (self *TopicAction) Init() {
	self.persis.MgoDB = mgodb.SetLocalDB(TopicServer...)
}

func (self *TopicAction) Persistence() {
	url := `http://ltpapi.voicecloud.cn/analysis/?api_key=YourApiKey&text=盗梦空间是一部好电影。大家对它的评价非常的高。&format=json`
	ipaddr := "202.120.87.152"
	bs := fetch.Do1(url, ipaddr)
	self.persis.Do(bs, "")
}

func (self *TopicAction) QueryOne() {
	one := self.persis.QueryOne()
	fmt.Println(one)
	bs := utils.I2Bytes(one)
	fmt.Println(string(bs))
}

func (self *TopicAction) Analyse() {
	one := self.persis.QuerySortedNewsOne(nil, "-unixdate")

	stringSaveChan = make(chan string, 5)
	TopicSet = make(TopicSlice, 0)
	go utils.SaveString(stringSaveChan)

	bsfirst := utils.I2Bytes(one.Content)
	self.analyse(one.Notice, bsfirst)
	FirstStep()
}

func (self *TopicAction) Search() {
	stringChan := utils.ReadAll("file.txt")
	stringSaveChan = make(chan string, 5)
	TopicSet = make(TopicSlice, 0)
	go utils.SaveString(stringSaveChan)
	for {
		// sentence := "人工智能技术在最近几年突然一下开始有了实质性的应用。"
		sentence := <-stringChan
		if sentence == "end" {
			break
		}
		url := `http://ltpapi.voicecloud.cn/analysis/?api_key=YourApiKey&text=` + sentence + `&format=json`
		ipaddr := "202.120.87.152"
		bs := fetch.Do1(url, ipaddr)
		self.persis.Do(bs, sentence)
		self.analyse(sentence, bs)
		// break
	}
	FirstStep()
	a := make(chan bool, 1)
	<-a
}

func (self *TopicAction) analyse(sentence string, data []byte) {
	// 第一层数组
	contentArrayOfFirstLayer := search.SearchArrays(data, []string{}...)
	// fmt.Println(contentArrayOfFirstLayer)

	// 第二层数组
	bssecond := utils.I2Bytes(contentArrayOfFirstLayer[0])
	contentArrayOfSecondLayer := search.SearchArrays(bssecond, []string{}...)
	// fmt.Println(contentArrayOfSecondLayer)

	for _, it := range contentArrayOfSecondLayer {
		// fmt.Println(it)
		arybs := utils.I2Bytes(it)
		ary := search.SearchArrays(arybs, []string{}...)
		// fmt.Println(ary)
		topics := make(TopicSlice, len(ary))
		for i, its := range ary {
			id := search.SearchFIValue(utils.I2Bytes(its), "id", []string{}...)
			cont := search.SearchSValue(utils.I2Bytes(its), "cont", []string{}...)
			relate := search.SearchSValue(utils.I2Bytes(its), "relate", []string{}...)
			parent := search.SearchFIValue(utils.I2Bytes(its), "parent", []string{}...)
			topic := Topic{id, cont, relate, parent, 0.0, 0}
			topics[i] = &topic
		}
		stringSaveChan <- sentence
		stringSaveChan <- processSentence(topics)
	}
}

// 处理句子成分
func processSentence(topicsOrigin TopicSlice) string {
	topicsStrOrigin := ""
	// sort.Sort(topicsOrigin)
	var hedTopic *Topic
	for _, it := range topicsOrigin {
		if (*it).isPicked(-2, "HED") {
			hedTopic = it
		}
		topicsStrOrigin += it.String()
	}
	topics := make(TopicSlice, 0)
	id := hedTopic.Id
	topics = append(topics, hedTopic)
	for _, v := range topicsOrigin {
		// if v.Parent == id {
		// 	topics = append(topics, v)
		// }
		// 句子核心句法成分
		if topics.Contain(v) {
			continue
		}
		if v.isPicked(id, []string{"SBV", "VOB"}...) {
			v.WeightUp(0.4)
			topics = append(topics, v)
		}
		// 其他关键字
		if topics.Contain(v) {
			continue
		}
		if v.isPicked(-2, []string{"SBV", "VOB", "COO", "CMP"}...) {
			v.WeightUp(0.2)
			topics = append(topics, v)
		}
		// 专有名词做定语
		if v.isPicked(-2, "ATT") && 3 <= len(v.Const) {
			v.WeightUp(0.5)
			topics = append(topics, v)
		}
		// 提取定语：SBV，HED的定语ATT
		// // ATT --> SBV
		// if topicsOrigin[v.Parent].isPicked(-2, "SBV") && v.isPicked(-2, "ATT") {
		// 	topics = append(topics, v)
		// }
		// ATT --> ATT --> ... --> ?
		att := v
		atts := make(TopicSlice, 0)
		for {
			if att.isPicked(-2, "ATT") {
				atts = append(atts, att)
				att.WeightUp(0.2)
			} else {
				if att.isPicked(-2, []string{"HED", "SBV", "ADV", "POB"}...) {
					if 1 >= len(atts) {
						break
					}
					att.WeightUp(0.1 * (float32)(len(atts)))
					atts = append(atts, att)
					for _, it := range atts {
						it.WeightUp(att.Weight)
					}
					topics = append(topics, atts...)
				} else {
					if 1 < len(atts) {
						for _, it := range atts {
							it.WeightUp(0.1)
						}
						topics = append(topics, atts...)
					}
				}
				break
			}
			if -1 == att.Parent {
				break
			}
			att = topicsOrigin[att.Parent]
		}
	}
	topics = *topics.EjRepeat()
	TopicSet = append(TopicSet, topics...)
	sort.Sort(topics)
	result := ""
	topicsStr := ""
	for _, it := range topics {
		result += it.Const
		topicsStr += it.String()
	}
	fmt.Printf("%s\n%s\n", topicsStrOrigin, topicsStr)
	return result + "\n" + topicsStrOrigin + "\n" + topicsStr + "\n"
}

func (self *TopicAction) Close() {
	self.persis.MgoDB.Close()
}

type Topic struct {
	Id     int64
	Const  string
	Relate string
	Parent int64
	Weight float32
	Freq   int32
}

type TopicMap map[int64]Topic

// 句法成分是否为指定条件
func (t Topic) isPicked(parent int64, relate ...string) bool {
	if -2 == parent {
		goto goodDaddy
	}
	if parent != t.Parent {
		return false
	}
goodDaddy:
	for _, it := range relate {
		if t.Relate == it {
			return true
		}
	}
	return false
}

// 增加权重
func (t *Topic) WeightUp(w float32) {
	t.Weight += w
}

func (t Topic) String() string {
	// tStr := fmt.Sprintf("Id %d\t ,Const %s\t ,Relate %s ,Parent %d\n", t.Id, t.Const, t.Relate, t.Parent)
	tStr := fmt.Sprintf("%d\t %s\t %d\t %s\t%.3f\n", t.Id, t.Relate, t.Parent, t.Const, t.Weight)
	return tStr
}

// 排序
type TopicSlice []*Topic

func (c TopicSlice) Len() int {
	return len(c)
}

func (c TopicSlice) Less(i, j int) bool {
	return c[i].Id < c[j].Id
}

func (c TopicSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (t *TopicSlice) Contain(topic *Topic) bool {
	for _, it := range *t {
		if it.Id == topic.Id {
			return true
		}
	}
	return false
}

// 排除重复值
func (t *TopicSlice) EjRepeat() *TopicSlice {
	length := len(*t)
	ejMap := make(map[int64]*Topic, length)
	for _, it := range *t {
		ejMap[it.Id] = it
	}
	i := 0
	result := make(TopicSlice, len(ejMap))
	for _, v := range ejMap {
		result[i] = v
		i++
	}
	return &result
}

// 统计

func FirstStep() {
	freq := Stating()
	cells := freq.Map2Slice()
	sort.Sort(sort.Reverse(cells))
	// cells.String()
	cells.OutFreqAndWeight()
}

type Stats map[string]int32

type Cell struct {
	Word string
	Freq int32
}

type CellSlice []*Cell

var sentence string
var threshold = int32(0)
var filter []string = []string{
	"的", "在", "和", "了", "也", "上", "还", "是", "年", "有", "，", "。", " ", "都", "而", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
}

func Stating() Stats {
	stats := make(Stats, 1)
	for _, it := range TopicSet {
		stats[it.Const]++
	}
	return stats
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
	topicmap2 := make(map[string]*Topic, 1)
	for _, v := range TopicSet {
		topicmap2[v.Const] = v
	}
	for _, v := range c {
		if threshold >= v.Freq {
			continue
		}
		strStat := fmt.Sprintf("%v\t%v\t%v\t%v\n", v.Freq, topicmap2[v.Word].Weight, topicmap2[v.Word].Const, topicmap2[v.Word].Relate)
		stringSaveChan <- strStat
	}
	fmt.Println("the end.")
}
