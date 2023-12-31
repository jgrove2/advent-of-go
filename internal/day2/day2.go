package day2

import (
	"fmt"
	"github.com/jgrove2/advent-of-go/internal/common"
)

func Day2() {
	fmt.Println("Runnin Day2...")
	fileLoc := "./internal/day2/input.txt"
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
