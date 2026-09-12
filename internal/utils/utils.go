// Package utils: shared utilities
package utils

import (
	"fmt"
	"os"
)

func ErrCheck(e error, msg string) {
	if e != nil {
		fmt.Println(msg)
		os.Exit(1)
	}
}

func ErrorExit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
