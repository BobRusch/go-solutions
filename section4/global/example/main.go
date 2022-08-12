package main

import "github.com/bobrusch/go-solutions/section4/global"

func main() {
	if err := global.UseLog(); err != nil {
		panic(err)
	}
}
