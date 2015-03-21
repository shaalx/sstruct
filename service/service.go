package service

type Service interface {
	Init()
	Persistence()
	PersistenceWithUnixDate(date int64)
	Analyse(int)
	AnalyseWithUnixDate(date int64)
	Search()
	Close()
	Log(date int64)
}
