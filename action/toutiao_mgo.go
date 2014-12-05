package action

import (
	"fmt"
	"github.com/shaalx/sstruct/bean"
	"github.com/shaalx/sstruct/fetch"
	"github.com/shaalx/sstruct/mgodb"
	"github.com/shaalx/sstruct/persistence"
	// "github.com/shaalx/sstruct/pkg3/mgo/bson"
	"github.com/shaalx/sstruct/search"
)

type ToutiaoAction struct {
	persis persistence.MgoPersistence
}

var (
	ToutiaoServer = []string{"", "newsmgo", "toutiao"}
)

func (self *ToutiaoAction) Init() {
	self.persis.MgoDB = mgodb.SetLocalDB(ToutiaoServer...)
}

func (self *ToutiaoAction) Persistence() {
	url := "http://ic.snssdk.com/2/article/v21/stream/?iid=2463135716&city=%E5%8D%97%E4%BA%AC%E5%B8%82%E5%B8%82%E8%BE%96%E5%8C%BA&longitude=121.3861245975094&latitude=31.23194430763986&user_city=%E5%8D%97%E4%BA%AC&min_behot_time=0&detail=1&image=1&count=20&ac=wifi&channel=App%20Store&app_name=news_article&aid=13&version_code=4.2&device_platform=ipad&os_version=8.1&device_type=iPad%20Mini%20Retina&vid=0256CB17-8BBE-4974-B25D-B4691079ACDC&openudid=1663351e7f057fe184db98ac159e9971e590aef8&idfa=D5A1F8CF-75C5-4DB9-8C3E-CCD702148275"
	// url := "http://ic.snssdk.com/2/article/v10/hot_comments/?iid=2463135716&ac=wifi&channel=App%20Store&app_name=news_article&aid=13&version_code=4.2&device_platform=ipad&os_version=8.1&device_type=iPad%20Mini%20Retina&vid=0256CB17-8BBE-4974-B25D-B4691079ACDC&openudid=1663351e7f057fe184db98ac159e9971e590aef8&idfa=D5A1F8CF-75C5-4DB9-8C3E-CCD702148275"
	ipaddr := "202.120.87.152"
	bs := fetch.Do(url, ipaddr)
	// fmt.Println(string(bs))
	self.persis.Do(bs)
}

func (self *ToutiaoAction) QueryOne() {
	one := self.persis.QueryOne()
	fmt.Println(one)
	bs := bean.I2Bytes(one)
	fmt.Println(string(bs))
}

func (self *ToutiaoAction) Analyse() {
	one := self.persis.QueryNewsOne(nil)
	buf := bean.I2Bytes(one.Content)
	// key := "display_info"
	// path := []string{"tips"}
	// sear := search.SearchSValue(buf, key, path...)
	// if sear == nil {
	// 	return
	// }
	// fmt.Println(*sear)

	// totn := search.SearchIValue(buf, "total_number", []string{}...)
	// fmt.Println(totn)

	ary := search.SearchArray(buf, "data", []string{}...)
	fmt.Println(len(ary))

	ary0buf := bean.I2Bytes(ary[0])
	abstract := search.SearchSValue(ary0buf, "abstract", []string{}...)
	if abstract != nil {
		fmt.Println(*abstract)
	}

	maps := search.TTStem(ary0buf)
	fmt.Println(maps)

	// abs := maps["abstract"]
	// abss, ok := abs.(*string)
	// if ok {
	// 	fmt.Println(*abss)
	// }
	TTShow(maps)
	TTShows(buf)
}

func TTShow(stem map[string]interface{}) {
	for key, val := range stem {
		// fmt.Println(key, val)
		switch key {
		case "tag", "keywords", "abstract", "title", "article_url", "middle_image":
			vals, ok := val.(*string)
			if ok && vals != nil {
				fmt.Printf("%s : %s\n", key, *vals)
			}
		case "has_image", "publish_time":
			fmt.Println(val)
		}
	}
}

func TTShows(data []byte) {
	arys := search.SearchArray(data, "data", []string{}...)
	for i, ary := range arys {
		aryb := bean.I2Bytes(ary)
		maps := search.TTStem(aryb)
		TTShow(maps)
		fmt.Println("*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*", i)
		fmt.Println()
	}
}

func (self *ToutiaoAction) LatestNews() {
	one := self.persis.QuerySortedNewsOne(nil, "unixdate")
	buf := bean.I2Bytes(one.Content)
	TTShows(buf)
}
