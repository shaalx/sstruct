package service

type Service interface {
	Init()
	Persistence()
	Analyse()
	Close()
}
