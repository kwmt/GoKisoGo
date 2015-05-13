package main

import (
	"fmt"
    "os"
//	"time"
)

func main() {
	go func() {
		fmt.Println("Gorutine")
		os.Exit(0)
	}()
	for {
//	for i:=0;;i++{
//		fmt.Println(i)
//		time.Sleep(1 * time.Second)
	}

}
