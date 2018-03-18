package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("OS........: %v %v\n", runtime.GOOS, runtime.GOARCH)
	fmt.Printf("Threads...: %v\n", runtime.GOMAXPROCS(-1))
}
