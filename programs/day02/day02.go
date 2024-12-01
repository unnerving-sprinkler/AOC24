package day02

import (
	util "AOC23/programs"
	"strconv"
	"strings"
	"time"
)

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ DAY 02A ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func Day2a(m int) (int, time.Duration) {
	start := time.Now()
	var lines []string
	if m == 0 {
		lines = util.Returnlines("inputdata/day_02/day_02_test.txt")
	} else {
		lines = util.Returnlines("inputdata/day_02/day_02_actual.txt")
	}

	//Todays VARS
	puzzleanswer := 0

	for i := 0; i < len(lines); i++ {
		gamepossible := true
		gameinformationstring := lines[i][(strings.Index(lines[i], ":"))+2:] //Strip The Name Of The Game Off String
		gameinformationstring = gameinformationstring + string("; ")         //Put a Ending ; To Make Our Jobs Easier Later
		var rounds []string
		for {
			rounds = append(rounds, gameinformationstring[:(strings.Index(gameinformationstring, ";"))])    //Append Each Round Info To rounds
			gameinformationstring = gameinformationstring[(strings.Index(gameinformationstring, ";") + 2):] //Remove The Round we Just Counted From The Game
			if len(gameinformationstring) < 4 {                                                             //Break Loop When There Are No More Rounds
				break
			}
		}
		for j := 0; j < len(rounds); j++ {
			roundinfo := ExtractRoundInformation(rounds[j]) //For Each Game Extract Information About How Many Cubes
			if roundinfo[0] > 12 {
				gamepossible = false
			}
			if roundinfo[1] > 13 {
				gamepossible = false
			}
			if roundinfo[2] > 14 {
				gamepossible = false
			}
		}
		if gamepossible {
			puzzleanswer += ExtractGameNumber(lines[i])
		}
	}
	return puzzleanswer, time.Since(start)
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ DAY 02B ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func Day2b(m int) (int, time.Duration) {
	start := time.Now()
	var lines []string
	if m == 0 {
		lines = util.Returnlines("inputdata/day_02/day_02_test.txt")
	} else {
		lines = util.Returnlines("inputdata/day_02/day_02_actual.txt")
	}

	//Todays VARS
	puzzleanswer := 0

	for i := 0; i < len(lines); i++ {
		var gameminimumblocks [3]int
		gameinformationstring := lines[i][(strings.Index(lines[i], ":"))+2:] //Strip The Name Of The Game Off String
		gameinformationstring = gameinformationstring + string("; ")         //Put a Ending ; To Make Our Jobs Easier Later
		var rounds []string
		for {
			rounds = append(rounds, gameinformationstring[:(strings.Index(gameinformationstring, ";"))])    //Append Each Round Info To rounds
			gameinformationstring = gameinformationstring[(strings.Index(gameinformationstring, ";") + 2):] //Remove The Round we Just Counted From The Game
			if len(gameinformationstring) < 4 {                                                             //Break Loop When There Are No More Rounds
				break
			}
		}
		for j := 0; j < len(rounds); j++ {
			roundinfo := ExtractRoundInformation(rounds[j]) //For Each Game Extract Information About How Many Cubes
			if roundinfo[0] > gameminimumblocks[0] {
				gameminimumblocks[0] = roundinfo[0]
			}
			if roundinfo[1] > gameminimumblocks[1] {
				gameminimumblocks[1] = roundinfo[1]
			}
			if roundinfo[2] > gameminimumblocks[2] {
				gameminimumblocks[2] = roundinfo[2]
			}
		}
		puzzleanswer += (gameminimumblocks[0] * gameminimumblocks[1] * gameminimumblocks[2])
	}
	return puzzleanswer, time.Since(start)
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ UTILITY ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
// Return The Game Number (Used For Day 2)
func ExtractGameNumber(s string) int {
	withoutprefix, _ := strings.CutPrefix(s, "Game ") //Remove the Prefix Of The Line
	w := strings.Index(withoutprefix, ":")            //Determine Where the : Char Is
	gamenumberstring := withoutprefix[:w]             //Strip Text String Based On That Information
	gamenumber, _ := strconv.Atoi(gamenumberstring)   //Convert Text To int
	return gamenumber
}

// Returns Information About The Round Fed Into The Function:
// Output Is Generated In This Format | [red, green, blue]
func ExtractRoundInformation(s string) [3]int {
	var colorinfo [3]int
	s = s + string(",  ") // Add , to the end will help us out later
	var roundinfo []string
	for {
		roundinfo = append(roundinfo, s[:strings.Index(s, ",")]) //Information About 1 Cube
		s = s[strings.Index(s, ",")+2:]                          // Trim Info We Just Took Out Of s
		if len(s) < 2 {                                          //Break After We've Completed All Handfulls
			break
		}
	}
	//fmt.Printf("%s\n", roundinfo[0])
	for i := 0; i < len(roundinfo); i++ {
		resultred, foundred := strings.CutSuffix(roundinfo[i], " red")
		resultgreen, foundgreen := strings.CutSuffix(roundinfo[i], " green")
		resultblue, foundblue := strings.CutSuffix(roundinfo[i], " blue")

		if foundred {
			colorinfo[0], _ = strconv.Atoi(resultred)
		} else if foundgreen {
			colorinfo[1], _ = strconv.Atoi(resultgreen)
		} else if foundblue {
			colorinfo[2], _ = strconv.Atoi(resultblue)
		}
	}
	//fmt.Printf("%d|%d|%d", colorinfo[0], colorinfo[1], colorinfo[2])
	return colorinfo
}
