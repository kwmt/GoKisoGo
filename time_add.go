package main

import (
	"fmt"
	"time"
)

func main(){
	t := time.Now()

	fmt.Println(t)
	
	add := t.Add(time.Hour*24 + time.Minute*30 + time.Second*5)

	fmt.Println(add)

	t= time.Now()
	add = t.AddDate(2147481635,2147483647,2147483647)
	fmt.Println(add)

}