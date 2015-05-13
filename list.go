package main

import (
	"container/list"
	"fmt"
)

func main(){
	l := list.New() //initialized *List

	fmt.Println(l.Back())
	
	l.PushBack("a")
	l.PushBack("b")
	l.PushBack(3)

	fmt.Println(l.Len())

	elem := l.Front()

	fmt.Println(elem)
	fmt.Println(l.Front().Next())
	fmt.Println(l.Front().Next().Next())

	if l.Front().Next().Next().Next() == nil {
		fmt.Println("nilです")
	}else{
		fmt.Println("じゃない")
	}


	for e:= l.Front(); e != nil; e=e.Next() {
		fmt.Println(e.Value)
	}
}

