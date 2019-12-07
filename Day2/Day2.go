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
		if values[i] == 99 {
			break
		}

		firstPosition := values[i+1]
		secondPosition := values[i+2]
		storagePosition := values[i+3]
		var result int
		if values[i] == 1 {
			result = values[firstPosition] + values[secondPosition]
		} else if values[i] == 2 {
			result = values[firstPosition] * values[secondPosition]
		}
		values[storagePosition] = result
	}

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
