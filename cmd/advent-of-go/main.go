package main

import (
	"flag"
	"fmt"
	"github.com/jgrove2/advent-of-go/internal/day1"
	"github.com/jgrove2/advent-of-go/internal/day2"
	"github.com/jgrove2/advent-of-go/internal/day3"
	"github.com/jgrove2/advent-of-go/internal/day4"
	"strconv"
)

func main() {
	fmt.Println("Running Advent of code")
	dfunc := []func(){day1.Day1, day2.Day2, day3.Day3, day4.Day4}
	flag.Parse()

	if len(flag.Args()) <= 0 {
		for i := 0; i < len(dfunc); i++ {
			dfunc[i]()
		}
	}
	for _, days := range flag.Args() {
		arg, _ := strconv.Atoi(days)
		if arg > 1 || arg < len(dfunc) {
			dfunc[arg-1]()
		}
	}
}
