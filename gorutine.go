package main

import(
	"fmt"
	"time"
)

func main(){
	test()
	fmt.Println("ゴルーチンでtestをコール")
	go test()
	time.Sleep(2*time.Second)
}

func test(){
	for i:=0;i<5; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}