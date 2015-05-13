package main

import (
	"fmt"
	"os"
//	"runtime"
)

func main() {
	go func() {
		fmt.Println("Gorutine")
		os.Exit(0)
	}()

	for {
		//runtime.Gosched()
		fmt.Println("無限ループ")
	}
}