package utils

import (
	"../pkg3/mgo/bson"
	"../service/log"
	"encoding/json"
)

// 解码 有 问题
func Bson2Bytes(b *bson.M) []byte {
	bs, err := bson.Marshal(b)
	if log.IsError("{can not convert bson to bytes}", err) {
		return nil
	}
	return bs
}

// in 其实为 *bson.M
func I2Bytes(in interface{}) []byte {
	out, err := json.Marshal(in)
	if log.IsError("{can not convert interface{} to bytes}", err) {
		return nil
	}
	return out
}
