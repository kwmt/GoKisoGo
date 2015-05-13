package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "alpha"
	s = s + "beta"
	s += "gamma"
	fmt.Println(s)

	arr := []string{"123", "456", "789"}

	fmt.Println(strings.Join(arr,"."))

	str := "This is a pen"

	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToLower(str))

}