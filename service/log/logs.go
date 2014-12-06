package log

import (
	"../../pkg3/logs"
)

var (
	LOGS = new(logs.BeeLogger)
)

func init() {
	LOGS = logs.NewLogger(10)
	LOGS.SetLogger("console", "")
	LOGS.Notice("init log")
}

func IsError(tip string, err error) bool {
	if err == nil {
		return false
	}
	LOGS.Error("%s !! Error info : %s", tip, err.Error())
	return true
}
