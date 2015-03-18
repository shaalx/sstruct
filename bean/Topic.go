package bean

import "fmt"
import (
	"sort"
)

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

type TopicMatix []TopicSlice

func (t *TopicMatix) Print(cells CellSlice) {
	constMap := make(map[string]int32, 1)
	var maxFr int32
	for _, v := range cells {
		constMap[v.Word] = v.Freq
		if maxFr < v.Freq {
			maxFr = v.Freq
		}
		// fmt.Println(v.Word, v.Freq)
	}
	constStr := ""
	var score float32
	str1 := ""
	sentences := make(Sens, len(*t))
	sum_length := 0
	for _, it := range *t {
		sum_length += len(it)
	}
	for i, it := range *t {
		// fmt.Println(i)
		score = 0.0
		str1 = ""
		var sen Sen
		for _, tpoic := range it {
			constStr = tpoic.Const
			fr, ok := constMap[constStr]
			if ok {
				score += float32(fr)/float32(maxFr)*0.1 + tpoic.Weight/float32(len(it))*0.9
				// fmt.Print(score)
				// fmt.Print("-", fr, "=", constStr)
			}
			str1 += tpoic.Const
			sen = Sen{Str: str1, Sum: score, Avg: score}
		}
		sentences[i] = &sen
	}
	sort.Sort(sentences)
	for _, it := range sentences {
		fmt.Println(it)
	}
}

type Sen struct {
	Str string
	Sum float32
	Avg float32
}

func (s *Sen) String() string {
	return fmt.Sprintf("%.3f\t %.3f\t %s\t", s.Avg, s.Sum, s.Str)
}

type Sens []*Sen

func (c Sens) Len() int {
	return len(c)
}

func (c Sens) Less(i, j int) bool {
	return c[i].Avg > c[j].Avg
}

func (c Sens) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
