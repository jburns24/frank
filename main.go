package main

import (
	"github.com/jburns24/frank/cmd"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func main() {
	cmd.Execute()
}
