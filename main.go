package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	"go.uber.org/automaxprocs/maxprocs"
)

func main() {
	v := flag.Bool("v", false, "display GOMAXPROCS before and after or if no change has occurred")
	flag.Parse()

	before := runtime.GOMAXPROCS(0)

	if _, err := maxprocs.Set(); err != nil {
		log.Fatal(err)
	}

	after := runtime.GOMAXPROCS(0)

	if *v {
		if before == after {
			fmt.Printf("GOMAXPROCS: %v (no change)\n", before)
		} else {
			fmt.Printf("GOMAXPROCS before: %v after: %v\n", before, after)
		}

		os.Exit(0)
	}

	fmt.Print(after)
}
