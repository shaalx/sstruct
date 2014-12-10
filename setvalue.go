package sstruct

import (
	// "fmt"
	"reflect"
)

// set value for an instance
func SetValue(instance interface{}, value []interface{}) interface{} {
	kind := reflect.TypeOf(instance).Kind()
	if reflect.Ptr == kind {
		ok := SetValueOfPtr(instance, value)
		if ok {
			return instance
		} else {
			return nil
		}
	}
	return SetValueOfCopy(reflect.TypeOf(instance), value)
}

// set value of instance
func SetValueOfCopy(typeOf reflect.Type, value []interface{}) interface{} {
	kind := typeOf.Kind()
	if reflect.Ptr == kind {
		return nil
	}
	newInstance := reflect.New(typeOf)
	app := newInstance.Interface()
	elem := reflect.ValueOf(app).Elem()
	// if length of value is over elem.NumField
	if len(value) > elem.NumField() {
		value = value[:elem.NumField()]
	}
	for i, v := range value {
		elem.Field(i).Set(reflect.ValueOf(v))
	}
	return app
}

// set value of ptr
func SetValueOfPtr(instance interface{}, value []interface{}) bool {
	kind := reflect.TypeOf(instance).Kind()
	if reflect.Ptr != kind {
		return false
	}
	elem := reflect.ValueOf(instance).Elem()
	// if length of value is over elem.NumField
	if len(value) > elem.NumField() {
		value = value[:elem.NumField()]
	}
	for i, v := range value {
		elem.Field(i).Set(reflect.ValueOf(v))
	}
	return true
}

// set value for instance at i
func SetValueAtI(instance interface{}, i int, value interface{}) bool {
	kind := reflect.TypeOf(instance).Kind()
	if reflect.Ptr != kind {
		return false
	}
	reflect.ValueOf(instance).Elem().Field(i).Set(reflect.ValueOf(value))
	return true
}
