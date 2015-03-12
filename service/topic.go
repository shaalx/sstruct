package service

import (
	"fmt"
	"github.com/shaalx/sstruct/persistence"
	"github.com/shaalx/sstruct/persistence/mgodb"
	"github.com/shaalx/sstruct/service/fetch"
	// "strings"
	// "github.com/shaalx/sstruct/service/log"
	// "github.com/shaalx/sstruct/bean"
	"github.com/shaalx/sstruct/service/search"
	"github.com/shaalx/sstruct/utils"
	"sort"
	// "time"
)

type Topic struct {
	Id     int64
	Const  string
	Relate string
	Parent int64
	Weight float32
}

func (t Topic) String() string {
	// tStr := fmt.Sprintf("Id %d\t ,Const %s\t ,Relate %s ,Parent %d\n", t.Id, t.Const, t.Relate, t.Parent)
	tStr := fmt.Sprintf("%d\t %s\t %d\t %s\t%.3f\n", t.Id, t.Relate, t.Parent, t.Const, t.Weight)
	return tStr
}

type TopicMap map[int64]Topic

type TopicAction struct {
	persis persistence.MgoPersistence
}

var (
	TopicServer    = []string{"", "sstruct", "topic"}
	stringSaveChan chan string
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
	go utils.SaveString(stringSaveChan)

	bsfirst := utils.I2Bytes(one.Content)
	self.analyse(one.Notice, bsfirst)
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
			topic := Topic{id, cont, relate, parent, 0.0}
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
		if it.isPicked(-2, "HED") {
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
			v.WeightUp(0.3)
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
				att.WeightUp(0.1)
			} else {
				if att.isPicked(-2, []string{"HED", "SBV", "ADV"}...) {
					att.WeightUp(0.1 * (float32)(len(atts)))
					topics = append(topics, atts...)
				}
				break
			}
			if -1 == att.Parent {
				break
			}
			att = topicsOrigin[att.Parent]
		}
	}
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

func (self *TopicAction) Search() {
	stringChan := utils.ReadAll("file.txt")
	stringSaveChan = make(chan string, 5)
	go utils.SaveString(stringSaveChan)
	for {
		// sentence := "政协委员朱维群、黄洁夫、胡晓义、李彦宏、俞敏洪谈促进民生改善与社会和谐稳定。"
		sentence := <-stringChan
		url := `http://ltpapi.voicecloud.cn/analysis/?api_key=YourApiKey&text=` + sentence + `&format=json`
		ipaddr := "202.120.87.152"
		bs := fetch.Do1(url, ipaddr)
		self.persis.Do(bs, sentence)
		self.analyse(sentence, bs)
		// break
	}
	a := make(chan bool, 1)
	<-a
}

func (self *TopicAction) Close() {
	self.persis.MgoDB.Close()
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
