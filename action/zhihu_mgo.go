package action

import (
	"fmt"
	"github.com/shaalx/sstruct/fetch"
	"github.com/shaalx/sstruct/mgodb"
	"github.com/shaalx/sstruct/persistence"
)

type ZhihuMgoAction struct {
	persis persistence.MgoPersistence
}

var (
	ZhihuServer = []string{"", "newsmgo", "zhihu"}
)

func (z *ZhihuMgoAction) Init() {
	z.persis.MgoDB = mgodb.SetLocalDB(ZhihuServer...)
}

func (z *ZhihuMgoAction) Persistence() {
	// url := "http://news-at.zhihu.com/api/3/theme/12"
	url := "http://news-at.zhihu.com/api/3/news/latest"
	// url := "https://itunes.apple.com/WebObjects/MZStore.woa/wa/viewRoom?fcId=560264462&genreIdString=36&mediaTypeString=Mobile+Software+Applications"
	ipaddr := "202.120.87.152"
	bs := fetch.Do(url, ipaddr)
	fmt.Println(string(bs))
	z.persis.Do(bs)
}
