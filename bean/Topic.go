package bean

import "fmt"

type Topic struct {
	Id     int64
	Const  string
	Relate string
	Parent int64
	Weight float32
	Freq   int32
}

type TopicIdMap map[int64]*Topic
type TopicConstMap map[string]*Topic

// 句法成分是否为指定条件
func (t Topic) IsPicked(parent int64, relate ...string) bool {
	if -2 == parent {
		goto goodDaddy
	}
	if parent != t.Parent {
		return false
	}
goodDaddy:
	for _, it := range relate {
		if t.Relate == it {
			return true
		}
	}
	return false
}

// 增加权重
func (t *Topic) WeightUp(w float32) {
	t.Weight += w
}

func (t Topic) String() string {
	// tStr := fmt.Sprintf("Id %d\t ,Const %s\t ,Relate %s ,Parent %d\n", t.Id, t.Const, t.Relate, t.Parent)
	tStr := fmt.Sprintf("%d\t %s\t %d\t %s\t%.3f\n", t.Id, t.Relate, t.Parent, t.Const, t.Weight)
	return tStr
}
