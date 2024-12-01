package until

import (
	"bufio"
	"os"
)

func Returnlines(filepath string) []string {
	var lines []string

	//Logic
	file, _ := os.Open(filepath)      //Open File Ignoring Errors
	scanner := bufio.NewScanner(file) //Scan File In
	for scanner.Scan() {
		lines = append(lines, scanner.Text()) //Append Each Line Into An Array
	}
	return lines
}
