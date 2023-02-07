package main

import (
	"go-blog/core"
	"go-blog/initialize"
)

func main() {
	initialize.InitAll()
	core.RunServer()
}
