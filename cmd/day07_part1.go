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

	//	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	day07Cmd.AddCommand(day07part1Cmd)
}

// day07part1Cmd represents the part1 command
var day07part1Cmd = &cobra.Command{
	Use:   "part1",
	Short: "Part 1 of Advent of Code Day 07",
	Run: func(cmd *cobra.Command, args []string) {
		testData = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
`
		solution = func(input *bufio.Scanner) (total int) {
			for input.Scan() {
				var operands []int

				lhs, rhs, notEmpty := strings.Cut(strings.TrimRight(input.Text(), "\r\n"), ": ")
				if !notEmpty {
					continue
				}
				calibration, _ := strconv.Atoi(lhs)

				for _, field := range strings.Fields(rhs) {
					operand, _ := strconv.Atoi(field)
					operands = append(operands, operand)
				}
				opSet := make([]bool, len(operands)-1)
				operations := [][]bool{opSet}
				for oi := 0; oi < len(opSet); oi++ {
					for _, ops := range operations {
						newOps := slices.Clone(ops)
						newOps[oi] = true
						operations = append(operations, newOps)
					}
				}

				for _, ops := range operations {
					equation := operands[0]

					for oi, multiply := range ops {
						if multiply {
							equation *= operands[oi+1]
						} else {
							equation += operands[oi+1]
						}
					}
					if calibration == equation {
						total += calibration
						break
					}
				}

			}
			return total
		}
	},
	PostRun: day07Cmd.Run,
}
