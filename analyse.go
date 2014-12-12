package sstruct

import (
	"container/list"
	"fmt"
	"reflect"
	"strings"
	// "unsafe"
)

// analyse the instance of struct
func Analyse(instance interface{}, linfo *list.List) *list.List {
	typeOf := reflect.TypeOf(instance)
	if Exist(linfo, typeOf) {
		return linfo
	}
	// linfoing1 := list.New()
	kind := typeOf.Kind()
	switch kind {
	case reflect.Ptr:
		AnalysePtr(instance, linfo)
		return linfo
	case reflect.Struct:
		linfo.PushFront(typeOf)
		/*linfoing1 = */ Analyseing(reflect.ValueOf(instance), linfo)

		elem := reflect.ValueOf(instance)
		for i := 0; i < elem.NumField(); i++ {
			t := elem.Field(i).Type()
			linfo.PushFront(t)
			// /*linfoing2 := */ Analyseing(elem.Field(i), linfo)
			// Join(linfoing1, linfoing2)
		}
		for i := 0; i < elem.NumField(); i++ {
			// t := elem.Field(i).Type()
			// linfo.PushFront(t)
			/*linfoing2 := */ Analyseing(elem.Field(i), linfo)
			// Join(linfoing1, linfoing2)
		}
	default:
		// return linfoing1
	}
	return linfo /*ing1*/
}

// analyse the instance of struct  ptr
func AnalysePtr(instance interface{}, linfo *list.List) *list.List {
	typeOf := reflect.TypeOf(instance)
	if reflect.Ptr != typeOf.Kind() || Exist(linfo, typeOf) {
		return linfo
	}
	linfo.PushFront(reflect.TypeOf(instance))
	elem := reflect.ValueOf(instance).Elem()
	for i := 0; i < elem.NumField(); i++ {
		t := elem.Field(i).Type()
		linfo.PushFront(t)
	}
	return linfo
}

func Analyseing(valueOf reflect.Value, linfo *list.List) *list.List {
	fmt.Println("\n.*.*.*.*.*.*.*.*.*.*.*.*.*.*.*")
	typeOf := valueOf.Type()
	// fmt.Println(typeOf)
	if Exist(linfo, typeOf) && reflect.Struct == typeOf.Kind() {
		return linfo
	}
	// linfo := list.New()
	linfo.PushFront(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	switch typeOf.Kind() {
	case reflect.Struct:
		fmt.Println("  Struct")
		for i := 0; i < valueOf.NumField(); i++ {
			t := valueOf.Field(i).Type()
			linfo.PushFront(t)
			// fmt.Println(t)
		}
	// case reflect.Ptr:
	// 	fmt.Println("  Ptr")
	// 	if sysType(typeOf.String()) {
	// 		fmt.Println("\tSystem type")
	// 	} else {
	// 		newInstance := reflect.New(typeOf)
	// 		app := newInstance.Interface()
	// 		fmt.Println(app)
	// 		// v := reflect.NewAt(typeOf, unsafe.New(app))
	// 		// fmt.Println(v)
	// 		// elem := reflect.ValueOf(app).Elem()
	// 		// for i := 0; i < elem.NumField(); i++ {
	// 		// 	fmt.Println(elem.Field(i))
	// 		// }

	// 	}
	default:
		// fmt.Println("  default system type")
		// linfo.PushFront(typeOf)
	}
	fmt.Println("-------------------------------\n")
	linfo.PushFront("\n")
	return linfo
}

// check the type exist in the list
func Exist(linfo *list.List, typeOf reflect.Type) bool {
	for e := linfo.Back(); e != nil; e = e.Prev() {
		etype, ok := e.Value.(reflect.Type)
		if ok && typeOf == etype {
			return true
		}
	}
	return false
}

func Join(ldst, lsrc *list.List) {
	for e := lsrc.Back(); e != nil; e = e.Prev() {
		ldst.PushFront(e.Value)
	}
}

func sysType(typeName string) bool {
	return !strings.Contains(typeName, ".")
}
