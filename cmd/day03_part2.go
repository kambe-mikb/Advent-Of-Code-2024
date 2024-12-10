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
    "strings"
	
	"github.com/spf13/cobra"
)

var mulRe *(regexp.Regexp)

func init() {
	day03Cmd.AddCommand(day03part2Cmd)
    mulRe = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
}

func evalMultiplies(expr string) (total int) {
    mulExprs := mulRe.FindAllStringSubmatch(expr, -1)
    for _, operation := range mulExprs {
        leftOperand, _ := strconv.Atoi(operation[1])
        rightOperand, _ := strconv.Atoi(operation[2])
        total += leftOperand * rightOperand
    }
    return total
}

// day03part2Cmd represents the part2 command
var day03part2Cmd = &cobra.Command{
	Use:   "part2",
	Short: "Part 2 of Advent of Code Day 03",
	Long: `What do you get if you add up all of the results of the multiplications?`,
	Run: func(cmd *cobra.Command, args []string) {
		testData = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`
		solution = func(input *bufio.Scanner) (total int) {
			// I hate using the "slurp the whole file into memory" tactic,
			// but sometimes it can't be helped. Nothing else seems to
			// work correctly
			var buffer strings.Builder 

			disableRe := regexp.MustCompile(`don't\(\)`)
			enableRe := regexp.MustCompile(`do\(\)`)
 			for input.Scan() {
 				buffer.WriteString(strings.TrimRight(input.Text(), "\r\n"))
 			}
			enabledExprs := disableRe.Split(buffer.String(), -1)
			for index, expr := range enabledExprs {
				if index == 0 {
					total += evalMultiplies(expr)
				} else {
					disabledExpr := enableRe.Split(expr, 2)
					if (len(disabledExpr) > 1) {
						total += evalMultiplies(disabledExpr[1])
					}
				}
			}
 			return total
 		}
	},
	PostRun: day03Cmd.Run,
}