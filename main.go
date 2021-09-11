package main

import (
	"github.com/EwRvp7LV7/1task2/parser"
	"log"
	"os"
	//"fmt"
)

func run() error {

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }

   log.SetOutput(file)
   log.Println("Поехали!")

	parser.Parser()

	return nil
}

func main() {

	if err := run(); err != nil {
		log.Fatal()
	}
}
