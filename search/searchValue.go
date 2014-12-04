package search

import (
	sjson "github.com/shaalx/sstruct/pkg3/go-simplejson"
)

// search value
// 查询某个路径path下的key值 string
func SearchPathSValue(data []byte, key string, path ...string) *string {
	if data == nil {
		return nil
	}
	js, err := sjson.NewJson(data)
	if checkError(err) {
		return nil
	}
	value, err := js.GetPath(path...).Get(key).String()
	if checkError(err) {
		return nil
	}
	return &value
}

// search value
// 查询某个路径path下的key值 int
func SearchPathIValue(data []byte, key string, path ...string) int64 {
	if data == nil {
		return -1
	}
	js, err := sjson.NewJson(data)
	if checkError(err) {
		return -1
	}
	value, err := js.GetPath(path...).Get(key).Int64()
	if checkError(err) {
		return -1
	}
	return value
}

// search key
// 查询某一个key值 string
func SearchSValue(data []byte, key string) *string {
	if data == nil {
		return nil
	}
	js, err := sjson.NewJson(data)
	if checkError(err) {
		return nil
	}
	value, err := js.Get(key).String()
	if checkError(err) {
		return nil
	}
	return &value
}

// search key
// 查询某一个key值 int
func SearchIValue(data []byte, key string) int64 {
	if data == nil {
		return -1
	}
	js, err := sjson.NewJson(data)
	if checkError(err) {
		return -1
	}
	value, err := js.Get(key).Int64()
	if checkError(err) {
		return -1
	}
	return value
}
