package sstruct

import (
// "container/list"
// "fmt"
// "testing"
)

type Apping struct {
	Name    string
	Address string
	Age     int
	Users   []string
	Appeds  Apped
	Apps    Apped
}

type Apped struct {
	Name    string
	Address string
	Age     int
	Apps    *Apping
	Users   []string
	N       Node
}
type Node struct {
	O  string
	T  int
	TH float32
}

var (
	apping = Apping{"Shang海", "ECNU", 25, []string{"User1", "User2"}, Apped{}, Apped{}}
)

// func Test_Analyse(t *testing.T) {
// 	fmt.Printf("%#v\n", apping)
// 	linfo := list.New()
// 	linfoing := Analyse(apping, linfo)
// 	for e := linfo.Back(); e != nil; e = e.Prev() {
// 		fmt.Println(e.Value)
// 	}
// 	fmt.Println()

// 	for e := linfoing.Back(); e != nil; e = e.Prev() {
// 		fmt.Println(e.Value)
// 	}
// 	fmt.Println()
// 	// linfoPtr := list.New()
// 	// Analyse(&apping, linfoPtr)
// 	// for e := linfoPtr.Back(); e != nil; e = e.Prev() {
// 	// 	fmt.Println(e.Value)
// 	// }
// 	// fmt.Println()
// }

// func Benchmark_Analyse(b *testing.B) {
// 	// for i := 0; i < 1; i++ { //use b.N for looping
// 	// 	Analyse(apping)
// 	// }
// }
