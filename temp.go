package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	dir := os.TempDir()
	fmt.Println(dir)

	file,_:= ioutil.TempFile(dir,"temp")
	fmt.Println(file.Name())
}