package action

import (
	"fmt"
	"github.com/shaalx/sstruct/fetch"
	"github.com/shaalx/sstruct/mgodb"
	"github.com/shaalx/sstruct/persistence"
)

type ItunesMgoAction struct {
	persis persistence.MgoPersistence
}

var (
	ItunesServer = []string{"", "newsmgo", "itunes"}
)

func (i *ItunesMgoAction) Init() {
	i.persis.MgoDB = mgodb.SetLocalDB(ItunesServer...)
}

func (i *ItunesMgoAction) Persistence() {
	// url := "http://news-at.zhihu.com/api/3/theme/12"
	// url := "http://fanyi.youdao.com/openapi.do?keyfrom=sasfasdfasf&key=1177596287&type=data&doctype=json&version=1.1&q=love"
	url := "https://itunes.apple.com/WebObjects/MZStore.woa/wa/viewRoom?fcId=560264462&genreIdString=36&mediaTypeString=Mobile+Software+Applications"
	ipaddr := "202.120.87.152"
	bs := fetch.Do(url, ipaddr)
	fmt.Println(string(bs))
	i.persis.Do(bs)
}
