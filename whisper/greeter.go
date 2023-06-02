package main

import (
	"fmt"
	"runtime"
)

func greet() string {
	return fmt.Sprintf("Hello, world from %s %s %s", runtime.Version(), runtime.GOOS, runtime.GOARCH)
}