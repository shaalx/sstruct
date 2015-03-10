package service

import (
	"fmt"
	"github.com/shaalx/sstruct/persistence"
	"github.com/shaalx/sstruct/persistence/mgodb"
	"github.com/shaalx/sstruct/service/fetch"
	// "github.com/shaalx/sstruct/service/log"
	// "github.com/shaalx/sstruct/service/search"
	"github.com/shaalx/sstruct/utils"
)

type TopicAction struct {
	persis persistence.MgoPersistence
}

var (
	TopicServer = []string{"", "sstruct", "topic"}
)

func (self *TopicAction) Init() {
	self.persis.MgoDB = mgodb.SetLocalDB(TopicServer...)
}

func (self *TopicAction) Persistence() {
	url := `http://ltpapi.voicecloud.cn/analysis/?api_key=YourApiKey&text=盗梦空间是一部好电影，大家对它的评价非常的高。&format=json`
	ipaddr := "202.120.87.152"
	bs := fetch.Do(url, ipaddr)
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
	fmt.Println(one.Content)
}

func (self *TopicAction) Search() {
	url := `http://ltpapi.voicecloud.cn/analysis/?api_key=YourApiKey&text=盗梦空间是一部好电影，` +
		`大家对它的评价非常的高。` +
		`&format=json`
	ipaddr := "202.120.87.152"
	bs := fetch.Do1(url, ipaddr)
	self.persis.Do(bs)
}

func (self *TopicAction) Close() {
	self.persis.MgoDB.Close()
}
