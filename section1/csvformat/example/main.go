package main

import (
	"fmt"

	"github.com/bobrusch/go-solutions/section1/csvformat"
)

func main() {
	if err := csvformat.AddMoviesFromText(); err != nil {
		panic(err)
	}

	if err := csvformat.WriteCSVOutput(); err != nil {
		panic(err)
	}

	buffer, err := csvformat.WriteCSVBuffer()
	if err != nil {
		panic(err)
	}

	fmt.Println("Buffer = ", buffer.String())
}
