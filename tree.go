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
	// linfo.PushFront(fish)
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

func SpiltFish(fishRunes []rune) *list.List {
	fmt.Println(string(fishRunes))
	linfo := list.New()
	linfo.PushFront(string(fishRunes))
	end := FirstIndexOf(fishRunes, []rune("}")[0])
	fmt.Println(end)
	if 0 > end {
		return linfo
	}
	start := LastIndexOf(fishRunes[:end], []rune("{")[0])
	fmt.Println(start)
	if 0 > start {
		return linfo
	}
	childFishOne := SpiltFish(fishRunes[start+1 : end])
	// linfo.PushFront(childFishOne)
	left := make([]rune, start+1)
	left = fishRunes[:start]
	left = append(left, fishRunes[end+1:]...)
	childFishTwo := SpiltFish(left)
	childFishTwo.PushFront(childFishOne)
	linfo.PushFront(childFishTwo)
	return linfo
}

func FirstIndexOf(runes []rune, r rune) int {
	for i, it := range runes {
		if it == r {
			return i
		}
	}
	return -1
}

func LastIndexOf(runes []rune, r rune) int {
	for i := len(runes); i > 0; i-- {
		if runes[i-1] == r {
			return i - 1
		}
	}
	return -1
}
