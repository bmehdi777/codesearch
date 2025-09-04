package cmd

import (
	"codesearch/server"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func newStartCmd() *cobra.Command {
	startCmd := cobra.Command{
		Use: "start",
		Run: handlerStart,
	}

	return &startCmd
}

func handlerStart(cmd *cobra.Command, args []string) {
	err := server.NewServer()
	if err != nil {
		fmt.Println("An error occured while launching the server : ", err)
		os.Exit(1)
	}
}
