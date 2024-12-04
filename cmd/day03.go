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
	"github.com/spf13/cobra"
)

// day03Cmd represents the day03 command
var day03Cmd = &cobra.Command{
	Use:   "day03",
	Short: "Day 3: Mull It Over",
	Long: `    "Our computers are having issues, so I have no idea if we have
any Chief Historians in stock! You're welcome to check the warehouse,
though," says the mildly flustered shopkeeper at the North Pole Toboggan
Rental Shop. The Historians head out to take a look.

The shopkeeper turns to you. "Any chance you can see why our computers are
having issues again?"

The computer appears to be trying to run a program, but its memory (your
puzzle input) is corrupted. All of the instructions have been jumbled up!

It seems like the goal of the program is just to multiply some numbers. It
does that with instructions like mul(X,Y), where X and Y are each 1-3 digit
numbers. For instance, mul(44,46) multiplies 44 by 46 to get a result of
2024. Similarly, mul(123,4) would multiply 123 by 4.

However, because the program's memory has been corrupted, there are also
many invalid characters that should be ignored, even if they look like part
of a mul instruction. Sequences like mul(4*, mul(6,9!, ?(12,34), or mul ( 
  2 , 4 ) do nothing.

For example, consider the following section of corrupted memory:

xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))

Only the four highlighted sections are real mul instructions. Adding up the
result of each instruction produces 161 (2*4 + 5*5 + 11*8 + 8*5).

Scan the corrupted memory for uncorrupted mul instructions.`,
	Run: func(cmd *cobra.Command, args []string) {
		inputFile = "input03.txt"
	},
	PersistentPostRun: rootCmd.Run,
}

func init() {
	rootCmd.AddCommand(day03Cmd)
}
