package main

import (
	"github.com/shaalx/sstruct/action"
)

func main() {
	// var act action.Action
	// act := action.ZhihuMgoAction{}
	act := action.ToutiaoAction{}
	// act := action.KuwoMgoAction{}
	// act := action.YodaoMgoAction{}

	act.Init()
	act.Persistence()
	// act.QueryOne()
}
