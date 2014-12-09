package main

import (
	// "bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type ST struct {
	one   *ST
	two   int
	three int
	four  int
	five  string
	six   string
	seven int
}

func main() {
	origin := "[{{} 3 1 2 .jpg /admin/shops/activity/home?id=15  1} {{} 4 1 7 .jpg /admin/shops/activity/home?id=21  1} {{} 5 1 7 .jpg /admin/shops/activity/home?id=24  1}]"

	// solutionOne(origin)
	solutionTwo(origin)
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
	testST()
}

func testST() {
	// var st ST
	st := ST{nil, 1, 2, 3, ".jpg", "/admin/a/b", 4}
	st.one = &st
	fmt.Println(st)
	fmt.Printf("#v %#v\n", st)
	fmt.Printf("v %v\n", st)
	fmt.Printf("T %T\n", st)
	fmt.Printf("c %c\n", st)
}
