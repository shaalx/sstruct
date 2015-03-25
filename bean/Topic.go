package bean

import (
	"fmt"
)

type Topic struct {
	Id     int64
	Const  string
	Relate string
	Parent int64
	Weight float32
	Freq   int32
	Pos    string
}

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

// ATT,ATT,SBV   (ATT)   t.IsCond(topics,[]string{"ATT","ATT","SBV"}...)
func (t Topic) IsCond(topics TopicSlice, relate ...string) (bool, TopicSlice) {
	id := t.Id
	resultTopics := make(TopicSlice, 0)
	for _, it := range relate {
		if id == -1 {
			return true, resultTopics
		}
		tNow := topics[id]
		if tNow.Relate != it {
			return false, nil
		}
		id = tNow.Parent
		resultTopics = append(resultTopics, tNow)
	}
	return true, resultTopics
}

// 增加权重
func (t *Topic) WeightUp(w float32) {
	t.Weight += w
}

func (t Topic) String() string {
	return fmt.Sprintf("%d\t %s\t %d\t %s\t%.3f\n", t.Id, t.Relate, t.Parent, t.Const, t.Weight)
}

func (t *Topic) CutWords() string {
	return t.Const
}
