package log

import (
	"fmt"
)

func IsError(err error) bool {
	if err == nil {
		return false
	}
	fmt.Println(err.Error())
	return true
}
