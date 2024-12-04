// /programs/day02.go

package day02

import (
	util "AOC24/programs"
	"math"
	"time"
)

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ DAY 02A ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func Day2a(m int) (int, time.Duration) {
	start := time.Now()
	lines := util.Returnlines("inputdata/day_02/day_02_actual.txt", "inputdata/day_02/day_02_test.txt", m)

	//Setup VARS For Today
	var report [][]int
	score := 0

	//Days Program

	for i := 0; i < len(lines); i++ { //Get Data Into Reasonable Format
		currentline := util.Splitlines(lines[i], " ")
		var currentlineasint []int
		for j := 0; j < len(currentline); j++ {
			currentlineasint = append(currentlineasint, util.StringtoNumber(currentline[j]))
		}
		report = append(report, currentlineasint)
	}

	for i := 0; i < len(report); i++ { //Loop Through Each Result
		isvalidresult := true

		if report[i][0] < report[i][1] { //Are All Numbers Moving Up
			for j := 1; j < len(report[i]); j++ {
				if report[i][j] < report[i][j-1] {
					isvalidresult = false

				}
			}
		} else if report[i][0] > report[i][1] { //Are All Numbers Moving Down
			for j := 1; j < len(report[i]); j++ {
				if report[i][j] > report[i][j-1] {
					isvalidresult = false

				}
			}
		} else { // Are Both First Numbers The Same
			isvalidresult = false

		}

		for j := 1; j < len(report[i]); j++ { //Is EACH Char Moving Up or Down By between 1 and 3 Units
			if math.Abs((float64(report[i][j]-report[i][j-1]))) > 0 && math.Abs((float64(report[i][j]-report[i][j-1]))) < 4 {
				//Do nothing
			} else {
				isvalidresult = false
			}

		}

		if isvalidresult { //Add Score Points
			score += 1
		}

	}

	return score, time.Since(start)
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ DAY 02B ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func Day2b(m int) (int, time.Duration) {
	start := time.Now()
	lines := util.Returnlines("inputdata/day_02/day_02_actual.txt", "inputdata/day_02/day_02_test.txt", m)

	//Setup VARS For Today
	var report [][]int
	score := 0

	//Days Program

	for i := 0; i < len(lines); i++ { //Get Data Into Reasonable Format
		currentline := util.Splitlines(lines[i], " ")
		var currentlineasint []int
		for j := 0; j < len(currentline); j++ {
			currentlineasint = append(currentlineasint, util.StringtoNumber(currentline[j]))
		}
		report = append(report, currentlineasint)
	}

	for i := 0; i < len(report); i++ { //Loop Through Each Result
		goodresult := false
		testreport := report[i]

		for j := 0; j < len(testreport); j++ { //Remove One Piece Of Each Array
			testslice := make([]int, 0, len(testreport)-1)
			testslice = append(testslice, testreport[:j]...)
			testslice = append(testslice, testreport[j+1:]...)

			if checkreport(testslice) {
				goodresult = true
			}
		}

		if goodresult {
			score += 1
		}

	}

	return score, time.Since(start)

}

func checkreport(report []int) bool {
	isvalidresult := true

	if report[0] < report[1] { //Are All Numbers Moving Up
		for j := 1; j < len(report); j++ {
			if report[j] < report[j-1] {
				isvalidresult = false

			}
		}
	} else if report[0] > report[1] { //Are All Numbers Moving Down
		for j := 1; j < len(report); j++ {
			if report[j] > report[j-1] {
				isvalidresult = false

			}
		}
	} else { // Are Both First Numbers The Same
		isvalidresult = false

	}

	for j := 1; j < len(report); j++ { //Is EACH Char Moving Up or Down By between 1 and 3 Units
		if math.Abs((float64(report[j]-report[j-1]))) > 0 && math.Abs((float64(report[j]-report[j-1]))) < 4 {
			//Do nothing
		} else {
			isvalidresult = false
		}

	}

	if isvalidresult == true { //Add Score Points
		return true
	}
	return false
}
