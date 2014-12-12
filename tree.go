package sstruct

import (
	"container/list"
	"fmt"
	"reflect"
	"strings"
)

func tree(valueOf reflect.Value, linfo *list.List) {
	typeOf := valueOf.Type()
	// linfo.PushFront(typeOf)
	switch typeOf.Kind() {
	case reflect.Struct:
		for i := 0; i < valueOf.NumField(); i++ {
			t := valueOf.Field(i).Type()
			// if existInTree(linfo, t) {
			// 	continue
			// }
			linfo.PushFront(t)
			tree(valueOf.Field(i), linfo)
		}
		linfo.PushFront("\tend of" + typeOf.String())
	default:
		// linfo.PushFront(typeOf)
	}
}

// check the type exist in the list
func existInTree(linfo *list.List, typeOf reflect.Type) bool {
	for e := linfo.Back(); e != nil; e = e.Prev() {
		etype, ok := e.Value.(reflect.Type)
		if ok && typeOf == etype {
			return true
		}
	}
	return false
}

func TreeFish(valueOf reflect.Value, linfo *list.List) *list.List {
	typeOf := valueOf.Type()
	fish := list.New()
	fish.PushFront(typeOf)
	switch typeOf.Kind() {
	case reflect.Struct:
		for i := 0; i < valueOf.NumField(); i++ {
			fishtree := TreeFish(valueOf.Field(i), linfo)
			fish.PushFront(fishtree)
		}
	default:
	}
	linfo.PushFront(fish)
	return fish
}

func ShowTreeFish(fishtree *list.List, length int) {
	for e := fishtree.Back(); e != nil; e = e.Prev() {
		childFish, ok := e.Value.(*list.List)
		if ok {
			ShowTreeFish(childFish, length+1)
		} else {
			fmt.Print(strings.Repeat("\t", length))
			fmt.Println(e.Value)
		}
	}
}
