package bean

// import (
// 	"fmt"
// )

// type Topic struct {
// 	Id     int64
// 	Const  string
// 	Relate string
// 	Parent int64
// 	Weight float32
// 	Freq   int32
// }

// type TopicMap map[int64]Topic

// // 句法成分是否为指定条件
// func (t Topic) isPicked(parent int64, relate ...string) bool {
// 	if -2 == parent {
// 		goto goodDaddy
// 	}
// 	if parent != t.Parent {
// 		return false
// 	}
// goodDaddy:
// 	for _, it := range relate {
// 		if t.Relate == it {
// 			return true
// 		}
// 	}
// 	return false
// }

// // 增加权重
// func (t *Topic) WeightUp(w float32) {
// 	t.Weight += w
// }

// func (t Topic) String() string {
// 	// tStr := fmt.Sprintf("Id %d\t ,Const %s\t ,Relate %s ,Parent %d\n", t.Id, t.Const, t.Relate, t.Parent)
// 	tStr := fmt.Sprintf("%d\t %s\t %d\t %s\t%.3f\n", t.Id, t.Relate, t.Parent, t.Const, t.Weight)
// 	return tStr
// }

// // 排序
// type TopicSlice []*Topic

// func (c TopicSlice) Len() int {
// 	return len(c)
// }

// func (c TopicSlice) Less(i, j int) bool {
// 	return c[i].Id < c[j].Id
// }

// func (c TopicSlice) Swap(i, j int) {
// 	c[i], c[j] = c[j], c[i]
// }

// func (t *TopicSlice) Contain(topic *Topic) bool {
// 	for _, it := range *t {
// 		if it.Id == topic.Id {
// 			return true
// 		}
// 	}
// 	return false
// }

// // 排除重复值
// func (t *TopicSlice) EjRepeat() *TopicSlice {
// 	length := len(*t)
// 	ejMap := make(map[int64]*Topic, length)
// 	for _, it := range *t {
// 		ejMap[it.Id] = it
// 	}
// 	i := 0
// 	result := make(TopicSlice, len(ejMap))
// 	for _, v := range ejMap {
// 		result[i] = v
// 		i++
// 	}
// 	return &result
// }
