package persistence

type Persistence struct {
	Do func(bytes []byte, notice string) bool
}
