package serv

import (
	"testing"
)

func TestServe(t *testing.T) {
	var service Serv
	service = new(ServA)
	service.Init()
	service.Serve()

	service = &(ServB{})
	service.Init()
	service.Serve()
	service.Close()
}
