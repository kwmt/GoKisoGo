package main 

import (
	"fmt"
	"time"
)

func main (){
	

	
	l,_ := time.LoadLocation("Local")
	t := time.Date(2001,2,3,11,22,33,44,l)
	fmt.Println(t)
}