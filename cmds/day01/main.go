package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	//"log"
	//"os"
)

const maxPositions int = 100

func next(position int, direction string, clicks int) (int, int) {
	crossZeroes := 0
	if direction == "L" {
		if position-clicks < 0 {
			crossZeroes = (clicks-position-1)/maxPositions + 1
		}
		position = (position - clicks) % maxPositions
		if position < 0 {
			position += maxPositions
		}
	} else {
		if position+clicks >= maxPositions {
			crossZeroes = (position + clicks) / maxPositions
		}

		position = (position + clicks) % maxPositions
	}
	return position, crossZeroes
}

func main() {
	fileName := flag.String("filename", "./combinations", "Combinations File")
	testRun := flag.Bool("test", false, "Whether to use initial combination")
	flag.Parse()

	var file *os.File
	var err error
	if *testRun {
		file = nil
	} else {
		file, err = os.Open(*fileName)
		if err != nil {
			panic(err)
		}

	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var password, crossZeroes, position int = 0, 0, 50
	combinations := []string{}
	if !(*testRun) {
		for scanner.Scan() {
			combination := scanner.Text()
			if err != nil {
				panic(err)
			}
			combinations = append(combinations, combination)
		}

	} else {
		combinations = []string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"}
	}

	for _, combination := range combinations {
		direction := string(combination[0])
		clicks, err := strconv.Atoi(combination[1:])
		if err != nil {
			panic(err)
		}
		cross := 0
		position, cross = next(position, direction, clicks)

		if position == 0 {
			password++
		}
		crossZeroes += cross
	}

	fmt.Println("Password is:", password+crossZeroes)

}
