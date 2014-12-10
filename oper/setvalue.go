package oper

import (
	"fmt"
	"reflect"
)

func SetValue(instance interface{}, value []interface{}) interface{} {
	kind := reflect.TypeOf(instance).Kind()
	if reflect.Ptr == kind {
		return SetValueOfPtr(instance, value)
	}
	return SetValueOfCopy(instance, value)
}

// set value of instance
func SetValueOfCopy(instance interface{}, value []interface{}) interface{} {
	kind := reflect.TypeOf(instance).Kind()
	if reflect.Ptr == kind {
		return nil
	}
	newInstance := reflect.New(reflect.TypeOf(instance))
	app := newInstance.Interface()
	elem := reflect.ValueOf(app).Elem()
	fmt.Println(elem.NumField())
	for i, v := range value {
		elem.Field(i).Set(reflect.ValueOf(v))
	}
	return app
}

// set value of ptr
func SetValueOfPtr(instance interface{}, value []interface{}) interface{} {
	kind := reflect.TypeOf(instance).Kind()
	if reflect.Ptr != kind {
		return nil
	}
	elem := reflect.ValueOf(instance).Elem()
	fmt.Println(elem.NumField())
	for i, v := range value {
		elem.Field(i).Set(reflect.ValueOf(v))
	}
	return instance
}

func SetValueAtI(instance interface{}, i int, value interface{}) bool {
	kind := reflect.TypeOf(instance).Kind()
	if reflect.Ptr != kind {
		return false
	}
	reflect.ValueOf(instance).Elem().Field(i).Set(reflect.ValueOf(value))
	return true
}
