package main

import (
	"codesearch/cmd"
	"fmt"
	"os"
)

func main() {
	if err := cmd.HandleCmd(); err != nil {
		fmt.Println("An error occured while handling subcommand : ", err)
		os.Exit(1)
	}
}
