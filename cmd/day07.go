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

// day07Cmd represents the day06 command
var day07Cmd = &cobra.Command{
	Use:   "day07",
	Short: "Day 7: ",
	Long: `The Historians take you to a familiar rope bridge over a river in
the middle of a jungle. The Chief isn't on this side of the bridge, though; maybe he's on the other side?

When you go to cross the bridge, you notice a group of engineers trying to repair it. (Apparently, it breaks
pretty frequently.) You won't be able to cross until it's fixed.

You ask how long it'll take; the engineers tell you that it only needs final calibrations, but some young
elephants were playing nearby and stole all the operators from their calibration equations! They could finish
the calibrations if only someone could determine which test values could possibly be produced by placing any
combination of operators into their calibration equations (your puzzle input).

For example:

190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20

Each line represents a single equation. The test value appears before the colon on each line; it is your job
to determine whether the remaining numbers can be combined with operators to produce the test value.

Operators are always evaluated left-to-right, not according to precedence rules. Furthermore, numbers in the
equations cannot be rearranged. Glancing into the jungle, you can see elephants holding two different types of
operators: add (+) and multiply (*).

Only three of the above equations can be made true by inserting operators:

    190: 10 19 has only one position that accepts an operator: between 10 and 19. Choosing + would give 29, but
	choosing * would give the test value (10 * 19 = 190).
    3267: 81 40 27 has two positions for operators. Of the four possible configurations of the operators, two
	cause the right side to match the test value: 81 + 40 * 27 and 81 * 40 + 27 both equal 3267 (when evaluated
	left-to-right)!
    292: 11 6 16 20 can be solved in exactly one way: 11 + 6 * 16 + 20.

The engineers just need the total calibration result, which is the sum of the test values from just the equations
that could possibly be true. In the above example, the sum of the test values for the three equations listed above

Determine which equations could possibly be true.

--- Part Two ---

The engineers seem concerned; the total calibration result you gave them is nowhere close to being within safety
tolerances. Just then, you spot your mistake: some well-hidden elephants are holding a third type of operator.

The concatenation operator (||) combines the digits from its left and right inputs into a single number. For example,
12 || 345 would become 12345. All operators are still evaluated left-to-right.

Now, apart from the three equations that could be made true using only addition and multiplication, the above example
has three more equations that can be made true by inserting operators:

    156: 15 6 can be made true through a single concatenation: 15 || 6 = 156.
    7290: 6 8 6 15 can be made true using 6 * 8 || 6 * 15.
    192: 17 8 14 can be made true using 17 || 8 + 14.

Adding up all six test values (the three that could be made before using only + and * plus the new three that can now
be made by also using ||) produces the new total calibration result of 11387.

Using your new knowledge of elephant hiding spots, determine which equations could possibly be true.`,
	Run: func(cmd *cobra.Command, args []string) {
		inputFile = "input07.txt"
	},
	PersistentPostRun: rootCmd.Run,
}

func init() {
	rootCmd.AddCommand(day07Cmd)
}
