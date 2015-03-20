package service

type Service interface {
	Init()
	Persistence()
	Analyse(int)
	Search()
	Close()
}
