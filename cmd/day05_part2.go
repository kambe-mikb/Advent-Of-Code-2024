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
	day05Cmd.AddCommand(day05part2Cmd)
}

func validatePageList(rules map[string]([]string), pages []string) bool {
	for pageIndex, page := range pages {
		if pageRule, exists := rules[page]; exists {
			for _, p := range pageRule {
				switch {
				case slices.Index(pages, p) == -1:
					continue
				case slices.Index(pages, p) < pageIndex:
					return false
				}
			}
		}
	}
	return true
}

func rearrangePages(rules map[string]([]string), pages []string) []string {
	for pageIndex, page := range pages {
		if pageRule, exists := rules[page]; exists {
			for _, p := range pageRule {
				index := slices.Index(pages, p)
				switch {
				case index == -1:
					continue
				case index < pageIndex:
					return append(slices.Delete(pages, index, index+1), p)
				}
			}
		}
	}
	return pages
}

// day05part1Cmd represents the part2 command
var day05part2Cmd = &cobra.Command{
	Use:   "part2",
	Short: "Part 2 of Advent of Code Day 05",
	Long:  `What do you get if you add up the middle page numbers after correctly ordering just those updates?`,
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
					if !validatePageList(rules, pages) {
						for !validatePageList(rules, pages) {
							pages = rearrangePages(rules, pages)
						}
						mid := len(pages) / 2
						mPage, _ := strconv.Atoi(pages[mid])
						total += mPage
					}
				}
			}
			return total
		}
	},
	PostRun: day05Cmd.Run,
}
