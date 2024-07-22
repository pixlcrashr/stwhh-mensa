// Package cmd
// Copyright © 2024 pixlcrashr (Vincent Heins)
package cmd

import (
	"context"
	"github.com/pixlcrashr/stwhh-mensa/pkg/crawler"
	"github.com/pixlcrashr/stwhh-mensa/pkg/storage"
	"github.com/spf13/cobra"
)

// crawlerCmd represents the crawler command
var crawlerCmd = &cobra.Command{
	Use:   "crawler",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		c := crawler.NewCrawler()

		res, err := c.Crawl(context.Background())
		if err != nil {
			panic(err)
		}

		dbPath, err := cmd.PersistentFlags().GetString("db-path")
		if err != nil {
			panic(err)
		}

		s, err := storage.New(dbPath)
		if err != nil {
			panic(err)
		}

		for _, day := range res {
			if err := s.AddDay(context.Background(), day); err != nil {
				panic(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(crawlerCmd)

	crawlerCmd.PersistentFlags().String("db-path", "./data.sqlite", "Path to the SQLite database file.")
}
