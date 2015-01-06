package service

import (
	"fmt"
	. "github.com/shaalx/sstruct/bean"
	"github.com/shaalx/sstruct/persistence"
	"github.com/shaalx/sstruct/persistence/mgodb"
	"github.com/shaalx/sstruct/service/fetch"
	"github.com/shaalx/sstruct/service/search"
	"github.com/shaalx/sstruct/utils"
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
	url := "http://toutiao.com/api/article/recent/?callback=jQuery17206307146116159856_1420533186372&category=news_society&count=20&max_behot_time=1420533186.73&offset=0&utm_source=toutiao&_=1420533189249"
	// url := "http://ic.snssdk.com/2/article/v21/stream/?iid=2463135716&city=%E5%8D%97%E4%BA%AC%E5%B8%82%E5%B8%82%E8%BE%96%E5%8C%BA&longitude=121.3861245975094&latitude=31.23194430763986&user_city=%E5%8D%97%E4%BA%AC&min_behot_time=0&detail=1&image=1&count=20&ac=wifi&channel=App%20Store&app_name=news_article&aid=13&version_code=4.2&device_platform=ipad&os_version=8.1&device_type=iPad%20Mini%20Retina&vid=0256CB17-8BBE-4974-B25D-B4691079ACDC&openudid=1663351e7f057fe184db98ac159e9971e590aef8&idfa=D5A1F8CF-75C5-4DB9-8C3E-CCD702148275"
	// url := "http://ic.snssdk.com/2/article/v10/hot_comments/?iid=2463135716&ac=wifi&channel=App%20Store&app_name=news_article&aid=13&version_code=4.2&device_platform=ipad&os_version=8.1&device_type=iPad%20Mini%20Retina&vid=0256CB17-8BBE-4974-B25D-B4691079ACDC&openudid=1663351e7f057fe184db98ac159e9971e590aef8&idfa=D5A1F8CF-75C5-4DB9-8C3E-CCD702148275"
	ipaddr := "202.120.87.152"
	bs := fetch.Do(url, ipaddr)
	// fmt.Println(string(bs))
	self.persis.Do(bs)
}

func (self *ToutiaoAction) QueryOne() {
	one := self.persis.QueryOne()
	fmt.Println(one)
	bs := utils.I2Bytes(one)
	fmt.Println(string(bs))
}

func (self *ToutiaoAction) Analyse() {
	one := self.persis.QueryNewsOne(nil)
	buf := utils.I2Bytes(one.Content)
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

	ary0buf := utils.I2Bytes(ary[0])
	abstract := search.SearchSValue(ary0buf, "abstract", []string{}...)
	if abstract != "" {
		fmt.Println(abstract)
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
			fmt.Printf("%s : %s\n", key, val)
		case "has_image", "publish_time":
			fmt.Println(val)
		}
	}
}

func TTContent(data []byte) []map[string]interface{} {
	arys := search.SearchArray(data, "data", []string{}...)
	result := make([]map[string]interface{}, len(arys))
	for i, ary := range arys {
		aryb := utils.I2Bytes(ary)
		maps := search.TTStem(aryb)
		result[i] = maps
	}
	return result
}

func TTContents(data []byte) []Toutiao {
	arys := search.SearchArray(data, "data", []string{}...)
	result := make([]Toutiao, len(arys))
	for i, ary := range arys {
		aryb := utils.I2Bytes(ary)
		maps := search.TTStem(aryb)
		// var tt Toutiao
		// tt.Title = getKey(maps, "title")
		tt := Toutiao{Title: getKey(maps, "title"), Abstract: getKey(maps, "abstract"), Keywords: getKey(maps, "keywords"), Middle_image: getKey(maps, "middle_image"), Article_url: getKey(maps, "article_url")}
		result[i] = tt
	}
	return result
}
func getKey(maps map[string]interface{}, key string) string {
	if maps == nil {
		return ""
	}
	val, ok := maps[key].(string)
	if ok {
		return val
	}
	return ""
}
func TTShows(data []byte) {
	arys := search.SearchArray(data, "data", []string{}...)
	for i, ary := range arys {
		aryb := utils.I2Bytes(ary)
		maps := search.TTStem(aryb)
		TTShow(maps)
		fmt.Println("*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*", i)
		fmt.Println()
	}
}

func (self *ToutiaoAction) LatestNews() []map[string]interface{} {
	one := self.persis.QuerySortedNewsOne(nil, "-unixdate")
	buf := utils.I2Bytes(one.Content)
	TTShows(buf)
	return TTContent(buf)
}

func (self *ToutiaoAction) TTLatestNews() []Toutiao {
	one := self.persis.QuerySortedNewsOne(nil, "-unixdate")
	buf := utils.I2Bytes(one.Content)
	return TTContents(buf)
}

func (self *ToutiaoAction) News(i int64) interface{} {
	fmt.Println(i)
	one := self.persis.QuerySortedNewsOne(nil, "-unixdate")
	buf := utils.I2Bytes(one.Content)
	return TTContents(buf)
}
