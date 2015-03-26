package action

import (
	"github.com/shaalx/sstruct/service"
	"github.com/shaalx/sstruct/service/log"
	"github.com/shaalx/sstruct/utils"
	. "github.com/shaalx/sstruct/vars"
	"time"
)

var DATE int64
var logMapping map[int64]string

func init() {
	logMapping = make(map[int64]string, 1)

	// logMapping[1426986912] = "C19-Computer0006.txt"
	// logMapping[1426989715] = "C34-Economy0002.txt"
	// logMapping[1426988787] = "C16-Electronics02.txt"
	// logMapping[1426943207] = "file2.txt"
	// logMapping[1426943167] = "file1.txt"
	// logMapping[1427166719] = "file.txt"
	// logMapping[1427101485] = "C32-Agriculture0002.txt"

	logMapping[1427352668] = "C34-Economy0003.txt"
	logMapping[1427352751] = "C34-Economy0004.txt"
	logMapping[1427352869] = "C34-Economy0006.txt"
	logMapping[1427353080] = "C34-Economy0008.txt"
	logMapping[1427353205] = "C34-Economy0010.txt"
	logMapping[1427353406] = "C34-Economy0012.txt"
	logMapping[1427353588] = "C34-Economy0014.txt"
	logMapping[1427353698] = "C34-Economy0016.txt"
	logMapping[1427353813] = "C34-Economy0018.txt"
	logMapping[1427354017] = "C34-Economy0020.txt"
	logMapping[1427354170] = "C34-Economy0022.txt"
	logMapping[1427354505] = "C34-Economy0024.txt"
	logMapping[1427354710] = "C34-Economy0026.txt"

	// logMapping[1427031288] = "C34-Economy0002.txt"
	// logMapping[1427297132] = "file.txt"
	// logMapping[1427297135] = "file1.txt"
	// logMapping[1427297138] = "file2.txt"

	DATE = 1426986912
	CURRENT_FILENAME = logMapping[DATE]
}

func TopicAction() {
	// TopicActionPersistence()
	// TopicActionAnalyseN()
	// TopicAction_PersistenceWithUnixDate() // 获得分词
	// TopicAction_AnalyseWithUnixDate() // 分析结果
	// AutoPersistence()
	AutoAnalyse()
}

func AutoAnalyse() {
	start := time.Now()

	var serv service.Service
	serv = &service.TopicAction{}
	serv.Init()
	defer serv.Close()

	for date, file := range logMapping {
		loopstart := time.Now()
		CURRENT_FILENAME = file
		log.LOGS.Alert("file %v", file)
		serv.AnalyseWithUnixDate(date)
		log.LOGS.Alert("Time costs : %v", time.Now().Sub(loopstart))
	}
	log.LOGS.Alert("Time costs : %v", time.Now().Sub(start))
	time.Sleep(time.Second)
}

func AutoPersistence() {
	start := time.Now()

	var serv service.Service
	serv = &service.TopicAction{}
	serv.Init()
	defer serv.Close()

	// ORIGIN_DIR
	ORIGIN_DIR = ORIGIN_DIR + "economy/"
	files := utils.ReadDir(ORIGIN_DIR)
	for i, file := range files {
		log.LOGS.Alert("%d:\t %s\n", i, file)
		loopstart := time.Now()
		CURRENT_FILENAME = file
		serv.PersistenceWithUnixDate(loopstart.Unix())
		serv.Log(loopstart.Unix())
		log.LOGS.Alert("Time costs : %v", time.Now().Sub(loopstart))
		time.Sleep(1)
	}

	log.LOGS.Alert("Time costs : %v", time.Now().Sub(start))
	time.Sleep(time.Second)
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

func TopicActionPersistence() {
	start := time.Now()

	var serv service.Service
	serv = &service.TopicAction{}
	serv.Init()

	serv.Persistence()
	defer serv.Close()

	log.LOGS.Alert("Time costs : %v", time.Now().Sub(start))
	time.Sleep(time.Second)
}
