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

type BDLENAction struct {
	persis persistence.MgoPersistence
}

var (
	BDLENServer = []string{"", "bdlen", "len"}
)

func (self *BDLENAction) Init() {
	self.persis.MgoDB = mgodb.SetLocalDB(BDLENServer...)
}

func (self *BDLENAction) Persistence() {
	url := "http://map.baidu.com/?newmap=1&reqflag=pcmap&biz=1&from=webmap&qt=nav&da_src=pcmappg.searchBox.button&c=1&sn=2$$$$$$%E5%8C%97%E4%BA%AC%E5%B8%82$$0$$$$&en=2$$$$$$%E7%9F%B3%E5%AE%B6%E5%BA%84%E5%B8%82$$0$$$$&sc=1&ec=1&rn=5&time_index=-1&day=-1&extinfo=63&tn=B_NORMAL_MAP&nn=0&ie=utf-8&l=7&b=(12184537.219999999,3297450.195;14087129.219999999,4903082.195)&t=1421214039805"
	ipaddr := "202.120.87.152"
	bs := fetch.Do(url, ipaddr)
	self.persis.Do(bs)
}

func (self *BDLENAction) QueryOne() {
	one := self.persis.QueryOne()
	fmt.Println(one)
	bs := utils.I2Bytes(one)
	fmt.Println(string(bs))
}

func (self *BDLENAction) Analyse() {
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

func Searching2(data []byte) {
	dis := search.SearchFIValue(data, "dis", []string{"content"}...)
	start := search.SearchSValue(data, "start_name", []string{"result"}...)
	end := search.SearchSValue(data, "end_name", []string{"result"}...)
	fmt.Printf("%s --> %s : %v \n", start, end, float64(dis)/1000.0)
}

func (self *BDLENAction) Search() {
	url := `http://ltpapi.voicecloud.cn/analysis/?api_key=YourApiKey&text=盗梦空间是一部好电影，` +
		`大家对它的评价非常的高。` +
		`&format=json`
	// url := "http://map.baidu.com/?newmap=1&reqflag=pcmap&biz=1&from=webmap&qt=nav&da_src=pcmappg.searchBox.button&c=1&sn=2$$$$$$上海市$$0$$$$&en=2$$$$$$石家庄市$$0$$$$&sc=1&ec=1&rn=5&time_index=-1&day=-1&extinfo=63&tn=B_NORMAL_MAP&nn=0&ie=utf-8&l=7&b=(12184537.219999999,3297450.195;14087129.219999999,4903082.195)&t=1421214039805"
	ipaddr := "202.120.87.152"
	bs := fetch.Do1(url, ipaddr)
	// Searching2(bs)
	self.persis.Do(bs)
}

func KYFWShow2(m map[string]interface{}) {
	fmt.Println("--------------------------------------------------------\n")
	fmt.Printf("%v : %v >>> %v  %v ~ %v  < %v > \n\n", m["station_train_code"], m["from_station_name"], m["to_station_name"], m["start_time"], m["arrive_time"], m["lishi"])
	fmt.Printf("[软座] ：%v\t", m["rz_num"])
	fmt.Printf("[硬座] ：%v\t", m["yz_num"])
	fmt.Printf("[硬卧] ：%v\n\n", m["yw_num"])
	fmt.Println("--------------------------------------------------------")
}

func (self *BDLENAction) Close() {
	self.persis.MgoDB.Close()
}
