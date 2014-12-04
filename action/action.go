package action

type Action struct {
	Do func() bool
}
