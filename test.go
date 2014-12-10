package main

import (
	"fmt"
	. "github.com/shaalx/sstruct/oper"
)

type App struct {
	Name    string
	Address string
	// Child   Node
	Age   int
	Apps  *App
	Users []string
}

func main() {
	test_set()
}

func test_set() {
	app := App{Name: "City", Address: "Shanghai"}
	app2 := App{Name: "City", Address: "Shanghai"}
	var v []interface{}
	v = make([]interface{}, 5)
	v[0] = "chengshi"
	v[1] = "yingyu"
	v[2] = 24
	v[3] = &app2
	v[4] = []string{"u1", "u2", "u3"}
	// res := SetValueOfCopy(&app, v)
	res := SetValue(app, v)
	SetValueAtI(&app, 0, "shijiazhuang")
	fmt.Println(".*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.* res")
	fmt.Printf("%#v\n", res)
	fmt.Println()
	// fmt.Println(".*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.* res json")
	// /*b := */ Jsonable(res)
	// fmt.Println()
	// fmt.Println(".*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.* res un-json")
	// Unjsonable(b)
	// fmt.Println()
	// // Analyse(res)
	// // str := Str{&res, res}
	// // Jsonable(str)
	// // fmt.Printf("%#v\n", str)
	// // fmt.Println("a", str.A)
	// // fmt.Println("b", str.B)
	fmt.Println(".*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.* app")
	fmt.Printf("%#v\n", app)
	fmt.Println()
	// fmt.Println(".*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.* app json")
	// Jsonable(app)
	// fmt.Println()
	// fmt.Println(".*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.* app un-json")
	// Jsonable(app)
	// fmt.Println()

}
