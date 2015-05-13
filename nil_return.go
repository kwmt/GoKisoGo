package main

import "fmt"

type MyError string

func (e MyError) Error() string {
	return string(e)
}

func main(){
	err := do()

	if err == nil {
		fmt.Printf("エラー無し\n")
	} else {
		fmt.Printf("エラーあり??: %v\n",err)
	}
}

func do() error {
	//var ret *MyError = nil

	return nil
}