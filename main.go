package main

import (
	"github.com/shaalx/sstruct/action"
	"time"
)

func main() {
	// var act action.Action
	// act := action.ZhihuMgoAction{}
	act := action.ToutiaoAction{}
	// act := action.KuwoMgoAction{}
	// act := action.YodaoMgoAction{}
	// act := action.ItunesMgoAction{}

	act.Init()
	for {
		act.Persistence()
		time.Sleep(time.Second * 7200)
	}
	// act.QueryOne()
	// act.Analyse()
	act.LatestNews()

	time.Sleep(time.Second * 4)
}
