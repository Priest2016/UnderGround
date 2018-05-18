package main

import (
	"fmt"
	"os"
	t "tools/src/undergroundTools"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Использование: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}

	target := os.Args[1]
	t.ScanIP(target)
}
