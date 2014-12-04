package main

import (
	"fmt"
	. "github.com/shaalx/sstruct/structs"
	// "reflect"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	// "encoding/json"
	"bytes"
)

type Structs_type Structs

func main() {
	// itype(Structs)
	jsonF()
	xmlF()
	gobF()
	fmt.Println([]byte("abcdefh"))
}

func itype(t interface{}) {
	// typ := reflect.TypeOf(t)
	// fmt.Println(typ)
	// error

	// i := new(t)
	// fmt.Println(i)
	// error

	// ss := Structs(t{})
	// fmt.Println(ss)
	// error
}

func jsonF() {
	ss := Structs{nil, 1, 2, 3, ".jpg", "/admin/a/b", 4}
	bs, _ := json.Marshal(ss)
	jsons := string(bs)
	fmt.Println(jsons)
}

func xmlF() {
	ss := Structs{nil, 1, 2, 3, ".jpg", "/admin/a/b", 4}
	bs, _ := xml.Marshal(ss)
	jsons := string(bs)
	fmt.Println(jsons)
}

func bsonF() {
	ss := Structs{nil, 1, 2, 3, ".jpg", "/admin/a/b", 4}
	bs, _ := json.Marshal(ss)
	jsons := string(bs)
	fmt.Println(jsons)
}

func gobF() {
	var b bytes.Buffer
	ss := Structs{nil, 1, 2, 3, ".jpg", "/admin/a/b", 4}
	encoder := gob.NewEncoder(&b)
	encoder.Encode(ss)
	fmt.Println(b.String())

	var ss2 Structs
	decoder := gob.NewDecoder(&b)
	decoder.Decode(&ss2)
	fmt.Println(ss2)
}
