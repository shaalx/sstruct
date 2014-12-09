package sstructs

import (
	"encoding/json"
	"fmt"
)

func Hello() {
	fmt.Println("hellp")
}

func Jsonable(i interface{}) []byte {
	b, err := json.Marshal(i)
	if err == nil {
		fmt.Println(string(b))
	} else {
		fmt.Println(err.Error())
	}
	return b
}

func Unjsonable(b []byte) App {
	app := App{}
	if nil == b {
		return app
	}
	err := json.Unmarshal(b, &app)
	if err == nil {
		fmt.Println(app)
	} else {
		fmt.Println(err.Error())
	}

	return app
}
