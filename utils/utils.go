package utils

import (
	"fmt"
	"runtime"
)

func Print(data ...any) {
	fmt.Println(data...)
}

func GetOS() string {
	return runtime.GOOS
}