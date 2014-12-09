package main

import (
	"fmt"
	. "github.com/shaalx/sstructs/sstructs"
	// "strings"
)

func main() {
	// test_normal()
	test_imnormal()
}

func test_normal() {
	app := App{Name: "City", Address: "Shanghai"}
	// newApp := sstructs.Set(app, []string{"city", "state"})
	// fmt.Println(newApp)

	// newApp2 := sstructs.Sets(app, []string{"city", "state"}...)
	// fmt.Println(newApp2)

	// // sstructs.Analyse(app)
	// // sstructs.Analyse(newApp2)
	// sstructs.AnalyseAddrLoop(&newApp2)
	// appi := Set(app, []string{"city", "state"})
	app.Println()
	newApp := NewInstance(app)
	fmt.Println(newApp)
	appNews, ok := newApp.(App)
	fmt.Println(appNews, ok)
}

func test_imnormal() {
	app := App{Name: "City", Address: "Shanghai"}
	iapp := ChangeWithIJ(app, 0)
	fmt.Println(iapp, app)
	a, ok := iapp.(*App)
	fmt.Println(a, ok)
	// app.Println()
	// i := Change(app)
	// fmt.Println(i)
	// ChangeWithI(i, app)
	// alloc1 := Alloc(app)
	// fmt.Println(alloc1)

	// alloc2 := Alloc(&app)
	// fmt.Println(alloc2)

	// alloc3 := Alloc(alloc2)
	// fmt.Println(alloc3)
	// Jsonable(app)
	// Jsonable(alloc1)
	// Jsonable(alloc2)
	// Jsonable(alloc3)
	// AnalyseLoop(i)
}
