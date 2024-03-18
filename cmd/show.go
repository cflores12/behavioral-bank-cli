/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	showTotal bool
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the contents of the CSV or SQLite database created by the program.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("show called")
	},
}

func init() {
	showCmd.PersistentFlags().BoolVar(&showTotal, "total", false, "Show total")
	rootCmd.AddCommand(showCmd)
}
