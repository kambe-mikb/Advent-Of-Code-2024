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

func debugState(blocks []int) {
	var printable []string = make([]string, len(blocks))

	if Debug {
		for i, sblock := range blocks {
			if sblock == -1 {
				printable[i] = "."
			} else {
				printable[i] = strconv.Itoa(sblock)
			}
		}
	}
	debugln(printable)
}

func day09part1(input *bufio.Scanner) (total int) {
	var isFree bool
	var blocks []int = make([]int, 0)
	var fileid int
	var inode []int

	input.Split(bufio.ScanRunes)
	for input.Scan() {
		digit, _ := strconv.Atoi(string(input.Text()))
		if isFree {
			inode = []int{-1}
			isFree = false
		} else {
			inode = []int{fileid}
			fileid++
			isFree = true
		}
		blocks = slices.Concat(blocks, slices.Repeat(inode, digit))
	}
	debugState(blocks)
	for b := slices.Index(blocks, -1); b > -1; b = slices.Index(blocks, -1) {
		down := slices.Backward(blocks)
		var r int
		var v int

		for r, v = range down {
			if v != -1 {
				break
			}
		}
		if r < b {
			break
		}
		blocks[b] = blocks[r]
		blocks[r] = -1
		debugState(blocks)
	}
	for i, v := range blocks {
		if v == -1 {
			break
		}
		total += i * v
	}
	return total
}
