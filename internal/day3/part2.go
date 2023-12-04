package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Part2(fp *os.File) (int, error) {
	fmt.Println("Running Pt 2...")
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
			total += checkRow2(inputArray)
		} else {
			startChecking = true
		}
	}
	inputArray[2] = inputArray[1]
	inputArray[1] = inputArray[0]
	inputArray[0] = nil
	total += checkRow2(inputArray)
	return total, nil
}
func checkRow2(inputArray [3][]rune) int {
	total := 0
	for index, r := range inputArray[1] {
		if r == '*' {
			total += checkForGear(index, inputArray)
		}
	}
	return total
}
func parseNum(arr []rune, start int) int {
	number := string(arr[start])
	right := 1
	left := -1

	for {
		if start+right < len(arr) && isDigit(arr[start+right]) {
			number = number + string(arr[start+right])
			right = right + 1
		}
		if left+start >= 0 && isDigit(arr[start+left]) {
			number = string(arr[start+left]) + number
			left = left - 1
		}
		if (start+right >= len(arr) || !isDigit(arr[start+right])) && (start+left < 0 || !isDigit(arr[start+left])) {
			break
		}
	}
	conv, _ := strconv.Atoi(number)
	return conv
}

func checkForGear(index int, arr [3][]rune) int {
	first := 0
	last := 0
	if len(arr[0]) != 0 {
		if isDigit(arr[0][index]) {
			if first != 0 {
				last = parseNum(arr[0], index)
			} else {
				first = parseNum(arr[0], index)
			}
		}
		if index+1 < len(arr[0]) && !isDigit(arr[0][index]) && isDigit(arr[0][index+1]) {
			if first != 0 {
				last = parseNum(arr[0], index+1)
			} else {
				first = parseNum(arr[0], index+1)
			}
		}
		if index-1 >= 0 && !isDigit(arr[0][index]) && isDigit(arr[0][index-1]) {
			if first != 0 {
				last = parseNum(arr[0], index-1)
			} else {
				first = parseNum(arr[0], index-1)
			}
		}
	}
	if len(arr[2]) != 0 {
		if isDigit(arr[2][index]) {
			if first != 0 {
				last = parseNum(arr[2], index)
			} else {
				first = parseNum(arr[2], index)
			}
		}
		if index+1 < len(arr[2]) && !isDigit(arr[2][index]) && isDigit(arr[2][index+1]) {
			if first != 0 {
				last = parseNum(arr[2], index+1)
			} else {
				first = parseNum(arr[2], index+1)
			}
		}
		if index-1 >= 0 && !isDigit(arr[2][index]) && isDigit(arr[2][index-1]) {
			if first != 0 {
				last = parseNum(arr[2], index-1)
			} else {
				first = parseNum(arr[2], index-1)
			}
		}
	}
	if index+1 < len(arr[1]) && isDigit(arr[1][index+1]) {
		if first != 0 {
			last = parseNum(arr[1], index+1)
		} else {
			first = parseNum(arr[1], index+1)
		}
	}
	if index-1 >= 0 && isDigit(arr[1][index-1]) {
		if first != 0 {
			last = parseNum(arr[1], index-1)
		} else {
			first = parseNum(arr[1], index-1)
		}
	}
	return first * last
}
