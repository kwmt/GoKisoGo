package main

import "fmt"

func main(){
	fmt.Println("開始")
	defer delay(1)
	defer delay(2)
	fmt.Println("終了")

}

func delay(i int) {
	fmt.Printf("delay %d\n",i)
}
