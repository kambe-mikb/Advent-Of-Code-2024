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
	"strconv"
)

func day09part2(input *bufio.Scanner) (total int) {
	var isFree bool
	var blocks []int = make([]int, 0)
	var inode []int
	var cursor int
	var directory [][]int = make([][]int, 0)
	var freelist [][]int = make([][]int, 0)

	input.Split(bufio.ScanRunes)
	for ; input.Scan(); cursor = len(blocks) {
		digit, _ := strconv.Atoi(string(input.Text()))
		if isFree {
			inode = []int{-1}
			isFree = false
			if digit == 0 {
				continue
			} else {
				freelist = append(freelist, []int{cursor, digit})
			}
		} else {
			inode = []int{len(directory)}
			directory = append(directory, []int{cursor, digit})
			isFree = true
		}
		blocks = slices.Concat(blocks, slices.Repeat(inode, digit))
	}
	debugState(blocks)
	debugln(directory)
	debugln(freelist)
	for di, dirCursor := range slices.Backward(directory) {
		debugln("Checking file", di, "at", dirCursor[0], "length", dirCursor[1])
	search:
		for fi, freeCursor := range freelist {
			switch {
			case freeCursor[0] > dirCursor[0]:
				break search
			case freeCursor[1] > dirCursor[1]:
				debugln("    Move into", freeCursor[0], "length", freeCursor[1])
				directory[di][0] = freeCursor[0]
				freelist[fi][0] += dirCursor[1]
				freelist[fi][1] -= dirCursor[1]
				break search
			case freeCursor[1] == dirCursor[1]:
				debugln("    Replace", freeCursor[0], "length", freeCursor[1])
				directory[di][0] = freeCursor[0]
				freelist = slices.Delete(freelist, fi, fi+1)
				break search
			}
		}
	}
	debugln(directory)
	debugln(freelist)

	blocks = slices.Repeat([]int{-1}, len(blocks))
	for di, dirCursor := range directory {
		inode = []int{di}
		if blocks[dirCursor[0]] == di {
			continue
		}
		copy(
			blocks[dirCursor[0]:(dirCursor[0]+dirCursor[1])],
			slices.Repeat(inode, dirCursor[1]),
		)
		debugln("Moving file", di, "to", dirCursor[0])
	}
	debugState(blocks)
	for i, v := range blocks {
		if v == -1 {
			continue
		}
		total += i * v
	}
	return total
}
