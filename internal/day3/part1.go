package day3

import (
	"bufio"
	"fmt"
	"os"
)

func Part1(fp *os.File) (int, error) {
	fmt.Println("Running Pt 1...")
	scanner := bufio.NewScanner(fp)
	var inputArray [3][]rune
	total := 0
	startChecking := false
	for scanner.Scan() {
		scannedLine := scanner.Text()
		inputArray[2] = inputArray[1]
		inputArray[1] = inputArray[0]
		inputArray[0] = []rune(scannedLine)
		if startChecking {
			total += checkRow(inputArray)
		} else {
			startChecking = true
		}
	}
	inputArray[2] = inputArray[1]
	inputArray[1] = inputArray[0]
	inputArray[0] = nil
	total += checkRow(inputArray)
	return total, nil
}
func checkRow(inputArray [3][]rune) int {
	num := 0
	isValid := false
	total := 0
	for index, r := range inputArray[1] {
		if isDigit(r) {
			if isValid {
				num = num*10 + int(r) - 48
			} else if checkForSym(index, inputArray) {
				isValid = true
				num = num*10 + int(r) - 48
			} else {
				num = num*10 + int(r) - 48
			}
		} else if isValid {
			isValid = false
			total += num
			num = 0
		} else if num != 0 {
			num = 0
		}
	}
	if isValid && num != 0 {
		total += num
	}
	return total
}

func isAlph(r rune) bool {
	if r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z' {
		return true
	}
	return false
}

func isDigit(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}

func isSymb(index int, arr []rune) bool {
	if !isDigit(arr[index]) && !isAlph(arr[index]) && arr[index] != '.' {
		return true
	}
	return false
}

func checkForSym(index int, arr [3][]rune) bool {
	if len(arr[0]) != 0 {
		if isSymb(index, arr[0]) {
			return true
		} else if index+1 < len(arr[0]) && isSymb(index+1, arr[0]) {
			return true
		} else if index-1 >= 0 && isSymb(index-1, arr[0]) {
			return true
		}
	}
	if len(arr[2]) != 0 {
		if isSymb(index, arr[2]) {
			return true
		} else if index+1 < len(arr[2]) && isSymb(index+1, arr[2]) {
			return true
		} else if index-1 >= 0 && isSymb(index-1, arr[2]) {
			return true
		}
	}
	if index+1 < len(arr[1]) && isSymb(index+1, arr[1]) {
		return true
	}
	if index-1 >= 0 && isSymb(index-1, arr[1]) {
		return true
	}
	return false
}
