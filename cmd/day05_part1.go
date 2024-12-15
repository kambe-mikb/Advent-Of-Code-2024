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
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	day05Cmd.AddCommand(day05part1Cmd)
}

// day05part1Cmd represents the part1 command
var day05part1Cmd = &cobra.Command{
	Use:   "part1",
	Short: "Part 1 of Advent of Code Day 05",
	Long: `What do you get if you add up the middle page number from those
correctly-ordered updates?`,
	Run: func(cmd *cobra.Command, args []string) {
		testData = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`
		solution = func(input *bufio.Scanner) (total int) {
			mode := "rules"
			rules := make(map[string]([]string))

		lines:
			for input.Scan() {
				line := strings.TrimRight(input.Text(), "\r\n")
				switch {
				case line == "":
					mode = "updates"
				case mode == "rules":
					rule := strings.Split(line, "|")
					rules[rule[0]] = append(rules[rule[0]], rule[1])
				case mode == "updates":
					pages := strings.Split(line, ",")
					for pageIndex, page := range pages {
						if pageRule, exists := rules[page]; exists {
							for _, p := range pageRule {
								switch {
								case slices.Index(pages, p) == -1:
									continue
								case slices.Index(pages, p) < pageIndex:
									fmt.Println(line, " is not valid")
									continue lines
								}
							}
						}
					}
					mid := len(pages) / 2
					fmt.Println(line, " is valid. Mid is ", pages[mid])
					mPage, _ := strconv.Atoi(pages[mid])
					total += mPage
				}
			}
			return total
		}
	},
	PostRun: day05Cmd.Run,
}
