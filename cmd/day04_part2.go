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
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	day04Cmd.AddCommand(day04part2Cmd)
}

// day04part1Cmd represents the part1 command
var day04part2Cmd = &cobra.Command{
	Use:   "part2",
	Short: "Part 2 of Advent of Code Day 04",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		testData = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMAS`
		solution = func(input *bufio.Scanner) (total int) {
			var rows []string

			for input.Scan() {
				rows = append(rows, strings.TrimRight(input.Text(), "\r\n"))
			}
			rowMax := len(rows) - 1

			for rIndex, row := range rows {
				columnMax := len(row) - 1
			row:
				for cIndex, letter := range row {
					if string(letter) != "A" {
						continue
					}

					if (rIndex == 0) || (rIndex == rowMax) || (cIndex == 0) || (cIndex == columnMax) {
						continue
					}
					cSearch := [2]int{cIndex - 1, cIndex + 1}

					valid := 0
					for _, c := range cSearch {
						var expectedLetter string

						switch string(rows[rIndex-1][c]) {
						case "M":
							expectedLetter = "S"
						case "S":
							expectedLetter = "M"
						default:
							continue row
						}

						if string(rows[rIndex+1][cIndex+(cIndex-c)]) == expectedLetter {
							valid++
						}
					}
					if valid >= 2 {
						total++
					}
				}
			}

			return total
		}
	},
	PostRun: day04Cmd.Run,
}
