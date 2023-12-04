package day2

import (
	"fmt"
	"github.com/jgrove2/advent-of-go/internal/common"
	"testing"
)

func TestPart1(t *testing.T) {
	fileLoc := "./part1_test_input.txt"
	answer, err := common.RunPart(Part1, fileLoc)
	if err != nil {
		fmt.Println(err)
		t.Errorf("There has been an error running Part1")
	} else {
		if answer != 8 {
			t.Errorf("The answer should be 8, got %d", answer)
		}
	}
}
