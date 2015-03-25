package utils

import (
	"fmt"
	"testing"
)

// func TestDeleteEnter(t *testing.T) {
// 	DeleteEnter("file.txt")
// }

func TestReadDir(t *testing.T) {
	files := ReadDir(".")
	for i, it := range files {
		fmt.Printf("%d:\t %s\n", i, it)
	}
}
