package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	api "github.com/jburns24/frank/api"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	key     string
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.frank.yaml)")
	rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("license", "apache")
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".frank" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".frank")
		viper.SetConfigType("yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}

	key = viper.GetString("claude.api_key")
	if key == "" {
		fmt.Println("No api key found")
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "frank",
	Short: "Passthrough to Claude Sonnet 3.5",
	Long: `
Simple wrapper for the Claude Sonnet 3.5 model.
	Expects a config at ~/.frank.yaml. Example command:

	$ frank
	> what is the distance to the moon?

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

			resp := api.SendChat(key, input)

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
