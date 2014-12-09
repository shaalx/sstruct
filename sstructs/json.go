package sstructs

import (
	"encoding/json"
	"fmt"
)

func Hello() {
	fmt.Println("hellp")
}

func Jsonable(i interface{}) {
	b, err := json.Marshal(i)
	if err == nil {
		fmt.Println(string(b))
	} else {
		fmt.Println(err.Error())
	}
}
