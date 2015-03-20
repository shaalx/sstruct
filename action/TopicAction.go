package action

import (
	"github.com/shaalx/sstruct/service"
)

func TopicAction() {
	var serv service.Service
	serv = &service.TopicAction{}
	serv.Init()
	// serv.Persistence()
	// serv.Search()
	serv.Analyse(33)
	defer serv.Close()
}
