package main

import "fmt"

type Calculator interface {
	Calculate(a int, b int) int
}

type Add struct {
	//フィールドは持たない
}

func (x Add) Calculate(a int, b int) int {
	return a + b
}

type Sub struct {
	//フィールドは持たない
}

func (x Sub) Calculate(a int, b int) int {
	return a - b
}


func main() {
	var add Add
	var sub Sub

	var cal Calculator

	cal = add

	calAdd := cal.Calculate(1, 2)
	fmt.Println("和:",calAdd)

	cal = sub

	calSub := cal.Calculate(1, 2)
	fmt.Println("差:",calSub)

}