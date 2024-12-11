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

// day04Cmd represents the day04 command
var day04Cmd = &cobra.Command{
	Use:   "day04",
	Short: "Day 4: Ceres Search",
	Long: `    "Looks like the Chief's not here. Next!" One of The Historians
    pulls out a device and pushes the only button on it. After a brief flash,
    you recognize the interior of the Ceres monitoring station!
    As the search for the Chief continues, a small Elf who lives on the station
    tugs on your shirt; she'd like to know if you could help her with her word
    search (your puzzle input). She only has to find one word: XMAS.

    This word search allows words to be horizontal, vertical, diagonal, written
    backwards, or even overlapping other words. It's a little unusual, though,
    as you don't merely need to find one instance of XMAS - you need to find
    all of them. Here are a few ways XMAS might appear, where irrelevant
    characters have been replaced with .:

    ..X...
    .SAMX.
    .A..A.
    XMAS.S
    .X....

    The actual word search will be full of letters instead. For example:

    MMMSXXMASM
    MSAMXMSMSA
    AMXSXMAAMM
    MSAMASMSMX
    XMASAMXAMM
    XXAMMXXAMA
    SMSMSASXSS
    SAXAMASAAA
    MAMMMXMMMM
    MXMXAXMASX

    In this word search, XMAS occurs a total of 18 times; here's the same word
    search again, but where letters not involved in any XMAS have been replaced
    with .:

    ....XXMAS.
    .SAMXMS...
    ...S..A...
    ..A.A.MS.X
    XMASAMX.MM
    X.....XA.A
    S.S.S.S.SS
    .A.A.A.A.A
    ..M.M.M.MM
    .X.X.XMASX

    Take a look at the little Elf's word search. 
    The Elf looks quizzically at you. Did you misunderstand the assignment?

    Looking for the instructions, you flip over the word search to find that
    this isn't actually an XMAS puzzle; it's an X-MAS puzzle in which you're
    supposed to find two MAS in the shape of an X. One way to achieve that is
    like this:

    M.S
    .A.
    M.S

    Irrelevant characters have again been replaced with . in the above diagram.
    Within the X, each MAS can be written forwards or backwards.

    Here's the same example from before, but this time all of the X-MASes have
    been kept instead:

    .M.S......
    ..A..MSMS.
    .M.S.MAA..
    ..A.ASMSM.
    .M.S.M....
    ..........
    S.S.S.S.S.
    .A.A.A.A..
    M.M.M.M.M.
    ..........

    In this example, an X-MAS appears 9 times.

    Flip the word search from the instructions back over to the word search
    side and try again.`,
	Run: func(cmd *cobra.Command, args []string) {
		inputFile = "input04.txt"
	},
	PersistentPostRun: rootCmd.Run,
}

func init() {
	rootCmd.AddCommand(day04Cmd)
}
