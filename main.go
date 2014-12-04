package main

import (
	"fmt"
	"github.com/shaalx/sstruct/fetch"
	"github.com/shaalx/sstruct/mgodb"
	// "labix.org/v2/mgo/bson"
	"encoding/json"
)

func main() {
	url := "http://news-at.zhihu.com/api/3/theme/12"
	// url := "http://fanyi.youdao.com/openapi.do?keyfrom=sasfasdfasf&key=1177596287&type=data&doctype=json&version=1.1&q=love"
	// url := "https://itunes.apple.com/WebObjects/MZStore.woa/wa/viewRoom?fcId=560264462&genreIdString=36&mediaTypeString=Mobile+Software+Applications"
	ipaddr := "202.120.87.152"
	bs := fetch.Do(url, ipaddr)
	fmt.Println(string(bs))
	test(bs)
}

// bs := bson.M{"a": 1}
// fmt.Println(bs)

func test(b []byte) {
	server := []string{"", "newsmgo", "firstbanner"}
	dbserver := mgodb.SetLocalDB(server...)
	var i interface{}
	json.Unmarshal(b, &i)
	dbserver.Save(i)
}
