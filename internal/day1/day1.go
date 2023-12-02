package day1

import (
	"bufio"
	"fmt"
	"github.com/jgrove2/advent-of-go/internal/common"
	"os"
	"strings"
)

func Day1() {
	fmt.Println("Runnin Day1...")
	fileLoc := "./internal/day1/input.txt"
	part1, err := common.RunPart(Part1, fileLoc)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(part1)
	part2, err := common.RunPart(Part2, fileLoc)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(part2)
}

func Part1(fp *os.File) (int, error) {
	fmt.Println("Running Pt 1...")
	totalValue := 0
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		scannedLine := scanner.Text()
		firstValue := -1
		lastValue := -1
		for _, char := range scannedLine {
			if char >= '0' && char <= '9' {
				if firstValue == -1 {
					firstValue = int(char) - 48
				} else {
					lastValue = int(char) - 48
				}
			}
		}
		if firstValue != -1 {
			if lastValue != -1 {
				totalValue += firstValue*10 + lastValue
			} else {
				totalValue += firstValue*10 + firstValue
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return -1, err
	}

	return totalValue, nil
}

func Part2(fp *os.File) (int, error) {
	fmt.Println("Running Day 1 Pt2...")
	totalValue := 0
	scanner := bufio.NewScanner(fp)
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
	for scanner.Scan() {
		scannedLine := scanner.Text()
		firstValue := -1
		lastValue := -1
	frontSide:
		for i := 0; i < len(scannedLine); i++ {
			if scannedLine[i] >= '0' && scannedLine[i] <= '9' {
				firstValue = int(scannedLine[i]) - 48
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
			if scannedLine[i] >= '0' && scannedLine[i] <= '9' {
				lastValue = int(scannedLine[i]) - 48
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
		if firstValue != -1 {
			if lastValue != -1 {
				i := firstValue*10 + lastValue
				totalValue += i
			} else {
				i := firstValue*10 + firstValue
				totalValue += i
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return -1, err
	}

	return totalValue, nil

}
