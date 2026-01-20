/*
Copyright © 2026 Err0r51
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// logGoGrepCmd represents the logGoGrep command
var logGoGrepCmd = &cobra.Command{
	Use:   "logGoGrep [flags] [file]",
	Short: "Search for patterns in log files",
	Long: `Search for a specific pattern or string in a log file and display matching lines.

The file can be specified either via the -f/--file flag or as a positional argument.
If both are provided, the flag takes precedence.

Examples:
  # Search for "error" in system.log (using positional argument)
  log-go-grep logGoGrep -s "error" system.log

  # Search using file flag
  log-go-grep logGoGrep -s "error" -f system.log

  # Search and display count of matching lines
  log-go-grep logGoGrep -s "warning" -c system.log

  # Combine flags
  log-go-grep logGoGrep --search "error" --count system.log`,
	Example: `  log-go-grep logGoGrep -s "error" system.log
  log-go-grep logGoGrep -s "error" -c system.log`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("logGoGrep called")
		search, _ := cmd.Flags().GetString("search")
		file, _ := cmd.Flags().GetString("file")
		count, _ := cmd.Flags().GetBool("count")
		var number_of_lines int

		if file == "" {
			if len(args) > 0 {
				file = args[len(args)-1] // Last argument
				fmt.Println("File:", file)
			} else {
				fmt.Println("Error: file path required as argument")
				return
			}
		}
		dat, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		scanner := bufio.NewScanner(strings.NewReader(string(dat)))
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, search) {
				number_of_lines++
				fmt.Println(line)
			}
		}
		if count {
			fmt.Println("Number of lines matching the pattern:", number_of_lines)
		}
	},
}

func init() {
	rootCmd.AddCommand(logGoGrepCmd)
	logGoGrepCmd.Flags().StringP("search", "s", "", "Search string/pattern to find in the log file (required)")
	logGoGrepCmd.MarkFlagRequired("search")
	logGoGrepCmd.Flags().StringP("file", "f", "", "Path to the log file to search (optional if file is provided as positional argument)")
	logGoGrepCmd.Flags().BoolP("count", "c", false, "Display the count of matching lines at the end")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logGoGrepCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logGoGrepCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
