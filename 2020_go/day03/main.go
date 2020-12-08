package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)


func main(){
	allLines, err := getStringsByLine("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part A answer: ", partASolution(allLines))
	fmt.Println("Part B answer: ", partBSolution(allLines))
}

func partASolution(allLines []string) int {
	treeCounter := 0
	xPos := 0
	yPos := 0
	treeValue := "#"

	for yPos < len(allLines){
		xPos = xPos % len(allLines[0])

		if treeValue == string(allLines[yPos][xPos]){
			treeCounter++
		}
		xPos = xPos + 3
		yPos++
	}
	return treeCounter
}

func genericTreeCrashCalc(allLines []string, yChange int, xChange int) int {
	treeCounter := 0
	xPos := 0
	yPos := 0
	treeValue := "#"

	for yPos < len(allLines){
		xPos = xPos % len(allLines[0])

		if treeValue == string(allLines[yPos][xPos]){
			treeCounter++
		}
		xPos = xPos + xChange
		yPos = yPos + yChange
	}
	return treeCounter

}

func partBSolution(allLines []string) int {
	totalTreeMulti := 1
	partBSlopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	for _, changeValues := range partBSlopes{
		totalTreeMulti = totalTreeMulti * genericTreeCrashCalc(allLines, changeValues[1], changeValues[0])
	}
	return totalTreeMulti
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