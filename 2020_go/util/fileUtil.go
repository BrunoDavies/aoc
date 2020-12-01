package util

import (
	"io/ioutil"
	"strconv"
	"strings"
)

//getValueByLine takes a file, reads the file and splits it by new line into a slice.
func getValueByLine(fileName string) ([]int, error) {
	fileContent, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	allContent := string(fileContent)
	fileLines := strings.Split(allContent, "\n")
	ints := make([]int, len(fileLines))

	//ensures there is no nil values returned
	for _, value := range fileLines {
		tempVal, err := strconv.Atoi(value)

		if err != nil {
			return nil, err
		}
		ints = append(ints, tempVal)
	}

	return ints, nil
}
