package error

import (
	"fmt"
)

func isError(err error) bool {
	if err == nil {
		return false
	}
	fmt.Println(err.Error())
	return true
}
