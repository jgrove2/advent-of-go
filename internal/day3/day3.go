package day3

import (
	"fmt"
	"github.com/jgrove2/advent-of-go/internal/common"
)

func Day3() {
	fmt.Println("Runnin Day3...")
	fileLoc := "./internal/day3/input.txt"
	part1, err := common.RunPart(Part1, fileLoc)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(part1)
	fmt.Println("Runnin Day3...")
	part2, err := common.RunPart(Part2, fileLoc)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(part2)
}
