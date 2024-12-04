package main

import (
	day01 "AOC24/programs/day01"
	day02 "AOC24/programs/day02"
	day03 "AOC24/programs/day03"
	"fmt"
)

/*


 */

func main() {
	fmt.Println("Starting The AdventOfCode2024 Script")

	// CHOOSE SETTINGS
	m := 1 // 0 = Test Input | 1 = Actual Input

	day1aans, day1atime := day01.Day1a(m) //Complete
	day1bans, day1btime := day01.Day1b(m) //Complete
	day2aans, day2atime := day02.Day2a(m) //Complete
	day2bans, day2btime := day02.Day2b(m) //Complete
	day3aans, day3atime := day03.Day3a(m) //Complete
	day3bans, day3btime := day03.Day3b(m) //Not Complete

	fmt.Printf("Day 01a | Answer: %-11d | Taking: %s\n", day1aans, day1atime)
	fmt.Printf("Day 01b | Answer: %-11d | Taking: %s\n", day1bans, day1btime)

	fmt.Printf("Day 02a | Answer: %-11d | Taking: %s\n", day2aans, day2atime)
	fmt.Printf("Day 02b | Answer: %-11d | Taking: %s\n", day2bans, day2btime)

	fmt.Printf("Day 03a | Answer: %-11d | Taking: %s\n", day3aans, day3atime)
	fmt.Printf("Day 03b | Answer: %-11d | Taking: %s\n", day3bans, day3btime)

}
