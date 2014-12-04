package action

import (
	"fmt"
	"github.com/shaalx/sstruct/fetch"
	"github.com/shaalx/sstruct/persistence"
)

type KuwoMgoAction struct {
	persis persistence.MgoPersistence
}

func (z *KuwoMgoAction) Do() {
	url := "http://qukudata.kuwo.cn/q.k?op=query&cont=ninfo&node=26&pn=0&rn=10&fmt=json&src=pad&callback=showNodeChild&r=2014123"
	ipaddr := "202.120.87.152"
	bs := fetch.Do(url, ipaddr)
	fmt.Println(string(bs))
	var p persistence.MgoPersistence
	p.Server = []string{"", "newsmgo", "kuwo"}
	p.Do(bs)
}
