package main

import (
	"flag"
	"fmt"
	"main/parser"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if parser.AssertXml(args[0]) {
		fmt.Println("Valid")
	} else {
		fmt.Println("Invalid")
	}
}
