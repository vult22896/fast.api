package main

import (
	"fmt"

	"fast.bibabo.vn/lib"
	"fast.bibabo.vn/server"
)

func main() {
	err := lib.SlackLog("Hello World")
	fmt.Println(err)
	server.Init()
}
