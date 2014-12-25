package service

type Service interface {
	Init()
	Persistence()
	Analyse()
	Search()
	Close()
}
