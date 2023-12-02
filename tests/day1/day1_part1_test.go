package tests

import (
	"fmt"
	"github.com/jgrove2/advent-of-go/internal/common"
	"github.com/jgrove2/advent-of-go/internal/day1"
	"testing"
)

func TestPart1(t *testing.T) {
	fileLoc := "../../internal/day1/part1_test_input.txt"
	answer, err := common.RunPart(day1.Part1, fileLoc)
	if err != nil {
		fmt.Println(err)
		t.Errorf("There has been an error running Part1")
	} else {
		if answer != 142 {
			t.Errorf("The answer should be 142, got %d", answer)
		}
	}
}
