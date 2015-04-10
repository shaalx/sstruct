package service

import (
	"fmt"
	"github.com/shaalx/sstruct/persistence"
	"github.com/shaalx/sstruct/persistence/mgodb"
	"github.com/shaalx/sstruct/pkg3/mgo/bson"
	"github.com/shaalx/sstruct/service/fetch"
	. "github.com/shaalx/sstruct/vars"
	// "strings"
	// "github.com/shaalx/sstruct/service/log"
	. "github.com/shaalx/sstruct/bean"
	"github.com/shaalx/sstruct/service/search"
	"github.com/shaalx/sstruct/utils"
	"sort"
	// "time"
	"strings"
)

type TopicAction struct {
	persis persistence.MgoPersistence
}

var (
	TopicServer       = []string{"", "sstruct", "topic"}
	stringSaveChan    chan string
	TopicMatrix       TopicMatix
	OriginTopicMatrix TopicMatix
	// posStop           = "|o|p|u|e|c|"
)

func (self *TopicAction) Init() {
	self.persis.MgoDB = mgodb.SetLocalDB(TopicServer...)
	OriginTopicMatrix = make(TopicMatix, 0)
}

func (self *TopicAction) Log(date int64) {
	utils.AppendFile(LOG_FILE, utils.UnixDateString(date)+"\t"+CURRENT_FILENAME)
}

func (self *TopicAction) Persistence() {
	CURRENT_FILENAME = "testStat.txt"
	// stringChan := utils.ReadAll(ORIGIN_DIR + CURRENT_FILENAME)
	stringSaveChan = make(chan string, 5)
	go utils.SaveStringAppend(stringSaveChan, "./"+CURRENT_FILENAME)
	i := 1
	for {
		sentence := "模块扇入是什么。"
		// sentence := <-stringChan
		if sentence == "end" {
			break
		}
		url := `http://ltpapi.voicecloud.cn/analysis/?api_key=YourApiKey&text=` + sentence + `&format=json`
		ipaddr := "202.120.87.152"
		bs := fetch.Do1(url, ipaddr)
		// fmt.Println(string(bs))
		// self.persis.Do(bs, sentence)
		fmt.Println(i)
		self.analyse(sentence, bs)
		i++
		break
	}
}

func (self *TopicAction) PersistenceWithUnixDate(date int64) {
	stringChan := utils.ReadAll(ORIGIN_DIR + CURRENT_FILENAME)
	i := 1
	for {
		// sentence := "人工智能技术在最近几年突然一下开始有了实质性的应用。"
		sentence := <-stringChan
		if sentence == "end" {
			break
		}
		url := `http://ltpapi.voicecloud.cn/analysis/?api_key=YourApiKey&text=` + sentence + `&format=json`
		ipaddr := "202.120.87.152"
		bs := fetch.Do1(url, ipaddr)
		self.persis.DoWithUnixDate(bs, sentence, date)
		fmt.Println(i)
		i++
	}
	// <-make(chan bool, 1)
}

func (self *TopicAction) AnalyseWithUnixDate(date int64) {
	selector := bson.M{"unixdate": date}
	newses := self.persis.QueryNewses(selector)
	stringSaveChan = make(chan string, 5)
	TopicMatrix = make(TopicMatix, 0)
	go utils.SaveString(stringSaveChan, SPILT_DIR+CURRENT_FILENAME)
	for _, it := range newses {
		bsfirst := utils.I2Bytes(it.Content)
		self.analyse(it.Notice, bsfirst)
	}
	// TopicMatrix.Statistics()
	TopicMatrix.StatisticsWithOrigin(&OriginTopicMatrix)
}

// 将分词结果输出到文件中
func (self *TopicAction) Analyse(n int) {
	newses := self.persis.QuerySortedLimitNNewses(nil, n, "-unixdate")
	stringSaveChan = make(chan string, 5)
	TopicMatrix = make(TopicMatix, 0)
	go utils.SaveString(stringSaveChan, SPILT_DIR+CURRENT_FILENAME)
	for _, it := range newses {
		bsfirst := utils.I2Bytes(it.Content)
		self.analyse(it.Notice, bsfirst)
	}
	// TopicMatrix.Statistics()
	// FirstStep()
}

