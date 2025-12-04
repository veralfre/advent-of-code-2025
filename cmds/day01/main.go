package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"veralfre.com/aoc-2025/internal/dial"
)

const maxPositions int = 100

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

	combinations := []string{}
	for scanner.Scan() {
		combination := scanner.Text()
		if err != nil {
			panic(err)
		}
		combinations = append(combinations, combination)
	}
	dial := dial.NewDial(maxPositions)

	password := dial.GetSimplePassword(combinations)
	fmt.Printf("Simple password %d\n", password)

	dial.Reset()
	password = dial.GetComplexPassword(combinations)
	fmt.Printf("Complex password %d\n", password)

}
