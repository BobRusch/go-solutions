package main

import "github.com/bobrusch/go-solutions/section1/tempfiles"

func main() {
	if err := tempfiles.WorkWithTemp(); err != nil {
		panic(err)
	}
}
