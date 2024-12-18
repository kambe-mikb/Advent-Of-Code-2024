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
	//	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	day06Cmd.AddCommand(day06part2Cmd)
}
func makeNewMap(curMap [][]bool, curGuard [3]int) (newMap [][]bool, newSeen [][]int, newGuard [3]int) {
	newMap = make([][]bool, len(curMap))
	for ri, rv := range curMap {
		newMap[ri] = slices.Clone(rv)
		newSeen = append(newSeen, make([]int, len(rv)))
	}
	newGuard = curGuard
    return
}

// day06part2Cmd represents the part2 command
var day06part2Cmd = &cobra.Command{
	Use:   "part2",
	Short: "Part 2 of Advent of Code Day 06",
	Long:  `How many different positions could you choose for this obstruction?`,
	Run: func(cmd *cobra.Command, args []string) {
		testData = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`
		solution = func(input *bufio.Scanner) (total int) {
			var startingMap [][]bool
			var startingGuard [3]int
			var row int

			for ; input.Scan(); row++ {
				line := strings.TrimRight(input.Text(), "\r\n")
				startingMap = append(startingMap, make([]bool, len(line)))
				for column, letter := range line {
					switch string(letter) {
					case "#":
						startingMap[row][column] = true
					case "^":
						startingMap[row][column] = false
						startingGuard = [3]int{row, column, 0}
					case ">":
						startingMap[row][column] = false
						startingGuard = [3]int{row, column, 1}
					case "v":
						startingMap[row][column] = false
						startingGuard = [3]int{row, column, 2}
					case "<":
						startingMap[row][column] = false
						startingGuard = [3]int{row, column, 3}
					default:
						startingMap[row][column] = false
					}
				}
			}
            curMap, seen, guard := makeNewMap(startingMap, startingGuard)
            for ri, rv := range curMap {
                for ci, cv := range rv {
                    if cv { // there's already an obstacle'
                        continue
                    }
                    curMap[ri][ci] = true // add an obstacle
                    seen[guard[0]][guard[1]] = guard[2] + 1
                    for off, guard, _ := step(curMap, guard); !off; off, guard, _ = step(curMap, guard) {
                        if seen[guard[0]][guard[1]] == 0 {
                            seen[guard[0]][guard[1]] = guard[2] + 1
                        } else {
                            if seen[guard[0]][guard[1]] == guard[2] + 1 {
                                fmt.Println("! Loop found with new obstacle at", ri, ",", ci)
                                // we've already been to this loc on this vector'
                                total++
                                break
                            }
                        }
                    }
                    curMap, seen, guard = makeNewMap(startingMap, startingGuard)                    
                }
            } 
			return total
		}
	},
	PostRun: day06Cmd.Run,
}
