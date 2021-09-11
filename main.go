package main

import (
	"github.com/EwRvp7LV7/1task2/parser"
	"log"
	//"fmt"
)

func run() error {
	parser.Parser()

	return nil
}

func main() {

	if err := run(); err != nil {
		log.Fatal()
	}
}
