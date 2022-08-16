package main

import "github.com/bobrusch/go-solutions/section6/rest"

func main() {
	if err := rest.Exec(); err != nil {
		panic(err)
	}
}
