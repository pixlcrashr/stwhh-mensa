// Package cmd
// Copyright © 2024 pixlcrashr (Vincent Heins)
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "stwhh-mensa",
	Short: "CLI-tool for crawling food data of the STWHH (Studierendenwerk Hamburg).",
	Long: `This CLI-tool provides two main parts:
- a crawler to save the food data data in a database
- a GraphQL API to fetch food data for further inspection
- an RSS feed/webpage for receiving updates on foods and meals
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
