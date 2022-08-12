package main

import (
	"fmt"

	"github.com/bobrusch/go-solutions/section3/errwrap"
)

func main() {
	errwrap.Wrap()
	fmt.Println()
	errwrap.Unwrap()
	fmt.Println()
	errwrap.StackTrace()
}
