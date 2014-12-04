package action

import (
	"fmt"
	"github.com/shaalx/sstruct/fetch"
	"github.com/shaalx/sstruct/mgodb"
	"github.com/shaalx/sstruct/persistence"
)

type YodaoMgoAction struct {
	persis persistence.MgoPersistence
}

var (
	YodaoServer = []string{"", "newsmgo", "yodao"}
)

func (y *YodaoMgoAction) Init() {
	y.persis.MgoDB = mgodb.SetLocalDB(YodaoServer...)
}

func (y *YodaoMgoAction) Persistence() {
	url := "http://dict.youdao.com/jsonapi?client=mobile&le=eng&q=simlpe"
	ipaddr := "202.120.87.152"
	bs := fetch.Do(url, ipaddr)
	fmt.Println(string(bs))
	y.persis.Do(bs)
}
