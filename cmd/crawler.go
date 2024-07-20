/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
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

		s, err := storage.New("test.sqlite")
		if err != nil {
			panic(err)
		}

		for _, d := range res.Dishes {
			if err := s.AddDish(d); err != nil {
				panic(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(crawlerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// crawlerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// crawlerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
