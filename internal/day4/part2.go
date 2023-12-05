package day4

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Part2(fp *os.File) (int, error) {
	fmt.Println("Running Pt 2...")
	scanner := bufio.NewScanner(fp)
	total := 0
	cardNumber := 1
	copies := map[int]int{1: 0}
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
				cardWinnings += 1
			}
		}
		for i := 1; i <= cardWinnings; i++ {
			copies[cardNumber+i] += 1 * (1 + copies[cardNumber])
		}

		total += 1 + cardWinnings*(1+copies[cardNumber])
		cardNumber += 1
	}
	return total, nil
}
