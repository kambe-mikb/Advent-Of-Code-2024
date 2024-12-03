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
		fmt.Println("Entering day01part1Cmd.Run")
		testData = `3   4
 4   3
 2   5
 1   3
 3   9
 3   3`
 		solution = func(input *bufio.Scanner) int {
 			var list1, list2 []int
 			
			fmt.Println("Entering Day 01 Part 1 Solution")
 			for input.Scan() {
 				fmt.Println("Current Line: ", input.Text())
 				fields := strings.Fields(input.Text())
 				value, _ := strconv.Atoi(fields[0])
 				fmt.Printf("Append to List1: %v", value)
 				list1 = append(list1, value)
 				value, _ = strconv.Atoi(fields[1])
 				fmt.Printf(" Append to List2: %v\n", value)
 				list2 = append(list2, value)
 			}
 			sort.Ints(list1)
 			sort.Ints(list2)
 			total := 0
 			for index, value1 := range list1 {
 				if distance := value1 - list2[index]; distance < 0 {
 					total -= distance
 				} else {
 					total += distance
 				}  
 			}
			fmt.Println("Exiting Day 01 Part 1 Solution")
 			return total
 		}
		fmt.Println("Exiting day01part1Cmd.Run")
	},
	PostRun: day01Cmd.Run,
}

func init() {
	day01Cmd.AddCommand(day01part1Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// part1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// part1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
