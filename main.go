package main

import (
	"fmt"
	"log"
	"runtime"

	"go.uber.org/automaxprocs/maxprocs"
)

func main() {
	if _, err := maxprocs.Set(); err != nil {
		log.Fatal(err)
	}

	fmt.Print(runtime.GOMAXPROCS(0))
}
