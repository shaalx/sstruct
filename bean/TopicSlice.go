package bean

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
