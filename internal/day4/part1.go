package day4

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Part1(fp *os.File) (int, error) {
	fmt.Println("Running Pt 1...")
	scanner := bufio.NewScanner(fp)
	total := 0
	for scanner.Scan() {
		scannedLine := scanner.Text()
		cardWinnings := 0
		splt := strings.Split(scannedLine, " | ")
		winningNumbers := strings.Split(splt[0], ": ")
		r := regexp.MustCompile(("(\\d+)"))
		winningNumbers = r.FindAllString(winningNumbers[1], -1)
		otherNumbers := r.FindAllString(splt[1], -1)
		winningMap := make(map[string]bool)
		for _, num := range winningNumbers {
			winningMap[num] = true
		}
		for _, num := range otherNumbers {
			if winningMap[num] {
				if cardWinnings == 0 {
					cardWinnings = 1
				} else {
					cardWinnings *= 2
				}
			}
		}
		total += cardWinnings
	}
	return total, nil
}
