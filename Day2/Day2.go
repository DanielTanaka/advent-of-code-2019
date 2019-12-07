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
	output, noun, verb := findNounVerb(19690720)
	result := applyFormula(noun, verb)
	fmt.Printf("Output = %d, Noun = %d, Verb = %d\n", output, noun, verb)
	fmt.Printf("Result: %d\n", result)
}

func applyFormula(noun, verb int) int {
	return 100*noun + verb
}

func findNounVerb(expectedOutput int) (int, int, int) {
	noun, verb := 0, 0
	maximumValue := 99
	fileName := "input.txt"
	output := readIntcodeAndApplyNounVerb(fileName, noun, verb)

	for noun < maximumValue && output != expectedOutput {
		noun++
		for verb < maximumValue && output != expectedOutput {
			verb++
			output = readIntcodeAndApplyNounVerb(fileName, noun, verb)
		}
		if output != expectedOutput { //maximum value was reached instead...
			verb = 0
		}
	}

	return output, noun, verb
}

func readIntcodeAndApplyNounVerb(fileName string, noun, verb int) int {
	file, error := os.Open(fileName)

	if error != nil {
		log.Fatal(error)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	values := make([]int, 0)
	convertToIntArray(strings.Split(line, ","), &values)

	values[1] = noun
	values[2] = verb

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
	output := values[0]
	return output
}

func convertToIntArray(stringArray []string, intArray *[]int) {
	for _, stringValue := range stringArray {
		intValue, conversionError := strconv.Atoi(stringValue)
		if conversionError != nil {
			log.Fatal(conversionError)
		}
		*intArray = append(*intArray, intValue)
	}
}
