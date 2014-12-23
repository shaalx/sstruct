package main

import (
	"github.com/shaalx/sstruct/service"
	// "time"
)

func main() {
	// var act action.Action
	// act := action.ZhihuMgoAction{}
	// act := service.ToutiaoAction{}
	// act := action.KuwoMgoAction{}
	// act := action.YodaoMgoAction{}
	// act := action.ItunesMgoAction{}
	act := service.KYFWAction{}

	act.Init()
	// for {
	// act.Persistence()
	// 	time.Sleep(time.Second * 7200)
	// }
	// act.QueryOne()
	// act.Analyse()
	act.Search()
	// act.LatestNews()

	// time.Sleep(time.Second * 4)
}
