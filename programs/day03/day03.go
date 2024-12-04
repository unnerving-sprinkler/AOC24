// /programs/day03.go

package day03

import (
	util "AOC24/programs"
	"regexp"
	"time"
)

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ DAY 01A ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func Day3a(m int) (int, time.Duration) {
	start := time.Now()
	var lines []string
	if m == 0 {
		lines = util.Returnlines("inputdata/day_03/day_03_test.txt")
	} else {
		lines = util.Returnlines("inputdata/day_03/day_03_actual.txt")
	}

	//Setup VARS For Today
	var matchingexpressions []string
	score := 0

	//Days Program
	for i := 0; i < len(lines); i++ {
		re := regexp.MustCompile(`mul\(\d+,\d+\)`)
		matchingexpressions = append(matchingexpressions, re.FindAllString(lines[i], -1)...)
	}

	for i := 0; i < len(matchingexpressions); i++ {
		re := regexp.MustCompile(`\d+`)
		numbers := re.FindAllString(matchingexpressions[i], -1)
		score += (util.StringtoNumber(numbers[0]) * util.StringtoNumber(numbers[1]))
	}

	return score, time.Since(start)
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ DAY 01B ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func Day3b(m int) (int, time.Duration) {
	start := time.Now()
	//var lines []string
	if m == 0 {
		//lines = util.Returnlines("inputdata/day_03/day_03_test.txt")
	} else {
		//lines = util.Returnlines("inputdata/day_03/day_03_actual.txt")
	}

	//Setup VARS For Today

	//Days Program

	return 0, time.Since(start)
}

type instruction struct {
	first  int
	second int
}
