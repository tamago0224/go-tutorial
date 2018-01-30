package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go run on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		//freedbsd, openbsd,
		//plan9, windows...
		fmt.Printf("%s.", os)
	}
}
