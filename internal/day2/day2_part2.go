package day2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Part2(fp *os.File) (int, error) {
	fmt.Println("Running Pt 2...")
	scanner := bufio.NewScanner(fp)
	total := 0
	redBlocks := 0
	greenBlocks := 0
	blueBlocks := 0
	for scanner.Scan() {
		scannedLine := scanner.Text()
		subStrings := strings.Split(scannedLine[8:], ";")
		for _, game := range subStrings {
			blocks := strings.Split(game, ",")
			for _, block := range blocks {
				num, err := ParseColor(block)
				if err != nil {
					return -1, err
				}
				if num[0] > redBlocks {
					redBlocks = num[0]
				}
				if num[1] > greenBlocks {
					greenBlocks = num[1]
				}
				if num[2] > blueBlocks {
					blueBlocks = num[2]
				}
			}
		}
		total += redBlocks * greenBlocks * blueBlocks
		redBlocks, greenBlocks, blueBlocks = 0, 0, 0
	}

	if err := scanner.Err(); err != nil {
		return -1, err
	}

	return total, nil
}
