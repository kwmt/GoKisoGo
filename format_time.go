package main 

import (
	"fmt"
	"time"
)

func main (){
	
	loc := time.FixedZone("JST",9)

	time :=time.Now()
	time=time.UTC()
	fmt.Println(time.Location().String())
	fmt.Printf("%04d/%02d/%02d %02d:%02d:%02d\n",
		time.Year(),
		time.Month(),
		time.Day(),
		time.Hour(),
		time.Minute(),
		time.Second(),
		)


	
	time=time.In(loc)

	fmt.Println(time.Location())

}