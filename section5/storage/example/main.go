package main

import "github.com/bobrusch/go-solutions/section5/storage"

func main() {
	if err := storage.Exec(); err != nil {
		panic(err)
	}
}
