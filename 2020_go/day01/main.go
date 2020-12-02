package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const goal = 2020

func main() {
	allInts, err := getValueByLine("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part a: ", roughPartASolution(allInts, goal), "Part b: ", roughPartBSolution(allInts, goal))
}

//This is supposed to go in /util but I do not know how to import it to this class
//getValueByLine takes a file, reads the file and splits it by new line into a slice.
func getValueByLine(fileName string) ([]int, error) {
	fileContent, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	allContent := string(fileContent)
	fileLines := strings.Split(strings.TrimSpace(allContent), "\n")
	ints := make([]int, 0, len(fileLines))

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

//My basic answer
func roughPartASolution(values []int, sum int) int {
	for i, valueA := range values {
		for _, valueB := range values[i+1:] {
			if sum == valueA+valueB {
				return valueA * valueB
			}
		}
	}
	return 0
}

//Solution A: from https://github.com/thlacroix/goadvent - so much cleaner / faster / interesting than mine
func findTwoValuesThatSum(values []int, sum int) int {
	heatmap := make(map[int]bool, len(values))

	for _, v := range values {
		if heatmap[v] {
			return v * (sum - v)
		}
		heatmap[sum-v] = true
	}
	return 0
}

func roughPartBSolution(values []int, sum int) int {
	for i, valueA := range values {
		for j, valueB := range values[i+1:] {
			for _, valueC := range values[j+1:] {
				if sum == valueA+valueB+valueC {
					return valueA * valueB * valueC
				}
			}
		}
	}
	return 0
}

//Solution B: from https://github.com/thlacroix/goadvent - couldn't do this in GoLang (did not enter answer for star)
//Included it since again a really clean solution that I learnt a lot from
func threeSum(data []int, target int) int {
	lookup := make(map[int]bool, len(data))

	for _, v := range data {
		lookup[v] = true
	}

	for i, x := range data {
		for _, y := range data[i+1:] {
			z := target - x - y
			if lookup[z] {
				return x * y * z
			}
		}
	}
	return 0
}
