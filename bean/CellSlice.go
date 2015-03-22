package bean

// 按照统计词频排序
import (
	"fmt"
)

type Cell struct {
	Word string
	Freq int32
}

func (c Cell) String() string {
	return fmt.Sprintf("%d\t %s", c.Freq, c.Word)
}

type CellSlice []*Cell

func (c CellSlice) Len() int {
	return len(c)
}

func (c CellSlice) Less(i, j int) bool {
	return c[i].Freq < c[j].Freq
}

func (c CellSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c CellSlice) String() {
	for i, v := range c {
		if v.Freq < int32(3) {
			continue
		}
		fmt.Printf("%d\t %s\n", i, v.String())
	}
}
