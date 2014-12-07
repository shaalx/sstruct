package service

import (
	"github.com/shaalx/sstruct/service/fetch"
)

type Action struct {
	zhihuMgoAction ZhihuMgoAction
	// Do             func() bool
	// Fetch          func(url, ipaddr string) []byte
	// Persistence    func(bs []byte) bool
}

func (a *Action) Do() bool {
	return true
}

func (a *Action) Fetch(url, ipaddr string) []byte {
	return fetch.Do(url, ipaddr)
}

// func (a *Action) ZhihuMgoPersistence(bs []byte) bool {
// 	return a.zhihuMgoAction.Persistence()
// }
