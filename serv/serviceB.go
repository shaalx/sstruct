package serv

import (
	"fmt"
)

type ServB struct {
	Name string
}

func (s *ServB) Init() {
	s.Name = "servBice"
	fmt.Println(s)
}

func (s *ServB) Serve() {
	fmt.Println(s.Name, ": serving...")
}
func (s *ServB) Close() {
	fmt.Println(s.Name, ": closing...")
}
