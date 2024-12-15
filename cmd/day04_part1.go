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
	//    "fmt"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	day04Cmd.AddCommand(day04part1Cmd)
}

// day04part1Cmd represents the part1 command
var day04part1Cmd = &cobra.Command{
	Use:   "part1",
	Short: "Part 1 of Advent of Code Day 04",
	Long:  `How many times does XMAS appear?`,
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
MXMXAXMASX`
		solution = func(input *bufio.Scanner) (total int) {
			var rows []string

			for input.Scan() {
				rows = append(rows, strings.TrimRight(input.Text(), "\r\n"))
			}
			rowMax := len(rows) - 1

			for rIndex, row := range rows {
				columnMax := len(row) - 1

				for cIndex, letter := range row {
					var rSearch []int
					var cSearch []int

					if string(letter) != "X" {
						continue
					}

					switch rIndex {
					case 0:
						rSearch = []int{0, 1}
					case rowMax:
						rSearch = []int{rIndex - 1, rIndex}
					default:
						rSearch = []int{rIndex - 1, rIndex, rIndex + 1}
					}

					switch cIndex {
					case 0:
						cSearch = []int{0, 1}
					case rowMax:
						cSearch = []int{cIndex - 1, cIndex}
					default:
						cSearch = []int{cIndex - 1, cIndex, cIndex + 1}
					}

					for _, r := range rSearch {
						for _, c := range cSearch {
							if (string(rows[r][c]) != "M") || ((rIndex > 0) && (r == 0)) || ((rIndex < rowMax) && (r == rowMax)) || ((cIndex > 0) && (c == 0)) || ((cIndex < columnMax) && (c == columnMax)) {
								continue
							}
							newr := r + (r - rIndex)
							newc := c + (c - cIndex)
							if (string(rows[newr][newc]) != "A") || ((r > 0) && (newr == 0)) || ((r < rowMax) && (newr == rowMax)) || ((c > 0) && (newc == 0)) || ((c < columnMax) && (newc == columnMax)) {
								continue
							}
							if string(rows[newr+(r-rIndex)][newc+(c-cIndex)]) == "S" {
								total++
							}
						}
					}
				}
			}

			return total
		}
	},
	PostRun: day04Cmd.Run,
}
