package main

import (
	"fmt"
	"math/rand"
	"time"
)

const gorutines  = 3

func main() {
	c := make(chan int)

	for i:=0; i < gorutines; i++ {
		go func(s chan <- int){
			duration := time.Duration(rand.Int63n(15)) * time.Second
			//fmt.Println("duration:",duration)
			time.Sleep(duration)
			fmt.Println("処理完了")
			
			c <- 0
		}(c)
			
	}

	for i:=0; i < gorutines; i++ {
		<- c
	}

	fmt.Println("すべて完了")
	
}
