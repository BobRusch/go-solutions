package main

import "github.com/bobrusch/go-solutions/section5/mongodb"

func main() {
	if err := mongodb.Exec("mongodb://proot:Pa55w0rd@localhost:27017"); err != nil {
		panic(err)
	}
}
