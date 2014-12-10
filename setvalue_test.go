package sstruct

// import (
// 	"fmt"
// 	"testing"
// )

// type App struct {
// 	Name    string
// 	Address string
// 	// Child   Node
// 	Age   int
// 	Apps  *App
// 	Users []string
// }

// var (
// 	app2 App           = App{Name: "City", Address: "Shanghai"}
// 	app  App           = App{Name: "City", Address: "Shanghai"}
// 	v    []interface{} = []interface{}{"chengshi", "yingyu", 24, &app2, []string{"u1", "u2", "u3"}}
// )

// func Test_SetValueCopy(t *testing.T) {
// 	res := SetValue(app, v)
// 	fmt.Println(".*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.* res")
// 	fmt.Printf("%#v\n", res)
// 	fmt.Println(".*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.* app")
// 	fmt.Printf("%#v\n", app)
// 	fmt.Println()
// }

// func Test_SetValuePtr(t *testing.T) {
// 	res := SetValue(&app, v)
// 	fmt.Println(".*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.* res")
// 	fmt.Printf("%#v\n", res)
// 	fmt.Println(".*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.* app")
// 	fmt.Printf("%#v\n", app)
// 	fmt.Println()
// }

// func Test_SetValueAtIPtr(t *testing.T) {
// 	for i, it := range v {
// 		res := SetValueAtI(&app, i, it)
// 		fmt.Println(".*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.* res")
// 		fmt.Printf("%#v\n", res)
// 		fmt.Println(".*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*.* app")
// 		fmt.Printf("%#v\n", app)
// 		fmt.Println()
// 	}

// }

// func Benchmark_SetValue(b *testing.B) {
// 	for i := 0; i < b.N; i++ { //use b.N for looping
// 		a := SetValue(app, v)
// 		SetValue(a, v)
// 	}
// }
