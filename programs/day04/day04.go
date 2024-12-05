// /programs/day04.go

package day04

import (
	util "AOC24/programs"
	"fmt"
	"regexp"
	"strings"
	"time"
)

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ DAY 04A ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func Day4a(m int) (int, time.Duration) {
	start := time.Now()
	lines := util.Returnlines("inputdata/day_04/day_04_actual.txt", "inputdata/day_04/day_04_test.txt", m)

	//Setup VARS For Today
	score := 0
	var inputarray [][]string
	var diagwords []string

	//Days Program
	for i := 0; i < len(lines); i++ { //Search For Across Words Using Basic REGEX
		re := regexp.MustCompile(`XMAS`)
		score += len(re.FindAllString(lines[i], -1))
		re = regexp.MustCompile(`SAMX`)
		score += len(re.FindAllString(lines[i], -1))

		//While Were Still In This Loop -- Break Lines (Strings) Into Arrays Of Charactors
		thisline := util.Splitlines(lines[i], "")
		inputarray = append(inputarray, thisline)
	}

	rotatedarray := rotatematrix(inputarray) //Flip Array

	for i := 0; i < len(rotatedarray); i++ { //Search For Up Down (Currently Across Because Of Flip)
		rotatedstring := strings.Join(rotatedarray[i], "")
		re := regexp.MustCompile(`XMAS`)
		score += len(re.FindAllString(rotatedstring, -1))
		re = regexp.MustCompile(`SAMX`)
		score += len(re.FindAllString(rotatedstring, -1))
	}

	//Not proud of this...
	paddedarray := util.LazyPadding2DSlice(rotatedarray, 5, "*")
	for i := 4; i < len(paddedarray)-4; i++ {
		for j := 4; j < len(paddedarray[i])-4; j++ {
			//fmt.Println(paddedarray[i][j])
			if paddedarray[i][j] == "X" || paddedarray[i][j] == "M" || paddedarray[i][j] == "A" || paddedarray[i][j] == "S" {

				diagone := []string{paddedarray[i][j], paddedarray[i-1][j+1], paddedarray[i-2][j+2], paddedarray[i-3][j+3]}
				diagtwo := []string{paddedarray[i][j], paddedarray[i+1][j+1], paddedarray[i+2][j+2], paddedarray[i+3][j+3]}

				diagwords = append(diagwords, strings.Join(diagone, ""))
				diagwords = append(diagwords, strings.Join(diagtwo, ""))

			}
		}
	}

	for i := 0; i < len(diagwords); i++ { //Search For Up Down (Currently Across Because Of Flip)
		re := regexp.MustCompile(`XMAS|SAMX`)
		score += len(re.FindAllString(diagwords[i], -1))
	}

	fmt.Println(diagwords)
	return score, time.Since(start)
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ DAY 04B ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func Day4b(m int) (int, time.Duration) {
	start := time.Now()
	//lines := util.Returnlines("inputdata/day_04/day_04_actual.txt", "inputdata/day_04/day_04_test.txt", m)

	//Setup VARS For Today
	score := 0

	//Days Program

	return score, time.Since(start)
}

// rotatematrix returns the matrix that you inputted, flipped on it's side.
// THE MATRIX MUST BE A RECTANGLE
func rotatematrix(matrix [][]string) [][]string {
	var returnarray [][]string
	for i := 0; i < len(matrix[0]); i++ {
		var thisrow []string
		for j := 0; j < len(matrix); j++ {
			thisrow = append(thisrow, matrix[j][i])
		}
		returnarray = append(returnarray, thisrow)
	}
	return returnarray
}
