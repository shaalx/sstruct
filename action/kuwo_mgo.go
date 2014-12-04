package action

import (
	"fmt"
	"github.com/shaalx/sstruct/fetch"
	"github.com/shaalx/sstruct/persistence"
	"strings"
)

type KuwoMgoAction struct {
	persis persistence.MgoPersistence
}

func (z *KuwoMgoAction) Do() {
	url := "http://artistlistinfo.kuwo.cn/mb.slist?stype=artistlist&category=0&order=hot&pn=0&rn=20&callback=showAreaArtistList&r=1417619717508"
	ipaddr := "202.120.87.152"
	bs := fetch.Do(url, ipaddr)
	stmp := string(bs)
	stmp = strings.Replace(stmp, "try {var jsondata=", "", -1)
	stmp = strings.Replace(stmp, ";showAreaArtistList(jsondata);}catch(e){jsonError(e)}", "", -1)

	stmp = strings.Replace(stmp, "'total':'49522','pn':'0','rn':'20','category':'0','new_album':'1','new_album_cnt':'107','artistlist':", "", -1)
	fmt.Println(string(stmp))
	var p persistence.MgoPersistence
	p.Server = []string{"", "newsmgo", "kuwo"}
	p.Do([]byte(stmp))
}
