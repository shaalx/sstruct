package main

import (
	"fmt"
	"labix.org/v2/mgo/bson"
)

func main() {
	bs := bson.M{"a": 1}
	fmt.Println(bs)
}
