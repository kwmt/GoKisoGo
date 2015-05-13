package main

import "fmt"

func main() {
	c := make(chan int, 3)

	c <- 0
	c <- 1
	c <- 2

	fmt.Println("cap:",cap(c))
	fmt.Println("len:",len(c))

	<- c
	<- c
	fmt.Println("cap:",cap(c))
	fmt.Println("len:",len(c))
}