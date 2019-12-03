package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	result := calculateResultFromFile("input.txt")
	fmt.Print(result)
}

func calculateResultFromFile(fileName string) int {
	file, error := os.Open(fileName)

	if error != nil {
		log.Fatal(error)
	}

	scanner := bufio.NewScanner(file)
	totalFuel := 0
	for scanner.Scan() {
		currentValue, conversionError := strconv.Atoi(scanner.Text())
		if conversionError != nil {
			log.Fatal(conversionError)
		}
		totalFuel = totalFuel + calculateFuel(currentValue)
	}

	file.Close()
	return totalFuel
}

func calculateFuel(mass int) int {
	return int(mass/3) - 2
}
