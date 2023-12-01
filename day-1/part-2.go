package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
	"strconv"
	"runtime/debug"
)

// I need to swap words for numbers to their corresponding values

func swapWordsForNumbers(input string) string {
	defer func() {
		if r := recover(); r != nil {
				fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()

	// one, two, three, four, five, six, seven, eight, and nine

	var numberMapper = map[string]string{
		"one": "1",	
		"two": "2",
		"three": "3",
		"four": "4",
		"five": "5",
		"six": "6",
		"seven": "7",
		"eight": "8",
		"nine": "9",
	}

	// for key, value := range numberMapper {
	//  	input = strings.ReplaceAll(input, key, value)
	// }

	// I want to replace the first instance of a number word and the last instance

	// I need to loop through the input and find the first instance of a number word
	// I need to loop through the input and find the last instance of a number word

	returnInput := input

	firstInstanceFound := false
	lastInstanceFound := false


	for index, _ := range input {		
		if firstInstanceFound {
			break
		}

		charsToCheck := input[0:index]

		// check if charsToCheck contains a number word

		for key, _ := range numberMapper {
			if strings.Contains(charsToCheck, key) {
				firstInstanceFound = true

				returnInput = strings.Replace(returnInput, key, numberMapper[key], 1)
				break
			}
		}

	}
	
	// then do the same for the last instance

	for index, _ := range returnInput {
		if lastInstanceFound {
			break
		}

		charsToCheck := returnInput[index:len(returnInput)]

		for key, _ := range numberMapper {

			if strings.Contains(charsToCheck, key) {

				lastInstanceFound = true

				returnInput = strings.Replace(returnInput, key, numberMapper[key], -1)

				break

			}

		}

	}

	return returnInput

}

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

		parsedCurrent := swapWordsForNumbers(current)

		currentNumbersOnly := extractNumbers(parsedCurrent)

		firstDigit := string(currentNumbersOnly[0])
		lastDigit  := string(currentNumbersOnly[len(currentNumbersOnly) - 1])
		rowValue, _ := strconv.Atoi(firstDigit + lastDigit)

		
		fmt.Println(current)

		fmt.Println(firstDigit)
		fmt.Println(lastDigit)
		
		fmt.Println(rowValue)
		fmt.Println("------")
	
		currentSum += rowValue

	}

	fmt.Println("------ TOTAL ------")
	fmt.Println(currentSum)
}

// 54165 too low

// 54199 too low

// 54205

// 54206 not right

// 54208 is right :D

// 54212 too high