package main

import (
	"DDoSModule/src/DDoSCore"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./ddos_module http://example.com rps")
		os.Exit(100)
	}

	rps, err := strconv.Atoi(os.Args[2])

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(101)
	}

	site := DDoSCore.Site(os.Args[1], rps)
	//site.FindPages()
	site.Start()
}
