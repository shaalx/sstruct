package sstructs

import (
	"fmt"
	"reflect"
	// "unsafe"
)

type App struct {
	Name    string
	Address string
	// Child   Node
	Age  int
	Apps *App
}

func (self *App) String() string {
	prints := fmt.Sprint(*self)
	// prints += "{1231 {345345 {4545656}}}"
	return prints
}

func (self *App) Println() {
	fmt.Printf("%v\n", self)
}

type Node struct {
	Stuff []string
	// Child *Node
}

func Set(instance interface{}, values []string) interface{} {
	typeOf := reflect.TypeOf(instance)
	newInstance := reflect.New(typeOf)
	app := newInstance.Interface()
	elem := reflect.ValueOf(app).Elem()
	for i, value := range values {
		elem.Field(i).SetString(value)
	}
	// fmt.Println(app)
	return app
}

func Sets(instance interface{}, values ...string) interface{} {
	typeOf := reflect.TypeOf(instance)
	newInstance := reflect.New(typeOf)
	app := newInstance.Interface()
	elem := reflect.ValueOf(app).Elem()

	for i, v := range values {

		valueOf := reflect.ValueOf(v)
		// for i, value := range values {
		// 	elem.Field(i).Set(value)
		// }
		// elem.
		elem.Field(i).Set(valueOf)
	}
	// fmt.Println(app)
	return app
}

// address the instance
func Analyse(instance interface{}) {
	typeOf := reflect.TypeOf(instance)
	name := typeOf.Name()
	if 0 == len(name) {
		AnalyseAddr(instance)
		return
	}
	fmt.Printf("name: %v\n", typeOf.Name())
	fmt.Printf("type: %v\n", typeOf)

	for i := 0; i < typeOf.NumField(); i++ {
		name := typeOf.Field(i).Name
		fmt.Println(name)
	}
}

// address the instance of pointer
func AnalyseAddr(instance interface{}) {
	// typeOf := reflect.TypeOf(instance)
	// fmt.Printf("name: -%v-\n", typeOf.Name())
	// fmt.Printf("type: %v\n", typeOf)

	valueOf := reflect.ValueOf(instance).Elem()
	fmt.Printf("value: %v\n", valueOf)

	typeValueOf := valueOf.Type()
	if 0 == len(typeValueOf.Name()) {
		return
	}
	for i := 0; i < typeValueOf.NumField(); i++ {
		name := typeValueOf.Field(i).Name
		fmt.Println(name)
	}
}

// address the instance of pointer-loop
func AnalyseLoop(instance interface{}) interface{} {
	i := address(instance)
	fmt.Println(i)
	return i
}

// the instance is pointor ?
func addressed(instance interface{}) bool {
	fmt.Println("type is :", reflect.TypeOf(instance))
	if 0 == len(reflect.TypeOf(instance).Name()) {
		return true
	}
	return false
}

// address the instance of pointer
func address(instance interface{}) interface{} {
	if !addressed(instance) {
		return instance
	}
	return newInstance(instance)
}

func newInstance(instance interface{}) interface{} {
	if !addressed(instance) {
		return instance
	}
	fmt.Print("new instance : ")
	valueOf := reflect.ValueOf(instance).Elem()
	typeValueOf := valueOf.Type()
	fmt.Println(typeValueOf)
	newInstance := reflect.New(typeValueOf)
	app := newInstance.Interface().(App)
	return app
}

func NewInstance(instance interface{}) interface{} {
	return newInstance(instance)
}
