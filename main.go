package main

import (
	day01 "AOC24/programs/day01"
	day02 "AOC24/programs/day02"
	day03 "AOC24/programs/day03"
	day04 "AOC24/programs/day04"
	day05 "AOC24/programs/day05"
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

/*


 */

func main() {
	fmt.Println("Starting The AdventOfCode2024 Script\n")

	// CHOOSE SETTINGS
	m := 1 // 0 = Test Input | 1 = Actual Input

	//Days Complete
	day1aans, day1atime := day01.Day1a(m) //Complete
	day1bans, day1btime := day01.Day1b(m) //Complete
	day2aans, day2atime := day02.Day2a(m) //Complete
	day2bans, day2btime := day02.Day2b(m) //Complete
	day3aans, day3atime := day03.Day3a(m) //Complete
	day3bans, day3btime := day03.Day3b(m) //Complete
	day4aans, day4atime := day04.Day4a(m) //Complete
	day4bans, day4btime := day04.Day4b(m) //Complete
	day5aans, day5atime := day05.Day5a(m) //Complete
	day5bans, day5btime := day05.Day5b(m) //Complete

	data := [][]string{
		[]string{"Day01a", strconv.Itoa(day1aans), fmt.Sprintf("%s", day1atime)},
		[]string{"Day01b", strconv.Itoa(day1bans), fmt.Sprintf("%s", day1btime)},
		[]string{"Day02a", strconv.Itoa(day2aans), fmt.Sprintf("%s", day2atime)},
		[]string{"Day02b", strconv.Itoa(day2bans), fmt.Sprintf("%s", day2btime)},
		[]string{"Day03a", strconv.Itoa(day3aans), fmt.Sprintf("%s", day3atime)},
		[]string{"Day03b", strconv.Itoa(day3bans), fmt.Sprintf("%s", day3btime)},
		[]string{"Day04a", strconv.Itoa(day4aans), fmt.Sprintf("%s", day4atime)},
		[]string{"Day04b", strconv.Itoa(day4bans), fmt.Sprintf("%s", day4btime)},
		[]string{"Day05a", strconv.Itoa(day5aans), fmt.Sprintf("%s", day5atime)},
		[]string{"Day05b", strconv.Itoa(day5bans), fmt.Sprintf("%s", day5btime)},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Day", "Answer", "Time"})
	for _, v := range data {
		table.Append(v)
	}
	table.SetColumnAlignment([]int{0, 1, 2})
	table.Render() // Send output

}
