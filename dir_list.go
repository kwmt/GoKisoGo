package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
)

func main(){
	goroot := runtime.GOROOT()
	fmt.Println(goroot)
	
	fileinfos,_ := ioutil.ReadDir(goroot)

	for _, fileinfo := range fileinfos {
		//if !fileinfo.IsDir() {
			fmt.Println(fileinfo.Name())
		//}
	}
}
	