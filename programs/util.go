package until

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Returnlines(filepathactual string, filepathtest string, m int) []string {
	var filepath string
	if m == 0 { //Choose Which Filepath To Follow
		filepath = filepathtest
	} else {
		filepath = filepathactual
	}

	var lines []string
	//Logic
	file, _ := os.Open(filepath)      //Open File Ignoring Errors
	scanner := bufio.NewScanner(file) //Scan File In
	for scanner.Scan() {
		lines = append(lines, scanner.Text()) //Append Each Line Into An Array
	}
	return lines
}

func Splitlines(line string, deliminator string) []string {
	eachcharstring := (strings.Split(line, deliminator))
	var cleanCharString []string
	for i := 0; i < len(eachcharstring); i++ {
		if eachcharstring[i] != "" {
			cleanCharString = append(cleanCharString, eachcharstring[i])
		}
	}
	return cleanCharString
}

func StringtoNumber(str string) int {
	stringasint, _ := strconv.Atoi(str)
	return stringasint
}

func LazyPadding2DSlice(inputarray [][]string, depth int, specialchar string) [][]string {
	var paddedarray [][]string

	topbotslice := make([]string, len(inputarray[0])+2*depth) //CREATE A SLICE FOR THE TOP AND BOTTOM
	for i := range topbotslice {
		topbotslice[i] = specialchar
	}

	leftrightslice := make([]string, depth) //CREATE A SLICE FOR THE LEFT AND RIGHT
	for i := range leftrightslice {
		leftrightslice[i] = specialchar
	}

	for i := 0; i < depth; i++ { //ADD THE SLICES TO THE TOP
		paddedarray = append(paddedarray, topbotslice)
	}

	for i := 0; i < len(inputarray); i++ { //Add Left And Right Padding
		paddedline := append(leftrightslice, inputarray[i]...)
		paddedline = append(paddedline, leftrightslice...)
		paddedarray = append(paddedarray, paddedline)
	}

	for i := 0; i < depth; i++ { //ADD THE SLICES TO THE TOP
		paddedarray = append(paddedarray, topbotslice)
	}

	return paddedarray
}
