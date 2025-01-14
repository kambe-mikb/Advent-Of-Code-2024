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

var Debug bool

var inputFile string

var testData string

var lines *bufio.Scanner

var solution func(*bufio.Scanner) int

func debugln(message ...any) {
	if Debug {
		fmt.Println(message...)
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aoc_2024",
	Short: "Advent of Code 2024 in Golang",
	Long: `My solutions to Advent of Code for 2024
This year I have decided to implement the solutions in Go.

To run a solution for a given day, Use go run main.go <day99> <part9>`,
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		var file io.Reader

		if Test {
			file = strings.NewReader(testData)
		} else {
			ifile, err := os.Open(inputFile)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer ifile.Close()
			file = bufio.NewReader(ifile)
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
	rootCmd.PersistentFlags().BoolVarP(&Test, "test", "t", false, "Use Test Input")
	rootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "Print Debug Messages to Console")
}
