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
	// serv = &service.KYFWAction{}
	serv = &service.ToutiaoAction{}
	serv = &service.BDLENAction{}
	serv.Init()
	serv.Persistence()
	serv.Analyse(0)
	for {
		serv.Search()
		time.Sleep(time.Second * 1)
	}
	defer serv.Close()
}
