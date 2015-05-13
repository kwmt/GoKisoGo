package main

import "fmt"
import "os"

func main() {
	file, err := os.Open("test1.txt")
	defer file.Close()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("OK")
}