package main

import (
	"fmt"
	"runtime"
)

//除非以 fallthrough 语句结束，否则分支会自动终止。
func main() {
	fmt.Println("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}
}
