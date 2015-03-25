package action

import (
	"github.com/shaalx/sstruct/service"
	"github.com/shaalx/sstruct/service/log"
	. "github.com/shaalx/sstruct/vars"
	"time"
)

var DATE int64
var logMapping map[int64]string

func init() {
	logMapping = make(map[int64]string, 1)
	logMapping[1426986912] = "C19-Computer0006.txt"
	logMapping[1426989715] = "C34-Economy0002.txt"
	logMapping[1426988787] = "C16-Electronics02.txt"
	logMapping[1426943207] = "file2.txt"
	logMapping[1426943167] = "file1.txt"
	logMapping[1427166719] = "file.txt"
	logMapping[1427101485] = "C32-Agriculture0002.txt"
	DATE = 1426943167
	CURRENT_FILENAME = logMapping[DATE]
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
