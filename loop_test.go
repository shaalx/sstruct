package sstruct

import (
	// "container/list"
	// "fmt"
	"reflect"
	"testing"
)

func Test_loop(t *testing.T) {
	loopAnalyse(reflect.ValueOf(apping))
	PrintMap()
}

// func Benchmark_Analyse(b *testing.B) {
// 	// for i := 0; i < 1; i++ { //use b.N for looping
// 	// 	Analyse(apping)
// 	// }
// }
