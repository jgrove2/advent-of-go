package day4

import (
	"fmt"
	"github.com/jgrove2/advent-of-go/internal/common"
	"testing"
)

func TestPart2(t *testing.T) {
	fileLoc := "./part2_test_input.txt"
	answer, err := common.RunPart(Part2, fileLoc)
	if err != nil {
		fmt.Println(err)
		t.Errorf("There has been an error running Part1")
	} else {
		if answer != 30 {
			t.Errorf("The answer should be 30, got %d", answer)
		}
	}
}
