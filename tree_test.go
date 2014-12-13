package sstruct

import (
	"container/list"
	"fmt"
	"reflect"
	"testing"
)

func Test_Tree(t *testing.T) {
	fmt.Println()
	linfo := list.New()
	fish := TreeFish(reflect.ValueOf(apping), linfo)
	// linfo.PushFront(fish)
	ShowTreeFish(linfo, 0)
	ShowTreeFish(fish, 0)

}

// func Benchmark_Analyse(b *testing.B) {
// 	// for i := 0; i < 1; i++ { //use b.N for looping
// 	// 	Analyse(apping)
// 	// }
// }
