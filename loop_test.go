package sstruct

import (
	// "container/list"
	"fmt"
	"github.com/astaxie/beego"
	"reflect"
	"testing"
	"time"
)

type Appin struct {
	Name    string
	Address string
	// Child   Node
	Age int
	// Apps   *Apping
	Users  []string
	Appeds Appe
	Apped  Appe
}

type Appe struct {
	Name    string
	Address string
	// Child   Node
	Age   int
	Apps  Apping
	Users []string
}

var (
	appin = Appin{"Shanghai", "ECNU", 25, []string{"User1", "User2"}, Appe{}, Appe{}}
)

func Test_loop(t *testing.T) {
	tim := time.ParseError{}
	fmt.Printf("%#v\n", appin)
	loopAnalyse(reflect.ValueOf(appin))
	// PrintMap()
	loopAnalyse(reflect.ValueOf(tim))
	// PrintMap()

	name := beego.Namespace{}
	loopAnalyse(reflect.ValueOf(name))
	PrintMap()
}

// func Benchmark_Analyse(b *testing.B) {
// 	// for i := 0; i < 1; i++ { //use b.N for looping
// 	// 	Analyse(apping)
// 	// }
// }
