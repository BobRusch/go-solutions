package main

import (
	"fmt"

	"github.com/bobrusch/go-solutions/section3/basicerrors"
)

func main() {
	basicerrors.BasicErrors()

	err := basicerrors.SomeFunc()
	fmt.Println("custom error: ", err)
}
