package service

type Service interface {
	Do() bool
	Fetch(url, ipaddr string) []byte
}
