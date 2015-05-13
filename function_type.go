package main

import "fmt"

func main(){
	var f func(int, int ) int 

	f = func(a, b int) int {
		return a + b
	}

	fmt.Println(f(1,2))
	
	f = multiply

	fmt.Println(f(2,3))
}

func multiply(x, y int) int {
	return x*y
}