// /programs/day03.go

package day03

import (
	util "AOC24/programs"
	"regexp"
	"time"
)

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ DAY 03A ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func Day3a(m int) (int, time.Duration) {
	start := time.Now()
	lines := util.Returnlines("inputdata/day_03/day_03_actual.txt", "inputdata/day_03/day_03_test.txt", m)

	//Setup VARS For Today
	var matchingexpressions []string
	score := 0

	//Days Program
	for i := 0; i < len(lines); i++ { //For Each Line In Problem Statement
		re := regexp.MustCompile(`mul\(\d+,\d+\)`)                                           //Define Regex Expression (https://regex101.com/)
		matchingexpressions = append(matchingexpressions, re.FindAllString(lines[i], -1)...) //For Each Matching Expression Add To Matching Expressions (use ... to unpack)
	}
	for i := 0; i < len(matchingexpressions); i++ {
		re := regexp.MustCompile(`\d+`)                                              //Define Regex (Groups Of Numbers)
		numbers := re.FindAllString(matchingexpressions[i], -1)                      //Find ALL Occurances Of This
		score += (util.StringtoNumber(numbers[0]) * util.StringtoNumber(numbers[1])) //Multply Occurances | Add To Score
	}
	return score, time.Since(start)
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ DAY 03B ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func Day3b(m int) (int, time.Duration) {
	start := time.Now()
	lines := util.Returnlines("inputdata/day_03/day_03_actual.txt", "inputdata/day_03/day_03_test.txt", m)

	//Setup VARS For Today
	score := 0
	dodontstatus := true
	var matchingexpressions []string

	//Days Program
	for i := 0; i < len(lines); i++ { //Same As Day 1, Except We Expand to Also Search For do() and don't()
		re := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
		matchingexpressions = append(matchingexpressions, re.FindAllString(lines[i], -1)...)
	}
	for i := 0; i < len(matchingexpressions); i++ {
		if matchingexpressions[i] == "do()" { //If do() Then Switch Math To True
			dodontstatus = true
		} else if matchingexpressions[i] == "don't()" { //If don't() Then Switch Math to False
			dodontstatus = false
		} else { //If It's a MUL() Then Add If Math Status Is True
			re := regexp.MustCompile(`\d+`)
			numbers := re.FindAllString(matchingexpressions[i], -1)
			if dodontstatus {
				score += (util.StringtoNumber(numbers[0]) * util.StringtoNumber(numbers[1]))
			}
		}
	}

	return score, time.Since(start)
}
