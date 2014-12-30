package action

import (
	. "github.com/shaalx/sstruct/bean"
	"github.com/shaalx/sstruct/service"
	"time"
)

func ToutiaoLatestNews() []Toutiao {
	toutiao := service.ToutiaoAction{}
	toutiao.Init()
	return toutiao.TTLatestNews()
}

func KyfwAction() {
	var serv service.Service
	serv = &service.KYFWAction{}
	serv.Init()
	// serv.Persistence()
	// serv.Analyse()
	for {
		serv.Search()
		time.Sleep(time.Second * 1)
	}
	serv.Close()
}
