package service

import (
	"fmt"
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
	self.persis.Do(bs, "")
}

func (self *KYFWAction) QueryOne() {
	one := self.persis.QueryOne()
	fmt.Println(one)
	bs := utils.I2Bytes(one)
	fmt.Println(string(bs))
}

func (self *KYFWAction) Analyse(n int) {
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

func Searching(data []byte) {
	// fmt.Println(news.DisplayDate)
	// buf := utils.I2Bytes(news.Content)
	status := search.KYFWStatus(data)
	if status {
		res := search.Data(data)
		// fmt.Println(res)
		for _, v := range res {
			train_code, ok := v["station_train_code"]
			if !ok {
				continue
			}
			code_str, ok := train_code.(string)
			if !ok {
				continue
			}
			if "Z198" == code_str || "Z270" == code_str {
				KYFWShow(v)
			}
		}
	} else {
		log.LOGS.Notice("%s", "查询失败，重试中...")
	}
}

func (self *KYFWAction) Search() {
	url := "https://kyfw.12306.cn/otn/leftTicket/queryT?leftTicketDTO.train_date=2015-02-27&leftTicketDTO.from_station=SJP&leftTicketDTO.to_station=SHH&purpose_codes=0X00"
	// url := "https://kyfw.12306.cn/otn/leftTicket/queryT?leftTicketDTO.train_date=2015-03-15&leftTicketDTO.from_station=SJP&leftTicketDTO.to_station=SHH&purpose_codes=0X00"
	ipaddr := "202.120.87.152"
	bs := fetch.Do(url, ipaddr)
	Searching(bs)
	self.persis.Do(bs, "")
}

func KYFWShow(m map[string]interface{}) {
	fmt.Println("--------------------------------------------------------\n")
	fmt.Printf("%v : %v >>> %v  %v ~ %v  < %v > \n\n", m["station_train_code"], m["from_station_name"], m["to_station_name"], m["start_time"], m["arrive_time"], m["lishi"])
	fmt.Printf("[软座] ：%v\t", m["rz_num"])
	fmt.Printf("[硬座] ：%v\t", m["yz_num"])
	fmt.Printf("[硬卧] ：%v\n\n", m["yw_num"])
	fmt.Println("--------------------------------------------------------")
}

func (self *KYFWAction) Close() {
	self.persis.MgoDB.Close()
}
