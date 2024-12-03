/*
Copyright © 2024 Mike Benson mike@kambe.com.au

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// The following variables are used in every solution
var Test bool

var inputFile string

var testData string

var lines *bufio.Scanner

var solution func(*bufio.Scanner) int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aoc_2024",
	Short: "Advent of Code 2024 in Golang",
	Long: `My solutions to Advent of Code for 2024
This year I have decided to implement the solutions in Go.

To run a solution for a given day, Use go run main.go <day99> <part9>`,
	TraverseChildren: true,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		var file io.Reader
		var err error
		if Test {
			file = strings.NewReader(testData)
		} else {
			file, err = os.Open(inputFile)
			if err != nil {
				fmt.Println(err)
				return
			}	
		}
		lines = bufio.NewScanner(file)
		fmt.Printf("Result is %v\n", solution(lines))
	},
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.aoc_2024.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().BoolVarP(&Test, "test", "t", false, "Use Test Input")
}


