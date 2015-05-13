package main

import "fmt"
import "time"

func main(){
	c:=make(chan int)
	go func(s chan<- int) {
		for i:=0; i<5; i++ {
			s <- i
			time.Sleep(1*time.Second)
		}
		close(s)
	}(c)

	for {
		value, ok := <- c
		if !ok {
			break
		}
		fmt.Println(value)
	}
}