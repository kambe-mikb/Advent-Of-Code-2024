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
    "regexp"
    "strconv"
	
	"github.com/spf13/cobra"
)

func init() {
	day03Cmd.AddCommand(day03part1Cmd)
}

// day03part1Cmd represents the part1 command
var day03part1Cmd = &cobra.Command{
	Use:   "part1",
	Short: "Part 1 of Advent of Code Day 03",
	Long: `What do you get if you add up all of the results of the multiplications?`,
	Run: func(cmd *cobra.Command, args []string) {
		testData = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))
`
		solution = func(input *bufio.Scanner) (total int) {
            mulRe := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
 			for input.Scan() {
                mulExprs := mulRe.FindAllStringSubmatch(input.Text(), -1)
                for _, operation := range mulExprs {
                    leftOperand, _ := strconv.Atoi(operation[1])
                    rightOperand, _ := strconv.Atoi(operation[2])
                    total += leftOperand * rightOperand
                }
 			}

 			return total
 		}
	},
	PostRun: day03Cmd.Run,
}