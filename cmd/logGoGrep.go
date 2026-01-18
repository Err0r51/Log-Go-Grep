/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// logGoGrepCmd represents the logGoGrep command
var logGoGrepCmd = &cobra.Command{
	Use:   "logGoGrep",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("logGoGrep called")
		file, err := cmd.Flags().GetString("file")
		if err != nil {
			fmt.Println("Error getting file:", err)
			return
		}
		fmt.Println("File:", file)
		search, err := cmd.Flags().GetString("search")
		if err != nil {
			fmt.Println("Error getting search:", err)
			return
		}
		fmt.Println("Search:", search)

	},
}

func init() {
	rootCmd.AddCommand(logGoGrepCmd)
	logGoGrepCmd.Flags().StringP("file", "f", "", "File to read")
	logGoGrepCmd.Flags().StringP("search", "s", "", "Search for a string")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logGoGrepCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logGoGrepCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
