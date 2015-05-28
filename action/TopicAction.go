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

	// logMapping[1427352668] = "C34-Economy0003.txt"
	// logMapping[1427353080] = "C34-Economy0008.txt"
	// logMapping[1427353406] = "C34-Economy0012.txt"
	// logMapping[1427353205] = "C34-Economy0010.txt"
	// logMapping[1427352751] = "C34-Economy0004.txt"
	// logMapping[1427352869] = "C34-Economy0006.txt"
	// logMapping[1427353588] = "C34-Economy0014.txt"
	// logMapping[1427353698] = "C34-Economy0016.txt"
	// logMapping[1427353813] = "C34-Economy0018.txt"
	// logMapping[1427354017] = "C34-Economy0020.txt"
	// logMapping[1427354170] = "C34-Economy0022.txt"
	// logMapping[1427354505] = "C34-Economy0024.txt"
	// logMapping[1427354710] = "C34-Economy0026.txt"
	// logMapping[1427468271] = "economy.txt"

	// logMapping[1426989715] = "C34-Economy0002.txt"
	// logMapping[1427770749] = "C19-Computer0010.txt"
	// logMapping[1427770652] = "C11-Space0027.txt"
	// logMapping[1427770830] = "C3-Art0015.txt"
	// logMapping[1427774288] = "C31-Enviornment0005.txt"
	logMapping[1427771023] = "C32-Agriculture0013.txt"
	// logMapping[1427771122] = "C38-Politics0014.txt"
	// logMapping[1427771284] = "C39-Sports0017.txt"
	// logMapping[1427788767] = "C6-Philosophy20.txt"
	// logMapping[1427771457] = "C7-History002.txt"

	logMapping[1428652919] = "C19-Computer0001.txt"
	logMapping[1428652919] = "C19-Computer0004.txt"
	logMapping[1428653198] = "C19-Computer0012.txt"
	logMapping[1428653350] = "C19-Computer0018.txt"
	logMapping[1428653674] = "C19-Computer0020.txt"
	logMapping[1428653915] = "C19-Computer0036.txt"
	logMapping[1428654104] = "C19-Computer0038.txt"
	logMapping[1428654375] = "C19-Computer0044.txt"
	logMapping[1428654517] = "C19-Computer0048.txt"
	logMapping[1428654721] = "C19-Computer0058.txt"
	logMapping[1428654871] = "C19-Computer0062.txt"

	logMapping[1428655413] = "C32-Agriculture0002.txt"
	logMapping[1428655564] = "C32-Agriculture0005.txt"
	logMapping[1428655845] = "C32-Agriculture0007.txt"
	logMapping[1428656176] = "C32-Agriculture0011.txt"
	logMapping[1428656296] = "C32-Agriculture0015.txt"
	logMapping[1428656537] = "C32-Agriculture0017.txt"
	logMapping[1428656894] = "C32-Agriculture0019.txt"
	logMapping[1428657148] = "C32-Agriculture0023.txt"
	logMapping[1428657572] = "C32-Agriculture0027.txt"
	logMapping[1428657859] = "C32-Agriculture0031.txt"
	logMapping[1428658151] = "C32-Agriculture0035.txt"

	logMapping[1428660297] = "C3-Art0003.txt"
	logMapping[1428660611] = "C3-Art0005.txt"
	logMapping[1428660978] = "C3-Art0007.txt"
	logMapping[1428661254] = "C3-Art0011.txt"
	logMapping[1428661714] = "C3-Art0017.txt"
	logMapping[1428662023] = "C3-Art0021.txt"
	logMapping[1428662429] = "C3-Art0023.txt"
	logMapping[1428662991] = "C3-Art0029.txt"
	logMapping[1428663342] = "C3-Art0035.txt"
	logMapping[1428663754] = "C3-Art0037.txt"

	logMapping[1428665231] = "C7-History004.txt"
	logMapping[1428665595] = "C7-History010.txt"
	logMapping[1428666105] = "C7-History014.txt"
	logMapping[1428666377] = "C7-History018.txt"
	logMapping[1428666686] = "C7-History021.txt"
	logMapping[1428666878] = "C7-History029.txt"
	logMapping[1428667033] = "C7-History037.txt"
	logMapping[1428667234] = "C7-History039.txt"
	logMapping[1428667552] = "C7-History041.txt"
	logMapping[1428667949] = "C7-History045.txt"

	logMapping[1428668619] = "C31-Enviornment0003.txt"
	logMapping[1428668938] = "C31-Enviornment0007.txt"
	logMapping[1428669332] = "C31-Enviornment0009.txt"
	logMapping[1428669526] = "C31-Enviornment0011.txt"
	logMapping[1428669872] = "C31-Enviornment0013.txt"
	logMapping[1428670300] = "C31-Enviornment0019.txt"
	logMapping[1428670495] = "C31-Enviornment0025.txt"
	logMapping[1428670850] = "C31-Enviornment0029.txt"
	logMapping[1428671130] = "C31-Enviornment0035.txt"
	logMapping[1428671325] = "C31-Enviornment0039.txt"

	logMapping[1428708617] = "C38-Politics0002.txt"
	logMapping[1428708717] = "C38-Politics0004.txt"
	logMapping[1428708833] = "C38-Politics0008.txt"
	logMapping[1428708955] = "C38-Politics0010.txt"
	logMapping[1428709057] = "C38-Politics0012.txt"
	logMapping[1428709208] = "C38-Politics0019.txt"
	logMapping[1428709289] = "C38-Politics0021.txt"
	logMapping[1428709461] = "C38-Politics0029.txt"
	logMapping[1428709560] = "C38-Politics0033.txt"
	logMapping[1428709656] = "C38-Politics0039.txt"
	logMapping[1428709810] = "C38-Politics0041.txt"
	logMapping[1428709952] = "C38-Politics0059.txt"
	logMapping[1428710072] = "C38-Politics0067.txt"

	logMapping[1428724255] = "C11-Space0003.txt"
	logMapping[1428724366] = "C11-Space0007.txt"
	logMapping[1428724422] = "C11-Space0009.txt"
	logMapping[1428724517] = "C11-Space0015.txt"
	logMapping[1428724705] = "C11-Space0017.txt"
	logMapping[1428724788] = "C11-Space0025.txt"
	logMapping[1428724933] = "C11-Space0029.txt"
	logMapping[1428725035] = "C11-Space0037.txt"
	logMapping[1428725208] = "C11-Space0043.txt"
	logMapping[1428725462] = "C11-Space0051.txt"

	logMapping[1428738766] = "C39-Sports0001.txt"
	logMapping[1428738905] = "C39-Sports0003.txt"
	logMapping[1428739048] = "C39-Sports0005.txt"
	logMapping[1428739184] = "C39-Sports0009.txt"
	logMapping[1428739344] = "C39-Sports0013.txt"
	logMapping[1428739446] = "C39-Sports0019.txt"
	logMapping[1428739581] = "C39-Sports0023.txt"
	logMapping[1428739677] = "C39-Sports0027.txt"
	logMapping[1428739947] = "C39-Sports0029.txt"
	logMapping[1428740191] = "C39-Sports0041.txt"
	logMapping[1428740349] = "C39-Sports0043.txt"
	logMapping[1428740549] = "C39-Sports0047.txt"
	logMapping[1428740701] = "C39-Sports0055.txt"

	logMapping[1428742389] = "C6-Philosophy22.txt"
	logMapping[1428742512] = "C6-Philosophy32.txt"
	logMapping[1428742806] = "C6-Philosophy39.txt"
	logMapping[1428744077] = "C6-Philosophy47.txt"
	logMapping[1428741958] = "C6-Philosophy08.txt"
	logMapping[1428742199] = "C6-Philosophy14.txt"
	logMapping[1428743084] = "C6-Philosophy43.txt"
	logMapping[1428751313] = "C6-Philosophy71.txt"
	logMapping[1428751376] = "C6-Philosophy80.txt"
	logMapping[1428751644] = "C6-Philosophy86.txt"

	logMapping[1432641887] = "C19-Computer0001.txt"
	logMapping[1432641960] = "C19-Computer0004.txt"
	logMapping[1432642047] = "C19-Computer0010.txt"
	logMapping[1432642098] = "C19-Computer0012.txt"
	logMapping[1432642141] = "C19-Computer0018.txt"
	logMapping[1432642227] = "C19-Computer0020.txt"
	logMapping[1432642289] = "C19-Computer0036.txt"
	logMapping[1432642352] = "C19-Computer0038.txt"
	logMapping[1432642423] = "C19-Computer0044.txt"
	logMapping[1432642481] = "C19-Computer0048.txt"
	logMapping[1432642624] = "C19-Computer0058.txt"
	logMapping[1432642668] = "C19-Computer0062.txt"

	DATE = 1427771023
	CURRENT_FILENAME = logMapping[DATE]
}

func TopicAction() {
	// TopicActionPersistence()
	// TopicActionAnalyseN()
	// TopicAction_PersistenceWithUnixDate() // 获得分词
	TopicAction_AnalyseWithUnixDate() // 分析结果
	// AutoPersistence()
	// AutoAnalyse()
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

	// ORIGIN_DIR/economy;computer/agriculture;art;history;environment;politics;space;sports;
	ORIGIN_DIR = ORIGIN_DIR + "philosophy/"
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
	serv.Analyse(1)
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
