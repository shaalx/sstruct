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
<<<<<<< HEAD:action/toutiao.go
	// serv = &service.KYFWAction{}
	serv = &service.ToutiaoAction{}
=======
	serv = &service.BDLENAction{}
>>>>>>> a16c14a9b4686175f372904a5c5718c682d6b78a:action/actions.go
	serv.Init()
	serv.Persistence()
	// serv.Analyse()
	// for {
<<<<<<< HEAD:action/toutiao.go
	// 	serv.Search()
	time.Sleep(time.Second * 1)
	// }
=======
	serv.Search()
	time.Sleep(time.Second * 1)
	// }
	serv.Close()
}

func TopicAction() {
	var serv service.Service
	serv = &service.TopicAction{}
	serv.Init()
	// serv.Persistence()
	// serv.Search()
	serv.Analyse()
	time.Sleep(time.Second * 1)
>>>>>>> a16c14a9b4686175f372904a5c5718c682d6b78a:action/actions.go
	serv.Close()
}
