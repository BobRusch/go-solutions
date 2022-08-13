package main

import (
	"fmt"

	"github.com/bobrusch/go-solutions/section4/panic"
)

func main() {
	fmt.Println("before panic")
	panic.Catcher()
	fmt.Println("after panic")
}
