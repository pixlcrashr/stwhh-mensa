// Package cmd
// Copyright © 2024 pixlcrashr (Vincent Heins)
package cmd

import (
	"github.com/pixlcrashr/stwhh-mensa/pkg/crawler"
	"github.com/pixlcrashr/stwhh-mensa/pkg/logger"
	"github.com/pixlcrashr/stwhh-mensa/pkg/storage"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"time"
)

// crawlerCmd represents the crawler command
var crawlerCmd = &cobra.Command{
	Use:   "crawler",
	Short: "Crawls the food overview of the STWHH website in intervals.",
	Run: func(cmd *cobra.Command, args []string) {
		l := logger.New()

		c := crawler.NewCrawler()

		dbPath, err := cmd.PersistentFlags().GetString("db-path")
		if err != nil {
			l.Fatal("could not open database", zap.Error(err))
		}

		s, err := storage.New(dbPath)
		if err != nil {
			l.Fatal("could not open database", zap.Error(err))
		}

		scheduler := crawler.NewScheduler(
			c,
			s,
			l,
		)

		if err := scheduler.StartAndBlock(time.Hour * 4); err != nil {
			l.Fatal("could not open database", zap.Error(err))
		}
	},
}

func init() {
	rootCmd.AddCommand(crawlerCmd)

	crawlerCmd.PersistentFlags().String("db-path", "./db.sqlite", "Path to the SQLite database file.")
}
