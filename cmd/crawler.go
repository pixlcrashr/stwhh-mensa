// Package cmd
// Copyright © 2024 pixlcrashr (Vincent Heins)
package cmd

import (
	"github.com/pixlcrashr/stwhh-mensa/pkg/crawler"
	db2 "github.com/pixlcrashr/stwhh-mensa/pkg/db"
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
		cr := crawler.NewCrawler()

		db, err := db2.FromConfig(
			c.Database,
		)
		if err != nil {
			logger.Fatal("could not open database")
		}

		s, err := storage.New(db)
		if err != nil {
			logger.Fatal("could not open database", zap.Error(err))
		}

		scheduler := crawler.NewScheduler(
			cr,
			s,
			logger,
		)

		if err := scheduler.StartAndBlock(time.Hour * 4); err != nil {
			logger.Fatal("could not open database", zap.Error(err))
		}
	},
}

func init() {
	rootCmd.AddCommand(crawlerCmd)
}
