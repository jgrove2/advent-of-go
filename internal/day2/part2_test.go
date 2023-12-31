package day2

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
		t.Errorf("There has been an error running Part2")
	} else {
		if answer != 2286 {
			t.Errorf("The answer should be 2286, got %d", answer)
		}
	}
}
