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
		totalFuel = totalFuel + calculateTotalFuel(currentValue)
	}

	file.Close()
	return totalFuel
}

func calculateTotalFuel(mass int) int {
	requiredFuel := calculateRequiredFuelForMass(mass)
	furtherFuel := calculateRequiredFuelForMass(requiredFuel)

	for furtherFuel > 0 {
		requiredFuel = requiredFuel + furtherFuel
		furtherFuel = calculateRequiredFuelForMass(furtherFuel)
	}
	return requiredFuel
}

func calculateRequiredFuelForMass(mass int) int {
	return int(mass/3) - 2
}
