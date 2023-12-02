package day2

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Part1(fp *os.File) (int, error) {
	fmt.Println("Running Pt 1...")
	scanner := bufio.NewScanner(fp)
	total := 0
	gameIndex := 1
	redBlocks := 12
	greenBlocks := 13
	blueBlocks := 14
	for scanner.Scan() {
		scannedLine := scanner.Text()
		subStrings := strings.Split(scannedLine, ";")
	subgame:
		for _, game := range subStrings {
			blocks := strings.Split(game, ",")
			for _, block := range blocks {
				num, err := ParseColor(block)
				if err != nil {
					return -1, err
				}
				if num[0] > redBlocks || num[1] > greenBlocks || num[2] > blueBlocks {
					total -= gameIndex
					break subgame
				}
			}
		}
		total += gameIndex

		gameIndex += 1
	}

	if err := scanner.Err(); err != nil {
		return -1, err
	}

	return total, nil
}

func ParseColor(blocks string) ([3]int, error) {
	number := regexp.MustCompile("(\\d+)")
	color := regexp.MustCompile("([a-z]+)")
	match_number := number.FindStringSubmatch(blocks)
	match_color := color.FindStringSubmatch(blocks)
	num, err := strconv.Atoi(match_number[0])
	if err != nil {
		return [3]int{0, 0, 0}, err
	}
	switch match_color[0] {
	case "red":
		return [3]int{num, 0, 0}, nil
	case "green":
		return [3]int{0, num, 0}, nil
	case "blue":
		return [3]int{0, 0, num}, nil
	default:
		return [3]int{0, 0, 0}, nil
	}
}
