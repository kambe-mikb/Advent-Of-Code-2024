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
	
	"github.com/spf13/cobra"
)

func init() {
	day04Cmd.AddCommand(day04part2Cmd)
}

// day04part2Cmd represents the part2 command
var day04part2Cmd = &cobra.Command{
	Use:   "part2",
	Short: "Part 2 of Advent of Code Day 04",
	Long: `What do you get if you add up all of the results of the multiplications?`,
	Run: func(cmd *cobra.Command, args []string) {
		testData = ``
		solution = func(input *bufio.Scanner) (total int) {
 			for input.Scan() {
 			}

 			return total
 		}
	},
	PostRun: day04Cmd.Run,
}