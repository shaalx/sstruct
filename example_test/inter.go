package main

import (
	// "bytes"
	"bytes"
	"encoding/json"
	"fmt"
	. "github.com/shaalx/sstruct/structs"
	"io"
	"reflect"
	"strings"
)

func main() {
	// origin := "[{{} 3 1 2 .jpg /admin/shops/activity/home?id=15  1} {{} 4 1 7 .jpg /admin/shops/activity/home?id=21  1} {{} 5 1 7 .jpg /admin/shops/activity/home?id=24  1}]"

	// solutionOne(origin)

	// solutionTwo(origin)
	// testST_print()
	// testST_reflect()
	testST_reflect2()
	// test_Map()
}

func solutionOne(origin string) {
	origin = strings.Replace(origin, "{", "", -1)
	origin = strings.Replace(origin, "}", "", -1)
	origin = strings.Replace(origin, "[", "", -1)
	origin = strings.Replace(origin, "]", "", -1)
	ss := strings.Split(origin, " ")
	for _, it := range ss {
		fmt.Println(it)
	}
}

func solutionTwo(origin string) {
	var st interface{}
	json.Unmarshal([]byte(origin), &st)
	fmt.Println(st)

}

func testST_print() {
	st := Structs{nil, 1, 2, 3, ".jpg", "/admin/a/b", 4}
	fmt.Println(st)
	fmt.Printf("#v %#v\n", st)
	fmt.Printf("v %v\n", st)
	fmt.Printf("T %T\n", st)

	st_1 := fmt.Sprintf("%v", st)
	st_2 := fmt.Sprint(st)
	fmt.Println(st_1)
	fmt.Println(st_2)
}

func testST_reflect() {
	st := Structs{nil, 1, 2, 3, ".jpg", "/admin/a/b", 4}
	typeOf := reflect.TypeOf(st)
	fmt.Println(typeOf)
	// elem := typeOf.Elem()
	// fmt.Println(elem.String())
	// var newst elem
	// fmt.Println(newst)
	// v := reflect.ValueOf(st)
	// newt := reflect.New(v.Type())
	// newt.Set(newt)
	// fmt.Println(newt)

	var r io.Reader
	b := new(bytes.Buffer)
	rv := reflect.ValueOf(&r).Elem()
	rv.Set(reflect.ValueOf(b))
	if r != b {
		fmt.Errorf("after Set: r=%T(%v)", r, r)
	}
	fmt.Println(rv)
	fmt.Println(b)
}

func testST_reflect2() {
	var r *Structs
	b := new(Structs)
	rv := reflect.ValueOf(&r).Elem()
	rv.Set(reflect.ValueOf(b))
	if r != b {
		fmt.Errorf("after Set: r=%T(%v)", r, r)
	}
	fmt.Println(rv)
	fmt.Println(b)

	fmt.Println(rv.Type())
	fmt.Println(rv.Elem())
}

func test_Map() {
	var a map[string]interface{}
	a = make(map[string]interface{}, 1)
	a["one"] = 1
	fmt.Println(a)
	fmt.Printf("%#v\n", a)
	fmt.Printf("%v\n", a)
	fmt.Printf("%T", a)

	// var ptlogin map[int64]map[string]interface{}
	// ptlogin = make(map[int64]map[string]interface{}, 1)
	// b := make(map[string]interface{}, 1)
	// b["one"] = 1
	// ptlogin[int64(1)] = b
	// fmt.Println(ptlogin)
}