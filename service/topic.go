package service

import (
	"fmt"
	"github.com/shaalx/sstruct/persistence"
	"github.com/shaalx/sstruct/persistence/mgodb"
	"github.com/shaalx/sstruct/service/fetch"
	// "strings"
	// "github.com/shaalx/sstruct/service/log"
	. "github.com/shaalx/sstruct/bean"
	"github.com/shaalx/sstruct/service/search"
	"github.com/shaalx/sstruct/utils"
	"sort"
	// "time"
	// "strings"
)

type TopicAction struct {
	persis persistence.MgoPersistence
}

var (
	TopicServer    = []string{"", "sstruct", "topic"}
	stringSaveChan chan string
	TopicMatrix    TopicMatix
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
	newses := self.persis.QuerySortedLimitNNewses(nil, 22, "-unixdate")
	stringSaveChan = make(chan string, 5)
	TopicSet = make(TopicSlice, 0)
	TopicMatrix = make(TopicMatix, 0)
	go utils.SaveString(stringSaveChan, "result.txt")
	for _, it := range newses {
		bsfirst := utils.I2Bytes(it.Content)
		self.analyse(it.Notice, bsfirst)
	}
	TopicMatrix.Statistics()
	// FirstStep()
}

func (self *TopicAction) Search() {
	stringChan := utils.ReadAll("file.txt")
	stringSaveChan = make(chan string, 5)
	TopicSet = make(TopicSlice, 0)
	TopicMatrix = make(TopicMatix, 0)
	i := 1
	go utils.SaveString(stringSaveChan, "result.txt")
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
		fmt.Println(i)
		i++
	}
	TopicMatrix.Statistics()
	// FirstStep()
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
		fmt.Println(sentence)
		stringSaveChan <- sentence
		stringSaveChan <- processSentence(topics)
	}
}