func (self *TopicAction) Search() {
	stringChan := utils.ReadAll(ORIGIN_DIR + CURRENT_FILENAME)
	stringSaveChan = make(chan string, 5)
	TopicMatrix = make(TopicMatix, 0)
	i := 1
	go utils.SaveString(stringSaveChan, SPILT_DIR+CURRENT_FILENAME)
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
	if len(contentArrayOfFirstLayer) <= 0 {
		return
	}
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
			pos := search.SearchSValue(utils.I2Bytes(its), "pos", []string{}...)
			// if strings.Contains(posStop, "|"+pos+"|") {
			// 	fmt.Println(pos, cont)
			// }
			cont = strings.TrimSpace(cont)
			topic := Topic{id, cont, relate, parent, 0.0, 0, pos}
			topics[i] = &topic
		}
		fmt.Print(".")
		// fmt.Println(sentence)
		stringSaveChan <- sentence + "\n"
		stringSaveChan <- processSentence(topics)
		OriginTopicMatrix = append(OriginTopicMatrix, topics)
	}
}

// 处理句子成分
func processSentence(topicsOrigin TopicSlice) string {
	topicsStrOrigin := ""
	// fmt.Println(topicsOrigin.String())
	// sort.Sort(topicsOrigin)
	var hedTopic *Topic
	for _, it := range topicsOrigin {
		if (*it).IsPicked(-2, "HED") {
			hedTopic = it
		}
		if it.Relate == "WP" {
			continue
		}
		topicsStrOrigin += it.CutWords() + "\t"
	}
	topics := make(TopicSlice, 0)
	hedTopic.WeightUp(0.3)
	// topics = append(topics, hedTopic)

	// inTopics := make(TopicSlice, 0)
	for _, v := range topicsOrigin {
		// if v.IsPicked(id, []string{"SBV", "VOB"}...) {
		// if v.IsPicked(-2, []string{"SBV", "VOB", "COO", "CMP"}...) {
		// // ATT --> ATT --> ... --> ?
		// 		if att.IsPicked(-2, []string{"HED", "SBV", "ADV", "POB"}...) {
		// TopicMatrix = append(TopicMatrix, topicsOrigin)
		if v.IsPicked(-2, []string{"SBV", "ATT", "VOB", "HED", "POB", "ADV", "COO", "RAD", "LAD", "FOB"}...) {
			// inTopics = append(inTopics, v)
			TopicMatrix = append(TopicMatrix, TopicSlice{v})
		}
		/*if ok, resultTopics := v.IsCond(topicsOrigin, []string{"ATT", "ATT", "SBV"}...); ok {
			topics = append(topics, resultTopics...)
			for _, it := range topics {
				it.WeightUp(0.2)
			}
			sort.Sort(topics)
			TopicMatrix = append(TopicMatrix, topics)
			// fmt.Println(topics.String())
			topics = make(TopicSlice, 0)
		}
		if ok, resultTopics := v.IsCond(topicsOrigin, []string{"ATT", "SBV"}...); ok {
			topics = append(topics, resultTopics...)
			for _, it := range topics {
				it.WeightUp(0.2)
			}
			sort.Sort(topics)
			TopicMatrix = append(TopicMatrix, topics)
			// fmt.Println(topics.String())
			topics = make(TopicSlice, 0)
		}
		if ok, resultTopics := v.IsCond(topicsOrigin, []string{"ADV", "ATT"}...); ok {
			topics = append(topics, resultTopics...)
			for _, it := range topics {
				it.WeightUp(0.2)
			}
			sort.Sort(topics)
			TopicMatrix = append(TopicMatrix, topics)
			// fmt.Println(topics.String())
			topics = make(TopicSlice, 0)
		}
		if ok, resultTopics := v.IsCond(topicsOrigin, []string{"ATT", "HED"}...); ok {
			topics = append(topics, resultTopics...)
			for _, it := range topics {
				it.WeightUp(0.2)
			}
			sort.Sort(topics)
			TopicMatrix = append(TopicMatrix, topics)
			// fmt.Println(topics.String())
			topics = make(TopicSlice, 0)
		}
		// if ok, resultTopics := v.IsCond(topicsOrigin, []string{"VOB", "HED"}...); ok {
		// 	topics = append(topics, resultTopics...)
		// 	for _, it := range topics {
		// 		it.WeightUp(0.2)
		// 	}
		// 	sort.Sort(topics)		//
		// 	TopicMatrix = append(TopicMatrix, topics)
		//	// fmt.Println(topics.String())
		// 	topics = make(TopicSlice, 0)
		// }
		if ok, resultTopics := v.IsCond(topicsOrigin, []string{"SBV", "VOB"}...); ok {
			topics = append(topics, resultTopics...)
			for _, it := range topics {
				it.WeightUp(0.1)
			}
			sort.Sort(topics)
			TopicMatrix = append(TopicMatrix, topics)
			// fmt.Println(topics.String())
			topics = make(TopicSlice, 0)
		}
		if ok, resultTopics := v.IsCond(topicsOrigin, []string{"ATT", "VOB"}...); ok {
			topics = append(topics, resultTopics...)
			for _, it := range topics {
				it.WeightUp(0.1)
			}
			sort.Sort(topics)
			TopicMatrix = append(TopicMatrix, topics)
			// fmt.Println(topics.String())
			topics = make(TopicSlice, 0)
		}
		if ok, resultTopics := v.IsCond(topicsOrigin, []string{"COO", "VOB"}...); ok {
			topics = append(topics, resultTopics...)
			for _, it := range topics {
				it.WeightUp(0.1)
			}
			sort.Sort(topics)
			TopicMatrix = append(TopicMatrix, topics)
			// fmt.Println(topics.String())
			topics = make(TopicSlice, 0)
		}
		if ok, resultTopics := v.IsCond(topicsOrigin, []string{"VOB", "COO"}...); ok {
			topics = append(topics, resultTopics...)
			for _, it := range topics {
				it.WeightUp(0.1)
			}
			sort.Sort(topics)
			TopicMatrix = append(TopicMatrix, topics)
			// fmt.Println(topics.String())
			topics = make(TopicSlice, 0)
		}
		if ok, resultTopics := v.IsCond(topicsOrigin, []string{"ATT", "ATT"}...); ok {
			topics = append(topics, resultTopics...)
			for _, it := range topics {
				it.WeightUp(0.1)
			}
			sort.Sort(topics)
			TopicMatrix = append(TopicMatrix, topics)
			// fmt.Println(topics.String())
			topics = make(TopicSlice, 0)
		}
		if ok, resultTopics := v.IsCond(topicsOrigin, []string{"SBV", "HED"}...); ok {
			topics = append(topics, resultTopics...)
			for _, it := range topics {
				it.WeightUp(0.1)
			}
			sort.Sort(topics)
			TopicMatrix = append(TopicMatrix, topics)
			// fmt.Println(topics.String())
			topics = make(TopicSlice, 0)
		}
		if ok, resultTopics := v.IsCond(topicsOrigin, []string{"ATT", "ATT", "POB"}...); ok {
			topics = append(topics, resultTopics...)
			for _, it := range topics {
				it.WeightUp(0.1)
			}
			sort.Sort(topics)
			TopicMatrix = append(TopicMatrix, topics)
			// fmt.Println(topics.String())
			topics = make(TopicSlice, 0)
		}
		if ok, resultTopics := v.IsCond(topicsOrigin, []string{"ATT", "POB"}...); ok {
			topics = append(topics, resultTopics...)
			for _, it := range topics {
				it.WeightUp(0.1)
			}
			sort.Sort(topics)
			TopicMatrix = append(TopicMatrix, topics)
			// fmt.Println(topics.String())
			topics = make(TopicSlice, 0)
		}*/
		if ok, resultTopics := v.IsCond(topicsOrigin, []string{"ATT", "ATT", "ATT"}...); ok {
			topics = append(topics, resultTopics...)
			for _, it := range topics {
				it.WeightUp(0.1)
			}
			sort.Sort(topics)
			TopicMatrix = append(TopicMatrix, topics)
			// fmt.Println(topics.String())
			topics = make(TopicSlice, 0)
		}
		if ok, resultTopics := v.IsCond(topicsOrigin, []string{"ATT", "ATT", "SBV"}...); ok {
			topics = append(topics, resultTopics...)
			for _, it := range topics {
				it.WeightUp(0.1)
			}
			sort.Sort(topics)
			TopicMatrix = append(TopicMatrix, topics)
			// fmt.Println(topics.String())
			topics = make(TopicSlice, 0)
		}
		if ok, resultTopics := v.IsCond(topicsOrigin, []string{"ATT", "SBV"}...); ok {
			topics = append(topics, resultTopics...)
			for _, it := range topics {
				it.WeightUp(0.1)
			}
			sort.Sort(topics)
			TopicMatrix = append(TopicMatrix, topics)
			// fmt.Println(topics.String())
			topics = make(TopicSlice, 0)
		}
		if ok, resultTopics := v.IsCond(topicsOrigin, []string{"ATT", "VOB"}...); ok {
			topics = append(topics, resultTopics...)
			for _, it := range topics {
				it.WeightUp(0.1)
			}
			sort.Sort(topics)
			TopicMatrix = append(TopicMatrix, topics)
			// fmt.Println(topics.String())
			topics = make(TopicSlice, 0)
		}
		if ok, resultTopics := v.IsCond(topicsOrigin, []string{"ATT", "HED"}...); ok {
			topics = append(topics, resultTopics...)
			for _, it := range topics {
				it.WeightUp(0.1)
			}
			sort.Sort(topics)
			TopicMatrix = append(TopicMatrix, topics)
			// fmt.Println(topics.String())
			topics = make(TopicSlice, 0)
		}
		if ok, resultTopics := v.IsCond(topicsOrigin, []string{"ATT", "POB"}...); ok {
			topics = append(topics, resultTopics...)
			for _, it := range topics {
				it.WeightUp(0.1)
			}
			sort.Sort(topics)
			TopicMatrix = append(TopicMatrix, topics)
			// fmt.Println(topics.String())
			topics = make(TopicSlice, 0)
		}
		if ok, resultTopics := v.IsCond(topicsOrigin, []string{"ADV", "ATT"}...); ok {
			topics = append(topics, resultTopics...)
			for _, it := range topics {
				it.WeightUp(0.1)
			}
			sort.Sort(topics)
			TopicMatrix = append(TopicMatrix, topics)
			// fmt.Println(topics.String())
			topics = make(TopicSlice, 0)
		}
		if ok, resultTopics := v.IsCond(topicsOrigin, []string{"FOB", "VOB"}...); ok {
			topics = append(topics, resultTopics...)
			for _, it := range topics {
				it.WeightUp(0.1)
			}
			sort.Sort(topics)
			TopicMatrix = append(TopicMatrix, topics)
			// fmt.Println(topics.String())
			topics = make(TopicSlice, 0)
		}
		if ok, resultTopics := v.IsCond(topicsOrigin, []string{"SBV", "SBV"}...); ok {
			topics = append(topics, resultTopics...)
			for _, it := range topics {
				it.WeightUp(0.1)
			}
			sort.Sort(topics)
			TopicMatrix = append(TopicMatrix, topics)
			// fmt.Println(topics.String())
			topics = make(TopicSlice, 0)
		}
		if ok, resultTopics := v.IsCond(topicsOrigin, []string{"SBV", "ATT"}...); ok {
			topics = append(topics, resultTopics...)
			for _, it := range topics {
				it.WeightUp(0.1)
			}
			sort.Sort(topics)
			TopicMatrix = append(TopicMatrix, topics)
			// fmt.Println(topics.String())
			topics = make(TopicSlice, 0)
		}
		if ok, resultTopics := v.IsCond(topicsOrigin, []string{"SBV", "VOB"}...); ok {
			topics = append(topics, resultTopics...)
			for _, it := range topics {
				it.WeightUp(0.1)
			}
			sort.Sort(topics)
			TopicMatrix = append(TopicMatrix, topics)
			// fmt.Println(topics.String())
			topics = make(TopicSlice, 0)
		}
		if ok, resultTopics := v.IsCond(topicsOrigin, []string{"SBV", "COO"}...); ok {
			topics = append(topics, resultTopics...)
			for _, it := range topics {
				it.WeightUp(0.1)
			}
			sort.Sort(topics)
			TopicMatrix = append(TopicMatrix, topics)
			// fmt.Println(topics.String())
			topics = make(TopicSlice, 0)
		}
		// ATT-POB
		if ok, resultTopics := v.IsCond(topicsOrigin, []string{"SBV", "HED"}...); ok {
			topics = append(topics, resultTopics...)
			for _, it := range topics {
				it.WeightUp(0.1)
			}
			sort.Sort(topics)
			TopicMatrix = append(TopicMatrix, topics)
			// fmt.Println(topics.String())
			topics = make(TopicSlice, 0)
		}
	}
	// inTopics = *inTopics.EjRepeat()
	// sort.Sort(inTopics)
	// TopicMatrix = append(TopicMatrix, inTopics)

	// sort.Sort(topics)
	// result := ""
	// topicsStr := ""
	// for _, it := range topics {
	// 	result += it.Const
	// 	topicsStr += it.String()
	// }
	// fmt.Printf("%s\n%s\n", topicsStrOrigin, topicsStr)
	return /*result + "\n" + */ topicsStrOrigin /*+ "\n" + topicsStr + "\n"*/
	// return topicsOrigin.String() + "\n"
}

func (self *TopicAction) Close() {
	self.persis.MgoDB.Close()
}
