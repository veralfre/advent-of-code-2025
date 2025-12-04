package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"veralfre.com/aoc-2025/internal/idranges"
)

func main() {
	fileName := flag.String("filename", "./input.txt", "Combinations File")
	flag.Parse()

	var file *os.File
	var err error
	file, err = os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// We read the long line
	scanner.Scan()
	rangeStr := scanner.Text()
	idRanges := idranges.FromString(rangeStr)
	totalInvalidIds := idranges.ComputeTotalInvalidIds(idRanges)
	fmt.Printf("Sum of invalid IDs: %d\n", totalInvalidIds)

}
