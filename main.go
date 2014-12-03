package main

import (
	"fmt"
	// "labix.org/v2/mgo/bson"
	"github.com/shaalx/sstruct/fetch"
)

func main() {
	url := "http://fanyi.youdao.com/openapi.do?keyfrom=sasfasdfasf&key=1177596287&type=data&doctype=json&version=1.1&q=love"
	// url := "https://itunes.apple.com/WebObjects/MZStore.woa/wa/viewRoom?fcId=560264462&genreIdString=36&mediaTypeString=Mobile+Software+Applications"
	ipaddr := "202.120.87.152"
	bs := fetch.Do(url, ipaddr)
	fmt.Println(string(bs))
}

// bs := bson.M{"a": 1}
// fmt.Println(bs)
