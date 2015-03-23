package action

import (
	"github.com/shaalx/sstruct/service"
	"github.com/shaalx/sstruct/service/log"
	. "github.com/shaalx/sstruct/vars"
	"time"
)

var DATE int64

func init() {
	CURRENT_FILENAME = "C19-Computer0006.txt"
	DATE = 1426986912
}
func TopicAction() {
	// TopicActionAnalyseN()
	// TopicAction_PersistenceWithUnixDate() // 获得分词
	TopicAction_AnalyseWithUnixDate() // 分析结果
}

func TopicAction_PersistenceWithUnixDate() {
	start := time.Now()
	var serv service.Service
	serv = &service.TopicAction{}
	serv.Init()
	serv.PersistenceWithUnixDate(start.Unix())
	serv.Log(start.Unix())
	defer serv.Close()

	log.LOGS.Alert("Time costs : %v", time.Now().Sub(start))
	time.Sleep(time.Second)
}

func TopicAction_AnalyseWithUnixDate() {
	start := time.Now()

	var serv service.Service
	serv = &service.TopicAction{}
	serv.Init()
	serv.AnalyseWithUnixDate(DATE)
	defer serv.Close()

	log.LOGS.Alert("Time costs : %v", time.Now().Sub(start))
	time.Sleep(time.Second)
}

func TopicActionAnalyseN() {
	start := time.Now()

	var serv service.Service
	serv = &service.TopicAction{}
	serv.Init()
	// serv.Search()
	serv.Analyse(33)
	defer serv.Close()

	log.LOGS.Alert("Time costs : %v", time.Now().Sub(start))
	time.Sleep(time.Second)
}
