package main

import (
	"flag"
	"fmt"
)

func main() {
	version := flag.Bool("version", false, "get the version")
	v := flag.Bool("v", false, "get the version")
	flag.Parse()
	if *version || *v {
		fmt.Println("0.0.0")
	}
}
