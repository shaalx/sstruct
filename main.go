package main

import (
	"github.com/shaalx/sstruct/action"
	// "time"
)

func main() {
	// action.TopicAction() // 测试
	action.TopicAction_PersistenceWithUnixDate() // 获得分词
	// action.TopicAction_AnalyseWithUnixDate() // 分析结果
}
