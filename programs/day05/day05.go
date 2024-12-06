// /programs/day05.go

package day05

import (
	util "AOC24/programs"
	"fmt"
	"time"
)

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ DAY 05A ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func Day5a(m int) (int, time.Duration) {
	start := time.Now()
	lines := util.Returnlines("inputdata/day_05/day_05_actual.txt", "inputdata/day_05/day_05_test.txt", m)

	//Setup VARS For Today
	score := 0
	var validpagesets [][]int

	//Days Program
	ruleset, pageset := parseday5(lines)

	for i := 0; i < len(pageset); i++ { //For Each Pageset
		isvalidpageset := true
		for j := 0; j < len(ruleset); j++ { //Loop Through Each Rule
			firstruleexists := false
			secondruleexists := false

			for k := 0; k < len(pageset[i]); k++ {
				var firstruleindex int
				var secondruleindex int
				if pageset[i][k] == ruleset[j][0] {
					firstruleexists = true
					firstruleindex = k
				}
				if pageset[i][k] == ruleset[j][1] {
					secondruleexists = true
					secondruleindex = k
				}

				if firstruleexists && secondruleexists {
					if firstruleindex > secondruleindex {
						isvalidpageset = false
					}

				}
			}
		}

		if isvalidpageset {
			middleindex := (len(pageset[i]) / 2) + (1 / 2)
			score += pageset[i][middleindex]
		}

	}
	fmt.Println(validpagesets)
	return score, time.Since(start)
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ DAY 05B ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func Day5b(m int) (int, time.Duration) {
	start := time.Now()
	//var lines []string
	if m == 0 {
		//lines = util.Returnlines("inputdata/day_05/day_05_test.txt")
	} else {
		//lines = util.Returnlines("inputdata/day_05/day_05_actual.txt")
	}

	//Setup VARS For Today
	score := 0

	//Days Program

	return score, time.Since(start)
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ UTIL ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

func parseday5(lines []string) ([][]int, [][]int) {
	mode := false
	var rulesets [][]int
	var pagesets [][]int
	//var pagesets []int

	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			mode = true
		}
		if !mode {
			thisline := util.Splitlines(lines[i], "|")
			firstrule := util.StringtoNumber(thisline[0])
			secondrule := util.StringtoNumber(thisline[1])
			thisruleset := []int{firstrule, secondrule}
			rulesets = append(rulesets, thisruleset)
		}

		if mode {
			thisline := util.Splitlines(lines[i], ",")
			var thispageset []int
			for i := 0; i < len(thisline); i++ {
				thispageset = append(thispageset, util.StringtoNumber(thisline[i]))
			}
			pagesets = append(pagesets, thispageset)
		}
	}
	pagesets = pagesets[1:]
	return rulesets, pagesets
}

/* type rule struct {
	firstpage  int
	secondpage int
}
*/
