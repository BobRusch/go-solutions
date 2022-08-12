package main

import (
	"fmt"

	"github.com/bobrusch/go-solutions/section3/log"
)

func main() {
	fmt.Println("basic logging and modification of logger:")
	log.Log()
	fmt.Println("logging 'handled' errors:")
	log.FinalDestination()
}
