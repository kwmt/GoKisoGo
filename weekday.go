package main
import (
	"time"
	"fmt"
)


func main(){

	t := time.Now()
	
	w := t.Weekday()

	fmt.Println(w)
}