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
	"sort"
	"strconv"
	"strings"
	
	"github.com/spf13/cobra"
)

// part1Cmd represents the part1 command
var day01part1Cmd = &cobra.Command{
	Use:   "part1",
	Short: "Part 1 of Advent of Code Day 01",
	Long: `What is the total distance between your lists?.`,
	Run: func(cmd *cobra.Command, args []string) {
		testData = `3   4
 4   3
 2   5
 1   3
 3   9
 3   3`
 		solution = func(input *bufio.Scanner) (total int) {
 			var list1, list2 []int
 			
 			for input.Scan() {
 				fields := strings.Fields(input.Text())
 				value, _ := strconv.Atoi(fields[0])
 				list1 = append(list1, value)
 				value, _ = strconv.Atoi(fields[1])
 				list2 = append(list2, value)
 			}
 			sort.Ints(list1)
 			sort.Ints(list2)

 			for index, value1 := range list1 {
 				if distance := value1 - list2[index]; distance < 0 {
 					total -= distance
 				} else {
 					total += distance
 				}  
 			}
 			return total
 		}
	},
	PostRun: day01Cmd.Run,
}

func init() {
	day01Cmd.AddCommand(day01part1Cmd)
}
