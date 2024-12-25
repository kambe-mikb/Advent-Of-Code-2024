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

func antiNodeContains(an []int) func([]int) bool {
	return func(el []int) bool {
		if (an[0] == el[0]) && (an[1] == el[1]) {
			return true
		} else {
			debugln("Found no repeat of", an, "at", el)
			return false
		}
	}
}

func day08part1(input *bufio.Scanner) (total int) {
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
	rowmax := len(startingMap) - 1
	colmax := len(startingMap[0]) - 1

	antiNodes := make([][]int, 0)
	for _, nodes := range antennae {
		for index1, node1 := range nodes[:(len(nodes) - 1)] {
			for _, node2 := range nodes[index1+1:] {
				cdistance := node2[1] - node1[1]
				rdistance := node2[0] - node1[0]
				rAntinode1 := node2[0] + rdistance
				if !((rAntinode1 < 0) || (rAntinode1 > rowmax)) {
					cAntinode1 := node2[1] + cdistance
					if !((cAntinode1 < 0) || (cAntinode1 > colmax)) {
						if antiNode := []int{rAntinode1, cAntinode1}; !slices.ContainsFunc(
							antiNodes,
							antiNodeContains(antiNode),
						) {
							total++
							antiNodes = append(antiNodes, antiNode)
						}
					}
				}
				rAntinode2 := node1[0] - rdistance
				if (rAntinode2 < 0) || (rAntinode2 > rowmax) {
					continue
				}
				cAntinode2 := node1[1] - cdistance
				if (cAntinode2 < 0) || (cAntinode2 > colmax) {
					continue
				}
				antiNode := []int{rAntinode2, cAntinode2}
				if slices.ContainsFunc(antiNodes, antiNodeContains(antiNode)) {
					continue
				}
				total++
				antiNodes = append(antiNodes, antiNode)
			}
		}
	}

	return total
}
