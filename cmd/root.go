// Package cmd
// Copyright © 2024 pixlcrashr (Vincent Heins)
package cmd

import (
	"github.com/pixlcrashr/stwhh-mensa/pkg/config"
	"github.com/pixlcrashr/stwhh-mensa/pkg/logging"
	"go.uber.org/zap"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configPath string = "config.yaml"
var c config.Config
var logger *zap.Logger

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "stwhh-mensa",
	Short: "CLI-tool for crawling food data of the STWHH (Studierendenwerk Hamburg).",
	Long: `This CLI-tool provides two main parts:
- a crawler to save the food data data in a database
- a GraphQL API to fetch food data for further inspection
- an RSS feed/webpage for receiving updates on foods and meals
`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		v := viper.New()
		v.SetConfigName("config")
		v.SetConfigType("yaml")
		v.SetConfigFile(configPath)

		if err := v.ReadInConfig(); err != nil {
			return err
		}

		if err := v.Unmarshal(&c); err != nil {
			return err
		}

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	logger = logging.New()
	defer logger.Sync()

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "config.yaml", "Provides the filepath of the config file.")
}
