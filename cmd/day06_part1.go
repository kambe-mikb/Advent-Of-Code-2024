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
	//	"slices"
	//	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	day06Cmd.AddCommand(day06part1Cmd)
}

var directions = [4]string{"^", ">", "v", "<"}

func step(obstacle [][]bool, theGuard [3]int) (bool, [3]int, int) {
    var r int
    var c int
    
    switch theGuard[2] {
    case 0:
        if theGuard[0] == 0 {
            return true, theGuard, 0
        }
        r = theGuard[0] - 1
        c = theGuard[1]
    case 1:
        if theGuard[1] == (len(obstacle[theGuard[0]]) - 1) {
            return true, theGuard, 0
        }
        r = theGuard[0]
        c = theGuard[1] + 1
    case 2:
        if theGuard[0] == (len(obstacle) - 1){
            return true, theGuard, 0
        }
        r = theGuard[0] + 1
        c = theGuard[1]
    case 3:
        if theGuard[1] == 0 {
            return true, theGuard, 0
        }
        r = theGuard[0]
        c = theGuard[1] - 1
    }
    if obstacle[r][c] {
        theGuard[2] = (theGuard[2] + 1) % 4
        return false, theGuard, 0 //step(obstacle, theGuard)
    } else {
        theGuard[0] = r
        theGuard[1] = c
        return false, theGuard, 1
    }
}

// day06part1Cmd represents the part1 command
var day06part1Cmd = &cobra.Command{
	Use:   "part1",
	Short: "Part 1 of Advent of Code Day 06",
	Long: `How many distinct positions will the guard visit before leaving the
    mapped area?`,
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
            var guard [3]int
            var seen [][]bool
            var row int
            
			for ;input.Scan(); row++  {
                line := strings.TrimRight(input.Text(), "\r\n")
                startingMap = append(startingMap, make([]bool, len(line)))
                seen = append(seen, make([]bool, len(line)))
                for column, letter := range line {
                    switch string(letter) {
                    case "#":
                        startingMap[row][column] = true
                    case "^":
                        startingMap[row][column] = false
                        guard = [3]int {row, column, 0}
                    case ">":
                        startingMap[row][column] = false
                        guard = [3]int {row, column, 1}
                    case "v":
                        startingMap[row][column] = false
                        guard = [3]int {row, column, 2}
                    case "<":
                        startingMap[row][column] = false
                        guard = [3]int {row, column, 3}
                    default:
                        startingMap[row][column] = false
                    }
                }
			}
            seen[guard[0]][guard[1]] = true
            total = 1
            for off, guard, n := step(startingMap, guard); !off; off, guard, n = step(startingMap, guard) {
                if !seen[guard[0]][guard[1]] {
                    total += n
                    seen[guard[0]][guard[1]] = true
                }
            }
			return total
		}
	},
	PostRun: day06Cmd.Run,
}
