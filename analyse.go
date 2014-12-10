package sstruct

import (
	"container/list"
	"fmt"
	"reflect"
)

// analyse the instance of struct
func Analyse(instance interface{}, linfo *list.List) *list.List {
	// linfo := list.New()
	typeOf := reflect.TypeOf(instance)
	if Exist(linfo, typeOf) {
		return linfo
	}
	kind := typeOf.Kind()
	linfo.PushFront(typeOf)
	switch kind {
	case reflect.Ptr:
		linfoPtr := AnalysePtr(instance, linfo)
		Join(linfo, linfoPtr)
		return linfo
	case reflect.Struct:
		// linfoLoop := Analyse(instance, linfo)
		// Join(linfo, linfoLoop)
		// return linfo
		fmt.Println("this is right for a struct")
	default:
		return linfo
	}
	elem := reflect.ValueOf(instance)
	for i := 0; i < elem.NumField(); i++ {
		t := elem.Field(i).Type()
		linfo.PushFront(t)
	}
	return linfo
}

// analyse the instance of struct  ptr
func AnalysePtr(instance interface{}, linfo *list.List) *list.List {
	// linfo := list.New()
	if Exist(linfo, reflect.TypeOf(instance)) {
		return linfo
	}
	elem := reflect.ValueOf(instance).Elem()
	for i := 0; i < elem.NumField(); i++ {
		t := elem.Field(i).Type()
		linfo.PushFront(t)
	}
	return linfo
}

// check the type exist in the list
func Exist(linfo *list.List, typeOf reflect.Type) bool {
	if 100 < linfo.Len() {
		return true
	}
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
