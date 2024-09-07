package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "frank",
	Short: "Passthrough to Claude Sonnet 3.5",
	Long: `
	Simple wrapper for the Claude Sonnet 3.5 model.
		Expects a config at ~/.frank.yaml. Example command

		> frank what is the distance to the moon?

		> The moon is 238,855 miles from the earth.`,
	Run: func(cmd *cobra.Command, args []string) {
		scanner := bufio.NewScanner(os.Stderr)

		fmt.Println("Enter 'q' to quit.")
		for {
			fmt.Print("> ")

			scanner.Scan()
			input := strings.TrimSpace(scanner.Text())

			if input == "q" {
				fmt.Println("Goodbye...")
				break
			}

			fmt.Printf("You said \"%s\"\n", input)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
