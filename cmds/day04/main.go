package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"veralfre.com/aoc-2025/internal/forklifts"
)

func main() {
	fileName := flag.String("filename", "./input.txt", "Combinations File")
	flag.Parse()

	data, err := os.ReadFile(*fileName)
	if err != nil {
		panic(err)
	}

	var stringGrid string = string(data)
	var lines []string = strings.Split(stringGrid, "\n")
	var grid [][] rune = make([][]rune, len(lines))
	for i, line := range lines{
		grid[i] = []rune(line)
	}
	forkliftGrid := forklifts.NewForkliftGrid(grid)


	accessibleRolls := forkliftGrid.CountAccessibile(4)
	fmt.Printf("Total accessible rolls: %d\n", accessibleRolls)
	totalRemoved := forkliftGrid.RemoveAllAccessibleRolls()
	fmt.Printf("Total removed rolls after all iterations: %d\n", totalRemoved)
	// totalOutputJoltage = banks.TotalOutputJoltage(12)
	// fmt.Printf("Total output joltage with 12 batteries: %d\n", totalOutputJoltage)

}
