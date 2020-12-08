package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var neededFeilds = []string{
"byr",
"iyr",
"eyr",
"hgt",
"hcl",
"ecl",
"pid",
}

func main(){
	allLinesByEmptyLine, err := getStringsByEmptyLine("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part A answer: ", partASolution(allLinesByEmptyLine))
	fmt.Println("Part B answer: ", partBSolution(allLinesByEmptyLine))
}

func partASolution(allLinesByEmptyLine []string) int {
	validPassports := 0

	for _, oneEntry := range allLinesByEmptyLine{
		if isValidPassport(oneEntry){
			validPassports++
		}
	}
	return validPassports
}

func isValidPassport(entry string) bool {
	neededFeilds := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}

	for _, feilds := range neededFeilds{
		if !strings.Contains(entry, feilds){
			return false
		}
	}
	return true
}

func partBSolution(allLinesByEmptyLine []string) int {
	validPassports := 0

	for _, oneEntry := range allLinesByEmptyLine{
		if isValidPassport(oneEntry){
			if passportValuesAreValid(oneEntry){
				validPassports++
			}
		}
	}
	return validPassports
}

func passportValuesAreValid(entry string) bool {
	counter := 0
	splitEntry := strings.Fields(entry)
	for _, dataPoint := range splitEntry {
		dataPoint := strings.Split(dataPoint, ":")
		switch dataPoint[0] {

			case neededFeilds[0]: //byr case
				enteryValue, _ := strconv.Atoi(dataPoint[1])
				if !(1920 <= enteryValue && enteryValue <= 2002) {
					counter++
				}

			case neededFeilds[1]://iyr case
				enteryValue, _ := strconv.Atoi(dataPoint[1])
				if !(2010 <= enteryValue && enteryValue <= 2020) {
					counter++
				}

			case neededFeilds[2]://eyr case
				enteryValue, _ := strconv.Atoi(dataPoint[1])
				if !(2020 <= enteryValue && enteryValue <= 2030) {
					counter++
				}

			case neededFeilds[3]://hgt case
				if strings.Contains(dataPoint[1], "cm") {
					onlyValueInString := strings.Split(dataPoint[1], "cm")
					onlyValueInInt, _ := strconv.Atoi(onlyValueInString[0])
					if !(150 <= onlyValueInInt && onlyValueInInt <= 193) {
						counter++
					}
				}

				if strings.Contains(dataPoint[1], "in") {
					onlyValueInString := strings.Split(dataPoint[1], "in")
					onlyValueInInt, _ := strconv.Atoi(onlyValueInString[0])
					if !(59 <= onlyValueInInt && onlyValueInInt <= 76) {
						counter++
					}
				}


			case neededFeilds[4]: //hcl case
				matched, _ := regexp.MatchString("^#[0-9a-f]{6}$", dataPoint[1])
				if !matched{
					counter++
				}

			case neededFeilds[5]: //ecl case
				allowedEyeColour := []string {
					"amb", "blu", "brn", "gry", "grn", "hzl", "oth",
				}
				containsFlag := 0
				for _, colour := range allowedEyeColour {
					if colour == dataPoint[1]{
						containsFlag = 1
					}
				}
				if containsFlag == 0{
					counter++
				}

			case neededFeilds[6]: //pid case
				matched, _ := regexp.MatchString("^[0-9]{9}$", dataPoint[1])
				if !matched{
					counter++
				}
		}
	}
	if counter == 0 {
		return true
	}
	return false
}




//This is supposed to go in /util but I do not know how to import it to this class
//getValueByLine takes a file, reads the file and splits it by new line into a slice.
func getStringsByEmptyLine(fileName string) ([]string, error) {
	fileContent, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	allContent := string(fileContent)
	fileLines := strings.Split(strings.TrimSpace(allContent), "\n\n")
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
