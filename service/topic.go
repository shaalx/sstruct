package service

import (
	"fmt"
	"github.com/shaalx/sstruct/persistence"
	"github.com/shaalx/sstruct/persistence/mgodb"
	"github.com/shaalx/sstruct/service/fetch"
	// "github.com/shaalx/sstruct/service/log"
	// "github.com/shaalx/sstruct/bean"
	"github.com/shaalx/sstruct/service/search"
	"github.com/shaalx/sstruct/utils"
	"time"
)

type Topic struct {
	Id     int64
	Const  string
	Relate string
	Parent int64
}

type TopicMap map[string]Topic

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
	self.persis.Do(bs)
}

func (self *TopicAction) QueryOne() {
	one := self.persis.QueryOne()
	fmt.Println(one)
	bs := utils.I2Bytes(one)
	fmt.Println(string(bs))
}

func (self *TopicAction) Analyse() {
	one := self.persis.QuerySortedNewsOne(nil, "-unixdate")
	// buf := utils.I2Bytes(one.Content)
	fmt.Println(one.DisplayDate)
	// fmt.Println(one.Content)

	// 第一层数组
	bsfirst := utils.I2Bytes(one.Content)
	contentArrayOfFirstLayer := search.SearchArrays(bsfirst, []string{}...)
	// fmt.Println(contentArrayOfFirstLayer)

	// 第二层数组
	bssecond := utils.I2Bytes(contentArrayOfFirstLayer[0])
	contentArrayOfSecondLayer := search.SearchArrays(bssecond, []string{}...)
	// fmt.Println(contentArrayOfSecondLayer)

	// contentArrayLength := len(contentArrayOfSecondLayer)
	TopicMaps := make(TopicMap, 15)
	for _, it := range contentArrayOfSecondLayer {
		// fmt.Println(it)
		arybs := utils.I2Bytes(it)
		ary := search.SearchArrays(arybs, []string{}...)
		// fmt.Println(ary)
		for _, its := range ary {
			// fmt.Println(its)
			id := search.SearchFIValue(utils.I2Bytes(its), "id", []string{}...)
			cont := search.SearchSValue(utils.I2Bytes(its), "cont", []string{}...)
			relate := search.SearchSValue(utils.I2Bytes(its), "relate", []string{}...)
			parent := search.SearchFIValue(utils.I2Bytes(its), "parent", []string{}...)
			topic := Topic{id, cont, relate, parent}
			// fmt.Printf("%v %s\t%s %v\n", id, cont, relate, parent)
			fmt.Printf("%v\n", topic)
			TopicMaps[topic.Relate] = topic
		}
		fmt.Println(TopicMaps["SBV"], TopicMaps["HED"], TopicMaps["VOB"])
	}
}

func (self *TopicAction) analyse(data []byte) {
	// 第一层数组
	contentArrayOfFirstLayer := search.SearchArrays(data, []string{}...)
	// fmt.Println(contentArrayOfFirstLayer)

	// 第二层数组
	bssecond := utils.I2Bytes(contentArrayOfFirstLayer[0])
	contentArrayOfSecondLayer := search.SearchArrays(bssecond, []string{}...)
	// fmt.Println(contentArrayOfSecondLayer)

	// contentArrayLength := len(contentArrayOfSecondLayer)
	TopicMaps := make(TopicMap, 25)
	for _, it := range contentArrayOfSecondLayer {
		// fmt.Println(it)
		arybs := utils.I2Bytes(it)
		ary := search.SearchArrays(arybs, []string{}...)
		// fmt.Println(ary)
		for _, its := range ary {
			// fmt.Println(its)
			id := search.SearchFIValue(utils.I2Bytes(its), "id", []string{}...)
			cont := search.SearchSValue(utils.I2Bytes(its), "cont", []string{}...)
			relate := search.SearchSValue(utils.I2Bytes(its), "relate", []string{}...)
			parent := search.SearchFIValue(utils.I2Bytes(its), "parent", []string{}...)
			topic := Topic{id, cont, relate, parent}
			// fmt.Printf("%v %s\t%s %v\n", id, cont, relate, parent)
			fmt.Printf("%v\n", topic)
			TopicMaps[topic.Relate] = topic
		}
		fmt.Println(TopicMaps["SBV"], TopicMaps["HED"], TopicMaps["VOB"])
		stringSaveChan <- TopicMaps["SBV"].Const + TopicMaps["HED"].Const + TopicMaps["VOB"].Const + TopicMaps["POB"].Const + TopicMaps["IOB"].Const + TopicMaps["FOB"].Const
	}
}

func (self *TopicAction) Search() {
	stringChan := utils.ReadAll("file.txt")
	stringSaveChan = make(chan string, 5)
	go utils.SaveString(stringSaveChan)
	for {
		// sentence := "佳洁士双效炫白牙膏被处罚603万元。这也是我国目前针对虚假违法广告的最大罚单。"
		sentence := <-stringChan
		url := `http://ltpapi.voicecloud.cn/analysis/?api_key=YourApiKey&text=` + sentence + `&format=json`
		ipaddr := "202.120.87.152"
		bs := fetch.Do1(url, ipaddr)
		self.persis.Do(bs)
		self.analyse(bs)
		time.Sleep(time.Second * 1)
	}
}

func (self *TopicAction) Close() {
	self.persis.MgoDB.Close()
}
