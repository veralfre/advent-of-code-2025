package main

import (
	"flag"
	"fmt"
	"os"

	"veralfre.com/aoc-2025/internal/joltage"
)

func main() {
	fileName := flag.String("filename", "./input.txt", "Combinations File")
	flag.Parse()

	data, err := os.ReadFile(*fileName)
	if err != nil {
		panic(err)
	}

	banks := joltage.NewBanksFromString(string(data))
	totalOutputJoltage := banks.TotalOutputJoltage(2)
	fmt.Printf("Total output joltage: %d\n", totalOutputJoltage)
	totalOutputJoltage = banks.TotalOutputJoltage(12)
	fmt.Printf("Total output joltage with 12 batteries: %d\n", totalOutputJoltage)

}
