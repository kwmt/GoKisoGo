package main

import (
	"container/list"
	"fmt"
)

func main(){
	l := list.New()
	for i:=0; i < 5; i++ {
		l.PushBack(i)
	}

	fmt.Println("要素数(Init前)：", l.Len())
	l.Init()

	fmt.Println("要素数(Init後)：", l.Len())
}