package main

import (
	"codesearch/server"
	"fmt"
	"os"
)


func main() {
	err := server.NewServer()
	if err != nil {
		fmt.Println("An error occured while launching the server : ", err)
		os.Exit(1)
	}
}
