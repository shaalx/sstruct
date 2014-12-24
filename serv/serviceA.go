package serv

import (
	"fmt"
)

type ServA struct {
	Name string
}

func (s *ServA) Init() {
	s.Name = "servAice"
	fmt.Println(s)
}

func (s *ServA) Serve() {
	fmt.Println(s.Name, ": serving...")
}

func (s *ServA) Close() {
	fmt.Println(s.Name, ": closing...")
}