// 处理句子成分
func processSentence(topicsOrigin TopicSlice) string {
	topicsStrOrigin := ""
	fmt.Println(topicsOrigin.String())
	// sort.Sort(topicsOrigin)
	var hedTopic *Topic
	for _, it := range topicsOrigin {
		if (*it).IsPicked(-2, "HED") {
			hedTopic = it
		}
		topicsStrOrigin += it.String()
	}
	topics := make(TopicSlice, 0)
	hedTopic.WeightUp(0.3)
	// topics = append(topics, hedTopic)

	inTopics := make(TopicSlice, 0)
	for _, v := range topicsOrigin {
		// if v.IsPicked(id, []string{"SBV", "VOB"}...) {
		// if v.IsPicked(-2, []string{"SBV", "VOB", "COO", "CMP"}...) {
		// // ATT --> ATT --> ... --> ?
		// 		if att.IsPicked(-2, []string{"HED", "SBV", "ADV", "POB"}...) {
		if v.IsPicked(-2, []string{"SBV", "ATT", "VOB", "HED"}...) {
			inTopics = append(inTopics, v)
			TopicMatrix = append(TopicMatrix, TopicSlice{v})
			TopicSet = append(TopicSet, v)
		}
		if ok, resultTopics := v.IsCond(topicsOrigin, []string{"ATT", "ATT", "SBV"}...); ok {
			topics = append(topics, resultTopics...)
			for _, it := range topics {
				it.WeightUp(0.2)
			}
			sort.Sort(topics)
			inTopics = append(inTopics, topics...)
			TopicMatrix = append(TopicMatrix, topics)
			TopicSet = append(TopicSet, topics...)
			fmt.Println(topics.String())
			topics = make(TopicSlice, 0)
		}
		if ok, resultTopics := v.IsCond(topicsOrigin, []string{"ATT", "SBV"}...); ok {
			topics = append(topics, resultTopics...)
			for _, it := range topics {
				it.WeightUp(0.2)
			}
			sort.Sort(topics)
			inTopics = append(inTopics, topics...)
			TopicMatrix = append(TopicMatrix, topics)
			TopicSet = append(TopicSet, topics...)
			fmt.Println(topics.String())
			topics = make(TopicSlice, 0)
		}
		if ok, resultTopics := v.IsCond(topicsOrigin, []string{"ATT", "HED"}...); ok {
			topics = append(topics, resultTopics...)
			for _, it := range topics {
				it.WeightUp(0.2)
			}
			sort.Sort(topics)
			inTopics = append(inTopics, topics...)
			TopicMatrix = append(TopicMatrix, topics)
			TopicSet = append(TopicSet, topics...)
			fmt.Println(topics.String())
			topics = make(TopicSlice, 0)
		}
		// if ok, resultTopics := v.IsCond(topicsOrigin, []string{"VOB", "HED"}...); ok {
		// 	topics = append(topics, resultTopics...)
		// 	for _, it := range topics {
		// 		it.WeightUp(0.2)
		// 	}
		// 	sort.Sort(topics)
		// 	inTopics = append(inTopics, topics...)
		// 	TopicMatrix = append(TopicMatrix, topics)
		// 	TopicSet = append(TopicSet, topics...)
		// 	fmt.Println(topics.String())
		// 	topics = make(TopicSlice, 0)
		// }
		if ok, resultTopics := v.IsCond(topicsOrigin, []string{"SBV", "VOB"}...); ok {
			topics = append(topics, resultTopics...)
			for _, it := range topics {
				it.WeightUp(0.1)
			}
			sort.Sort(topics)
			inTopics = append(inTopics, topics...)
			TopicMatrix = append(TopicMatrix, topics)
			TopicSet = append(TopicSet, topics...)
			fmt.Println(topics.String())
			topics = make(TopicSlice, 0)
		}
		if ok, resultTopics := v.IsCond(topicsOrigin, []string{"ATT", "VOB"}...); ok {
			topics = append(topics, resultTopics...)
			for _, it := range topics {
				it.WeightUp(0.1)
			}
			sort.Sort(topics)
			inTopics = append(inTopics, topics...)
			TopicMatrix = append(TopicMatrix, topics)
			TopicSet = append(TopicSet, topics...)
			fmt.Println(topics.String())
			topics = make(TopicSlice, 0)
		}
		if ok, resultTopics := v.IsCond(topicsOrigin, []string{"COO", "VOB"}...); ok {
			topics = append(topics, resultTopics...)
			for _, it := range topics {
				it.WeightUp(0.1)
			}
			sort.Sort(topics)
			inTopics = append(inTopics, topics...)
			TopicMatrix = append(TopicMatrix, topics)
			TopicSet = append(TopicSet, topics...)
			fmt.Println(topics.String())
			topics = make(TopicSlice, 0)
		}
		if ok, resultTopics := v.IsCond(topicsOrigin, []string{"VOB", "COO"}...); ok {
			topics = append(topics, resultTopics...)
			for _, it := range topics {
				it.WeightUp(0.1)
			}
			sort.Sort(topics)
			inTopics = append(inTopics, topics...)
			TopicMatrix = append(TopicMatrix, topics)
			TopicSet = append(TopicSet, topics...)
			fmt.Println(topics.String())
			topics = make(TopicSlice, 0)
		}
		if ok, resultTopics := v.IsCond(topicsOrigin, []string{"ATT", "ATT"}...); ok {
			topics = append(topics, resultTopics...)
			for _, it := range topics {
				it.WeightUp(0.1)
			}
			sort.Sort(topics)
			inTopics = append(inTopics, topics...)
			TopicMatrix = append(TopicMatrix, topics)
			TopicSet = append(TopicSet, topics...)
			fmt.Println(topics.String())
			topics = make(TopicSlice, 0)
		}
		if ok, resultTopics := v.IsCond(topicsOrigin, []string{"SBV", "HED"}...); ok {
			topics = append(topics, resultTopics...)
			for _, it := range topics {
				it.WeightUp(0.1)
			}
			sort.Sort(topics)
			inTopics = append(inTopics, topics...)
			TopicMatrix = append(TopicMatrix, topics)
			TopicSet = append(TopicSet, topics...)
			fmt.Println(topics.String())
			topics = make(TopicSlice, 0)
		}
		if ok, resultTopics := v.IsCond(topicsOrigin, []string{"ATT", "ATT", "POB"}...); ok {
			topics = append(topics, resultTopics...)
			for _, it := range topics {
				it.WeightUp(0.1)
			}
			sort.Sort(topics)
			inTopics = append(inTopics, topics...)
			TopicMatrix = append(TopicMatrix, topics)
			TopicSet = append(TopicSet, topics...)
			fmt.Println(topics.String())
			topics = make(TopicSlice, 0)
		}
		if ok, resultTopics := v.IsCond(topicsOrigin, []string{"ATT", "POB"}...); ok {
			topics = append(topics, resultTopics...)
			for _, it := range topics {
				it.WeightUp(0.1)
			}
			sort.Sort(topics)
			inTopics = append(inTopics, topics...)
			TopicMatrix = append(TopicMatrix, topics)
			TopicSet = append(TopicSet, topics...)
			fmt.Println(topics.String())
			topics = make(TopicSlice, 0)
		}
	}
	// inTopics = *inTopics.EjRepeat()
	// sort.Sort(inTopics)
	// TopicSet = append(TopicSet, inTopics...)
	// TopicMatrix = append(TopicMatrix, inTopics)

	// sort.Sort(topics)
	// result := ""
	// topicsStr := ""
	// for _, it := range topics {
	// 	result += it.Const
	// 	topicsStr += it.String()
	// }
	// fmt.Printf("%s\n%s\n", topicsStrOrigin, topicsStr)
	return /*result + "\n" + */ topicsStrOrigin + "\n" /*+ topicsStr + "\n"*/
}

func (self *TopicAction) Close() {
	self.persis.MgoDB.Close()
}

// 统计

func FirstStep() {
	freq := Stating()
	cells := freq.Map2Slice()
	sort.Sort(sort.Reverse(cells))
	// cells.String()
	cells.OutFreqAndWeight()
	TopicMatrix.Statistics()
}

func Stating() Stats {
	stats := make(Stats, 1)
	for _, it := range TopicSet {
		stats[it.Const]++
	}
	return stats
}
