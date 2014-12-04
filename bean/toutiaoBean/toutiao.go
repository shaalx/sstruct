package toutiaoBean

// import (
// 	"encoding/json"
// 	"fmt"
// 	"github.com/shaalx/sstruct/mgo/bson"
// )

// // 解码 有 问题
// func Bson2Bytes(b *bson.M) []byte {
// 	bs, err := bson.Marshal(b)
// 	if err != nil {
// 		fmt.Println("binary.Write failed:", err)
// 	}
// 	return bs
// }

// // in 其实为 *bson.M
// func I2Bytes(in interface{}) []byte {
// 	out, err := json.Marshal(in)
// 	if err != nil {
// 		fmt.Println("binary.Write failed:", err)
// 	}
// 	return out
// }
