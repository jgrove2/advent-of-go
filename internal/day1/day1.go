package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Part1() int {
	fmt.Println("Running Day 1 Pt 1...")
	file, err := os.Open("./internal/day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	totalValue := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scannedLine := scanner.Text()
		firstValue := ""
		lastValue := ""
		for _, char := range scannedLine {
			if unicode.IsNumber(char) {
				if firstValue == "" {
					firstValue = string(char)
				} else {
					lastValue = string(char)
				}
			}
		}
		if firstValue != "" {
			if lastValue != "" {
				i, err := strconv.Atoi(firstValue + lastValue)
				if err != nil {
					log.Fatal(err)
				}
				totalValue += i
			} else {
				i, err := strconv.Atoi(firstValue + firstValue)
				if err != nil {
					log.Fatal(err)
				}
				totalValue += i
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return totalValue
}

func Part2() int {
	fmt.Println("Running Day 1 Pt2...")
	file, err := os.Open("./internal/day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	totalValue := 0
	scanner := bufio.NewScanner(file)
	speltNumbers := [10]string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
	numbers := map[string]int{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}
	for scanner.Scan() {
		scannedLine := scanner.Text()
		firstValue := 0
		lastValue := 0
	frontSide:
		for i := 0; i < len(scannedLine); i++ {
			n, ok := numbers[string(scannedLine[i])]
			if ok {
				firstValue = n
				break frontSide
			}
			frontWindow := scannedLine[0 : i+1]
			for i, str := range speltNumbers {
				if strings.Contains(frontWindow, str) {
					firstValue = i
					break frontSide
				}
			}
		}
	backSide:
		for i := len(scannedLine) - 1; i >= 0; i-- {
			n, ok := numbers[string(scannedLine[i])]
			if ok {
				lastValue = n
				break backSide
			}
			backWindow := scannedLine[i:]
			for i, str := range speltNumbers {
				if strings.Contains(backWindow, str) {
					lastValue = i
					break backSide
				}
			}
		}
		if firstValue != 0 {
			if lastValue != 0 {
				i, err := strconv.Atoi(fmt.Sprintf("%d%d", firstValue, lastValue))
				if err != nil {
					log.Fatal(err)
				}
				totalValue += i
			} else {
				i, err := strconv.Atoi(fmt.Sprintf("%d%d", firstValue, firstValue))
				if err != nil {
					log.Fatal(err)
				}
				totalValue += i
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return totalValue

}
