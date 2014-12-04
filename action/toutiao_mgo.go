package action

import (
	"fmt"
	"github.com/shaalx/sstruct/fetch"
	"github.com/shaalx/sstruct/mgodb"
	"github.com/shaalx/sstruct/persistence"
)

type ToutiaoAction struct {
	persis persistence.MgoPersistence
}

var (
	ToutiaoServer = []string{"", "newsmgo", "toutiao"}
)

func (t *ToutiaoAction) Init() {
	t.persis.MgoDB = mgodb.SetLocalDB(ToutiaoServer...)
}

func (t *ToutiaoAction) Persistence() {
	url := "http://ic.snssdk.com/2/article/v21/stream/?iid=2463135716&city=%E5%8D%97%E4%BA%AC%E5%B8%82%E5%B8%82%E8%BE%96%E5%8C%BA&longitude=121.3861245975094&latitude=31.23194430763986&user_city=%E5%8D%97%E4%BA%AC&min_behot_time=0&detail=1&image=1&count=20&ac=wifi&channel=App%20Store&app_name=news_article&aid=13&version_code=4.2&device_platform=ipad&os_version=8.1&device_type=iPad%20Mini%20Retina&vid=0256CB17-8BBE-4974-B25D-B4691079ACDC&openudid=1663351e7f057fe184db98ac159e9971e590aef8&idfa=D5A1F8CF-75C5-4DB9-8C3E-CCD702148275"
	// url := "http://ic.snssdk.com/2/article/v10/hot_comments/?iid=2463135716&ac=wifi&channel=App%20Store&app_name=news_article&aid=13&version_code=4.2&device_platform=ipad&os_version=8.1&device_type=iPad%20Mini%20Retina&vid=0256CB17-8BBE-4974-B25D-B4691079ACDC&openudid=1663351e7f057fe184db98ac159e9971e590aef8&idfa=D5A1F8CF-75C5-4DB9-8C3E-CCD702148275"
	ipaddr := "202.120.87.152"
	bs := fetch.Do(url, ipaddr)
	fmt.Println(string(bs))
	t.persis.Do(bs)
}

func (t *ToutiaoAction) QueryOne() {
	one := t.persis.QueryOne()
	fmt.Println(one)
}

func (t *ToutiaoAction) Analyse() {

}
