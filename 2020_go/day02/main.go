package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	allLines, err := getStringsByLine("input.txt")

	fmt.Println("Part A solution: ", partASolution(allLines), err)
	fmt.Println("Part B solution: ", partBSolution(allLines), err)
}

func partASolution(fullArray []string) int{
	correctPasswords := 0

	for _, oneLine := range fullArray{
		if isPasswordCorrect(strings.Fields(oneLine)) {
			correctPasswords++
		}
	}

	return correctPasswords
}

func isPasswordCorrect(line []string) bool{
	letterAppears := timesLetterInPassword(line[1], line[2])

	rangeForLetter := line[0]
	rangeForLetterSplit := strings.Split(rangeForLetter, "-")

	lowerRange, _ := strconv.Atoi(rangeForLetterSplit[0])
	upperRange, _ := strconv.Atoi(rangeForLetterSplit[1])

	if (letterAppears >= lowerRange && letterAppears <= upperRange) {
		return true
	}

	return false
}

func timesLetterInPassword(letter string, password string) int{
	counter := 0
	for _, char := range password{
		if strings.Trim(letter, ":") == string(char) {
			counter++
		}
	}
	return counter
}

func partBSolution(fullArray []string) int {
	correctPasswords := 0
	
	for _, oneLine := range fullArray{
		if isPasswordCorrectPartB(strings.Fields(oneLine)){
			correctPasswords++
		}
	}
	return correctPasswords
}

func isPasswordCorrectPartB(line []string) bool {
	letter := strings.Trim(line[1], ":")

	rangeForLetter := line[0]
	rangeForLetterSplit := strings.Split(rangeForLetter, "-")

	lowerRange, _ := strconv.Atoi(rangeForLetterSplit[0])
	upperRange, _ := strconv.Atoi(rangeForLetterSplit[1])

	lowerBool := letter == string(line[2][lowerRange-1])
	upperBool := letter == string(line[2][upperRange-1])
	if lowerBool != upperBool {
		return true
	}
	return false
}








//This is supposed to go in /util but I do not know how to import it to this class
//getValueByLine takes a file, reads the file and splits it by new line into a slice.
func getStringsByLine(fileName string) ([]string, error) {
	fileContent, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	allContent := string(fileContent)
	fileLines := strings.Split(strings.TrimSpace(allContent), "\n")
	fullLine := make([]string, 0, len(fileLines))

	//ensures there is no nil values returned
	for _, oneLine := range fileLines {

		if err != nil {
			return nil, err
		}
		fullLine = append(fullLine, oneLine)
	}
	return fullLine, nil
}