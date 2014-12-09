package sstructs

import (
	"fmt"
	"reflect"
	"unsafe"
)

func Alloc(instance interface{}) interface{} {
	if !addressed(instance) {
		return instance
	}
	fmt.Print("alloc : ")
	valueOf := reflect.ValueOf(instance).Elem()
	typeValueOf := valueOf.Type()
	fmt.Println("type ValueOf", typeValueOf)
	newInstance := reflect.New(typeValueOf)
	app := newInstance.Interface()
	fmt.Println(reflect.ValueOf(app), reflect.TypeOf(app))
	fmt.Println(reflect.ValueOf(instance), reflect.TypeOf(instance))
	e := (*interface{})(unsafe.Pointer(&app))
	fmt.Println(*e)
	fmt.Println(reflect.ValueOf(&app).CanSet())
	fmt.Println(reflect.TypeOf(app).Elem().Field(0))
	// reflect.
	app1 := App{Name: "C-i-t-y", Address: "S-h-a-n-g-h-a-i"}
	vp := reflect.ValueOf(&app1).Elem()
	reflect.ValueOf(&app).Elem().Set(vp.Addr())

	reflect.ValueOf(&app).Elem().Set(reflect.ValueOf("itt"))
	// 	vvv := []string{"vvv", "ttt"}
	// vvvp := reflect.ValueOf(&vvv)
	// reflect.ValueOf(&app).Elem().Set(vvvp.Elem())
	return app
}

func Alloc2(instance interface{}) interface{} {
	typeOf := reflect.TypeOf(instance)
	var v reflect.Value
	switch typeOf.Kind() {
	case reflect.Ptr:
		v = reflect.ValueOf(instance).Elem()
	default:
		v = reflect.ValueOf(instance).Field(0)
	}
	fmt.Println(v)
	return instance
}

func Change(i interface{}) interface{} {
	// set with string
	// reflect.ValueOf(&i).Elem().Set(reflect.ValueOf("itt"))

	// set with int
	canSetValue := reflect.ValueOf(&i).Elem()
	canSetValue.Set(reflect.ValueOf(100))
	fmt.Println(i)

	// .Set(*.Addr())
	// fmt.Println(canSetValue.Addr())

	// set with []string
	newValue := reflect.ValueOf([]string{"a", "b"})
	canSetValue.Set(newValue)
	fmt.Println(i)

	newValue2 := reflect.ValueOf(interface{}([]string{"a", "b"}))
	canSetValue.Set(newValue2)
	fmt.Println(i)
	return i
}

// Elem 指针指向的内存内容
// Field 指针不可操作

func ChangeWithI(i, j interface{}) interface{} {

	canSetValue := reflect.ValueOf(&i).Elem()

	newValue2 := reflect.ValueOf(j)
	canSetValue.Set(newValue2)
	fmt.Println(i)
	return i
}

func ChangeWithIJ(i interface{}, j int) interface{} {
	kind := reflect.TypeOf(i).Kind()
	if reflect.Ptr == kind {
		reflect.ValueOf(&i).Elem().Set(reflect.ValueOf("ptr"))
	} else {
		newInstance := reflect.New(reflect.TypeOf(i))
		app := newInstance.Interface()
		elem := reflect.ValueOf(app).Elem()
		elem.Field(j).Set(reflect.ValueOf("inst"))
		return app
	}

	fmt.Println(i)
	return i
}

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
	// reflect.ValueOf(&instance).Elem().Set(reflect.ValueOf(value))
	// return instance

	// newInstance := reflect.New(reflect.ValueOf(instance).Elem().Type())
	// app := newInstance.Interface()
	elem := reflect.ValueOf(instance).Elem()
	for i, v := range value {
		elem.Field(i).Set(reflect.ValueOf(v))
	}
	return instance
}
