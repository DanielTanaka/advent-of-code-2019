package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	readIntcodeFromFile("input.txt")
}

func readIntcodeFromFile(fileName string) {
	file, error := os.Open(fileName)

	if error != nil {
		log.Fatal(error)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	values := make([]int, 0)
	convertToInt(strings.Split(line, ","), &values)

	step := 4
	for i := 0; i < len(values); i = i + step {
		haltOpcode := 99
		if values[i] == haltOpcode {
			break
		}

		firstPosition := values[i+1]
		secondPosition := values[i+2]
		storagePosition := values[i+3]

		sumOpcode := 1
		multiplicationOpcode := 2

		var result int
		if values[i] == sumOpcode {
			result = values[firstPosition] + values[secondPosition]
		} else if values[i] == multiplicationOpcode {
			result = values[firstPosition] * values[secondPosition]
		}
		values[storagePosition] = result
	}
	file.Close()
	fmt.Println(values)
}

func convertToInt(stringArray []string, intArray *[]int) {
	for _, value := range stringArray {
		intValue, conversionError := strconv.Atoi(value)
		if conversionError != nil {
			log.Fatal(conversionError)
		}
		*intArray = append(*intArray, intValue)
	}
}
