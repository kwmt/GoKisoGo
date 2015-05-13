package main 

import (
	"fmt"
	"math"
)

func main() {
	var degree float64 = 45

	radian := degree * math.Pi / 180

	fmt.Println("radian:",radian)

	fmt.Println("sin:", math.Sin(radian))
	fmt.Println("cos:",math.Cos(radian))
	fmt.Println("tan:",math.Tan(radian))

}