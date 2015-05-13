package main 

import (
	"fmt"
	"strconv"
)

func main() {
	var i64 int64 = 1234567890

	fmt.Println(strconv.FormatInt(i64,10))

	fmt.Println(strconv.FormatInt(i64,16))
}