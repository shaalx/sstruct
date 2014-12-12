package sstruct

import (
	"container/list"
	"fmt"
	"reflect"
	"testing"
)

// type Apping struct {
// 	Name    string
// 	Address string
// 	// Child   Node
// 	Age int
// 	// Apps   *Apping
// 	Users  []string
// 	Appeds Apped
// }

// type Apped struct {
// 	Name    string
// 	Address string
// 	// Child   Node
// 	Age int
// 	// Apps  *Apping
// 	Users []string
// }

// var (
// 	appi = Apping{"Shanghai", "ECNU", 25, []string{"User1", "User2"}, Apped{}}
// )

func Test_Tree(t *testing.T) {
	fmt.Println()
	linfo := list.New()
	tree(reflect.ValueOf(appin), linfo)
	for e := linfo.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}

}

// func Benchmark_Analyse(b *testing.B) {
// 	// for i := 0; i < 1; i++ { //use b.N for looping
// 	// 	Analyse(apping)
// 	// }
// }
