package persistence

type Persistence struct {
	Do func([]byte) bool
}
