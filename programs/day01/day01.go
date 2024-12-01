// /programs/day01.go

package day01

import (
	util "AOC24/programs"
	"math"
	"sort"
	"time"
)

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ DAY 01A ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func Day1a(m int) (int, time.Duration) {
	start := time.Now()
	var lines []string
	if m == 0 {
		lines = util.Returnlines("inputdata/day_01/day_01_test.txt")
	} else {
		lines = util.Returnlines("inputdata/day_01/day_01_actual.txt")
	}

	//Setup VARS For Today
	var firstnumbers []int
	var secondnumbers []int
	var score float64

	//Days Program
	for i := 0; i < len(lines); i++ {
		splitline := util.Splitlines(lines[i], " ")                            //Split Line Into Parts
		firstnumbers = append(firstnumbers, util.StringtoNumber(splitline[0])) //Convert Strings Into INT
		secondnumbers = append(secondnumbers, util.StringtoNumber(splitline[1]))
	}

	sort.Ints(firstnumbers) //Use GO to Organize
	sort.Ints(secondnumbers)

	for i := 0; i < len(firstnumbers); i++ {
		score += math.Abs(float64(firstnumbers[i]) - float64(secondnumbers[i])) //Do Scoring
	}

	return int(score), time.Since(start)
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ DAY 01B ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func Day1b(m int) (int, time.Duration) {
	start := time.Now()
	var lines []string
	if m == 0 {
		lines = util.Returnlines("inputdata/day_01/day_01_test.txt")
	} else {
		lines = util.Returnlines("inputdata/day_01/day_01_actual.txt")
	}

	//Setup VARS For Today
	var firstnumbers []int
	var secondnumbers []int
	var score int

	//Days Program
	for i := 0; i < len(lines); i++ {
		splitline := util.Splitlines(lines[i], " ")
		firstnumbers = append(firstnumbers, util.StringtoNumber(splitline[0]))
		secondnumbers = append(secondnumbers, util.StringtoNumber(splitline[1]))
	}

	sort.Ints(firstnumbers)
	sort.Ints(secondnumbers)

	for i := 0; i < len(firstnumbers); i++ {
		occurances := 0
		for j := 0; j < len(secondnumbers); j++ {
			if firstnumbers[i] == secondnumbers[j] {
				occurances++
			}
		}
		score += firstnumbers[i] * occurances
	}

	return score, time.Since(start)
}
