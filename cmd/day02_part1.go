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
	"strconv"
	"strings"
	
	"github.com/spf13/cobra"
)

func init() {
	day02Cmd.AddCommand(day02part1Cmd)
}

// day02part1Cmd represents the part1 command
var day02part1Cmd = &cobra.Command{
	Use:   "part1",
	Short: "Part 1 of Advent of Code Day 02",
	Long: `What is their similarity score?`,
	Run: func(cmd *cobra.Command, args []string) {
		testData = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9

`
		solution = func(input *bufio.Scanner) (total int) {
 			for input.Scan() {
 				var safe bool
 				var isIncreasing bool

 				fields := strings.Fields(input.Text())
                if len(fields) == 0 {
                    break
                }
 				reports := make([]int, len(fields))
                line:
     				for index, value := range fields {
    	 				reports[index], _ = strconv.Atoi(value)
    	 				switch index {
    	 				case 0:
    	 					safe = true
    	 				case 1:
    		 				comparison := reports[index] - reports[index - 1]
                            safe = ((comparison != 0) && (comparison >= -3) && (comparison <= 3))
    	 					if (!safe) {
    	 						break line
    	 					}
    	 					isIncreasing = (comparison > 0)
    	 				default:
    		 				comparison := reports[index] - reports[index - 1]
                            safe = (comparison != 0)
    	 					if (!safe) {
    	 						break line
    	 					}
    	 					// Using a switch, rather than if-then-else for readability
     						switch {
     						case (isIncreasing && (comparison > 0) && (comparison <= 3)):
     							safe = true
     						case ((!isIncreasing) && (comparison < 0) && (comparison >= -3)):
     							safe = true
     						default:
     							safe = false
     							break line
     						}
     						isIncreasing = (comparison > 0)
    	 				}
     				}
 				if safe {
 					total++
 				}
 			}

 			return total
 		}
	},
	PostRun: day02Cmd.Run,
}