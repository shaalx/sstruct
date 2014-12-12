package sstruct

import (
	"fmt"
	"reflect"
)

var (
	mapType map[reflect.Type][]reflect.Type // = map[reflect.Type][]reflect.Type
)

func init() {
	mapType = make(map[reflect.Type][]reflect.Type, 15)
}

func loopAnalyse(value reflect.Value) {
	typeOf := value.Type()
	if Exists(typeOf) {
		fmt.Println(mapType)
		return
	}
	if reflect.Struct == typeOf.Kind() {
		length := value.NumField()
		newTypes := make([]reflect.Type, length)
		for i := 0; i < length; i++ {
			newTypes[i] = value.Field(i).Type()
		}
		mapType[typeOf] = newTypes
		for i := 0; i < length; i++ {
			loopAnalyse(value.Field(i))
		}
	}
}

// check type exist in the mapType
func Exists(typeOf reflect.Type) bool {
	_, ok := mapType[typeOf]
	if ok {
		return true
	}
	return false
}

func PrintMap() {
	for k, v := range mapType {
		fmt.Println(k)
		for i, it := range v {
			fmt.Printf("\t< %d > %v\n", i, it)
		}
	}
}
