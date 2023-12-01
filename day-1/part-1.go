package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
	"strconv"
)

func extractNumbers(input string) string {
	var result strings.Builder

	for _, char := range input {
		if unicode.IsDigit(char) {
			result.WriteRune(char)
		}
	}

	return result.String()
}

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	currentSum := 0

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		current := scanner.Text()

		currentNumbersOnly := extractNumbers(current)

		firstDigit := string(currentNumbersOnly[0])
		lastDigit  := string(currentNumbersOnly[len(currentNumbersOnly) - 1])
		rowValue, _ := strconv.Atoi(firstDigit + lastDigit)
	
		currentSum += rowValue

	}

	fmt.Println(currentSum)
}