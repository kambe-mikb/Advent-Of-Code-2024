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
	"slices"
)

func day08part2(input *bufio.Scanner) (total int) {
	var startingMap [][]string
	var antennae map[string]([]([]int)) = make(map[string][][]int)
	var row int

	for ; input.Scan(); row++ {
		line := input.Text()
		startingMap = append(startingMap, make([]string, len(line)))
		for column, theRune := range line {
			letter := string(theRune)
			if letter == "." {
				continue
			}
			startingMap[row][column] = letter
			if _, exists := antennae[letter]; !exists {
				antennae[letter] = make([][]int, 0)
			}
			antennae[letter] = append(antennae[letter], []int{row, column})
		}
	}

	antiNodes := make([][]int, 0)
	for _, nodes := range antennae {
		for index1, node1 := range nodes[:(len(nodes) - 1)] {
			for _, node2 := range nodes[index1+1:] {
				A := node2[1] - node1[1]
				B := node2[0] - node1[0]
				C := B*node1[1] - A*node1[0]
				for rindex, row := range startingMap {
					for cindex := range row {
						if ((A * rindex) - (B * cindex) + C) == 0 {
							// we are on the line
							if antiNode := []int{rindex, cindex}; !slices.ContainsFunc(
								antiNodes,
								antiNodeContains(antiNode),
							) {
								total++
								antiNodes = append(antiNodes, antiNode)
							}
						}
					}
				}
			}
		}
	}

	return total
}
