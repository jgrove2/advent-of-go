package day3

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
		if answer != 467835 {
			t.Errorf("The answer should be 467835, got %d", answer)
		}
	}
}
func TestParse(t *testing.T) {
	strin := []rune{'1', '2', '3', '4', '4'}
	answer := parseNum(strin, 2)
	if answer != 12344 {
		t.Errorf("The parsedNum should be 12344, got %d", answer)
	}
}
func TestPart21(t *testing.T) {
	fileLoc := "./part2_1_test_input.txt"
	answer, err := common.RunPart(Part2, fileLoc)
	if err != nil {
		fmt.Println(err)
		t.Errorf("There has been an error running Part2")
	} else {
		if answer != 1 {
			t.Errorf("The answer should be 1, got %d", answer)
		}
	}
}
