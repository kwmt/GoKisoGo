package main

import (
	"os"
	"log"
)

func main() {
	file, _ := os.OpenFile("test.log", os.O_CREATE| os.O_WRONLY| os.O_APPEND, 0666)
	defer func() {
		file.Close()
	}()

	log.SetOutput(file)
	log.Print("test")

	logger := log.New(file, "[mylogger]", log.LstdFlags)
	logger.Print("test1")
	logger.Print("test2")
}