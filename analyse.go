package sstruct

import (
	"container/list"
	"fmt"
	"reflect"
)

// analyse the instance of struct
func Analyse(instance interface{}, linfo *list.List) *list.List {
	typeOf := reflect.TypeOf(instance)
	if Exist(linfo, typeOf) {
		return linfo
	}
	kind := typeOf.Kind()
	switch kind {
	case reflect.Ptr:
		AnalysePtr(instance, linfo)
		return linfo
	case reflect.Struct:
		linfo.PushFront(typeOf)
		elem := reflect.ValueOf(instance)
		for i := 0; i < elem.NumField(); i++ {
			t := elem.Field(i).Type()
			// 	switch t.Kind() {
			// 	case reflect.Ptr, reflect.Struct:
			// 		Analyse(instance, linfo)
			// 	default:
			linfo.PushFront(t)
			// 	}
		}
	default:
		return linfo
	}
	return linfo
}

// analyse the instance of struct  ptr
func AnalysePtr(instance interface{}, linfo *list.List) *list.List {
	if Exist(linfo, reflect.TypeOf(instance)) {
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
