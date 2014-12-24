package serv

type Serv interface {
	Init()
	Serve()
	Close()
}
