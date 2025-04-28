package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	config         string
	displayVersion bool

	// version number of dom
	cliVersion string
)

var rootCmd = &cobra.Command{
	Use:   "dom",
	Short: "dom is a configurable, symlink-based dotfile manager",
	Long:  `dom (Polish for home) is a configurable, symlink-based dotfile manager. It helps you back up and reuse the most important files in your $HOME directory with ease.`,
	Run: func(cmd *cobra.Command, _ []string) {
		if displayVersion {
			fmt.Println(cliVersion)
			return
		}
		_ = cmd.Help()
	},
}

func Execute(version string) {
	cliVersion = version
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&config, "config", "c", "", "config file path")
	rootCmd.Flags().BoolVar(&displayVersion, "version", false, "version")
}
