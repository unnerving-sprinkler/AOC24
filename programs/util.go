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
