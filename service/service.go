package service

type Service interface {
	Init()
	Persistence()
	PersistenceWithUnixDate(date int64)
	Analyse(int)
	Search()
	Close()
}
