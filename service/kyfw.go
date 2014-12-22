package service

import (
	"encoding/json"
	"fmt"
	. "github.com/shaalx/sstruct/bean"
	"github.com/shaalx/sstruct/persistence"
	"github.com/shaalx/sstruct/persistence/mgodb"
	"github.com/shaalx/sstruct/service/fetch"
	"github.com/shaalx/sstruct/service/log"
	"github.com/shaalx/sstruct/service/search"
	"github.com/shaalx/sstruct/utils"
)

type KYFWAction struct {
	persis persistence.MgoPersistence
}

var (
	KYFWServer = []string{"", "kyfw", "ssh"}
)

func (self *KYFWAction) Init() {
	self.persis.MgoDB = mgodb.SetLocalDB(KYFWServer...)
}

func (self *KYFWAction) Persistence() {
	url := "https://kyfw.12306.cn/otn/leftTicket/queryT?leftTicketDTO.train_date=2015-02-27&leftTicketDTO.from_station=SJP&leftTicketDTO.to_station=SHH&purpose_codes=0X00"
	ipaddr := "202.120.87.152"
	bs := fetch.Do(url, ipaddr)
	self.persis.Do(bs)
}

func (self *KYFWAction) QueryOne() {
	one := self.persis.QueryOne()
	fmt.Println(one)
	bs := utils.I2Bytes(one)
	fmt.Println(string(bs))
}

func (self *KYFWAction) Analyse() {
	one := self.persis.QuerySortedNewsOne(nil, "-unixdate")
	buf := utils.I2Bytes(one.Content)
	fmt.Println(one.DisplayDate)
	status := search.KYFWStatus(buf)
	fmt.Println(status)
	if status {
		res := search.Data(buf)
		// fmt.Println(res)
		for _, v := range res {
			KYFWShow(v)
		}
	} else {
		log.LOGS.Notice("%s", "查询失败，重试中...")
	}
}

func Searching(news News) {
	fmt.Println(news.DisplayDate)
	buf := utils.I2Bytes(news.Content)
	status := search.KYFWStatus(buf)
	fmt.Println(status)
	if status {
		res := search.Data(buf)
		// fmt.Println(res)
		for _, v := range res {
			KYFWShow(v)
		}
	} else {
		log.LOGS.Notice("%s", "查询失败，重试中...")
	}
}

func B2News(data []byte) News {
	var newsV News
	err := json.Unmarshal(data, &newsV)
	if log.IsError("{bytes --> News error}", err) {
		return newsV
	}
	return newsV
}

func (self *KYFWAction) Search() {
	url := "https://kyfw.12306.cn/otn/leftTicket/queryT?leftTicketDTO.train_date=2015-02-27&leftTicketDTO.from_station=SJP&leftTicketDTO.to_station=SHH&purpose_codes=0X00"
	ipaddr := "202.120.87.152"
	bs := fetch.Do(url, ipaddr)
	news := B2News(bs)
	Searching(news)
}

func KYFWShow(m map[string]interface{}) {
	fmt.Printf("[ %v ] ---> [ %v ] :  %v ( %v )\n", m["from_station_name"], m["to_station_name"], m["station_train_code"], m["lishi"])
	fmt.Printf("[硬座] ：%v\n", m["yz_num"])
	fmt.Printf("[软座] ：%v\n", m["rz_num"])
	fmt.Printf("[硬卧] ：%v\n", m["yw_num"])
	fmt.Println("--------------------------------------------------------\n")
}

func (self *KYFWAction) LatestNews() []map[string]interface{} {
	one := self.persis.QuerySortedNewsOne(nil, "-unixdate")
	buf := utils.I2Bytes(one.Content)
	TTShows(buf)
	return TTContent(buf)
}

func (self *KYFWAction) TTLatestNews() []Toutiao {
	one := self.persis.QuerySortedNewsOne(nil, "-unixdate")
	buf := utils.I2Bytes(one.Content)
	return TTContents(buf)
}

func (self *KYFWAction) News(i int64) interface{} {
	fmt.Println(i)
	one := self.persis.QuerySortedNewsOne(nil, "-unixdate")
	buf := utils.I2Bytes(one.Content)
	return TTContents(buf)
}
