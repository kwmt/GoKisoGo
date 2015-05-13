package main 


import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main () {
	response,err := http.Get("http://golang.jp/hello.html")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
//	fmt.Println(response)

	body,_ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))

	status := response.Status

	fmt.Println(status)
	
}