package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type OperandFunc func(int, int) int

func loadOperands(line string) []OperandFunc {
	operandsParts := strings.Fields(line)
	operands := make([]OperandFunc, len(operandsParts))
	for i, ch := range operandsParts {
		switch ch {
		case "+":
			operands[i] = func(a, b int) int { return a + b }
		case "*":
			operands[i] = func(a, b int) int { return a * b }
		}
	}
	return operands
}

func solveAlgorithm(lines []string, operands []OperandFunc) []int {
	n := len(lines)
	// We want to get rid of all the spaces?

	firstLine := strings.Fields(lines[0])
	// fmt.Printf("First line: %d\n", len(firstLine))
	// fmt.Printf("Operands: %d\n", len(operands))
	row := make([]int, len(firstLine))
	for i, ch := range firstLine {
		val, err := strconv.Atoi(ch)
		if err != nil {
			panic(err)
		}
		row[i] = val
	}

	// Process each line
	for _, line := range lines[1 : n-1] {
		numbers := strings.Fields(line)
		for i, numStr := range numbers {
			val, err := strconv.Atoi(numStr)
			if err != nil {
				panic(err)
			}
			row[i] = operands[i](row[i], val)
		}
	}
	return row
}

func getTestLines() []string {
	return []string{
		"123 328  51 64 ",
		" 45 64  387 23 ",
		"  6 98  215 314",
		"*   +   *   +  ",
	}
}
func testAlgorithm() {
	// Simple test case
	// 123 328  51 64
	//  45 64  387 23
	//   6 98  215 314
	// *   +   *   +
	lines := getTestLines()
	operands := loadOperands(lines[3])
	result := solveAlgorithm(lines, operands)
	expected := []int{(123 * 45 * 6), (328 + 64 + 98), (51 * 387 * 215), (64 + 23 + 314)}
	fmt.Printf("Operands: %v\n", operands)
	fmt.Printf("Result: %v\n", result)
	for i, val := range result {
		if val != expected[i] {
			panic(fmt.Sprintf("Test failed at index %d: got %d, expected %d", i, val, expected[i]))
		}
	}
	fmt.Println("Test passed!")
}

func partTwo(lines []string) []int {
	output := []int{}
	accumulator := []int{}
	cols, rows := len(lines[0]), len(lines)
	chars := make([][]rune, rows)
	for i, line := range lines {
		chars[i] = []rune(line)
	}

	for col := cols - 1; col >= 0; col-- {
		colAcc := ""
		for row := 0; row < rows-1; row++ {
			char := chars[row][col]
			colAcc += string(char)
		}

		intVal, err := strconv.Atoi(strings.TrimSpace(colAcc))
		if err != nil {
			panic(err)
		}
		accumulator = append(accumulator, intVal)
		possibleOperand := chars[rows-1][col]
		switch possibleOperand {
		case '+', '*':
			var operand OperandFunc = nil
			switch possibleOperand {
			case '+':
				operand = func(a, b int) int { return a + b }
			case '*':
				operand = func(a, b int) int { return a * b }
			}
			// fmt.Printf("Found %c with accumulator: %v\n", possibleOperand, accumulator)
			problemSolution := accumulator[0]
			for _, val := range accumulator[1:] {
				problemSolution = operand(problemSolution, val)
			}
			output = append(output, problemSolution)
			accumulator = []int{}
			// We also skip one column as there will only be spaces
			col--
		default:
			//continue
		}

	}
	return output
}

func main() {
	fileName := flag.String("filename", "./input.txt", "Combinations File")
	flag.Parse()

	data, err := os.ReadFile(*fileName)
	if err != nil {
		panic(err)
	}

	// testAlgorithm()

	stringData := string(data)
	lines := strings.Split(stringData, "\n")
	n := len(lines)
	print("Number of columns:", n, "\n")

	operands := loadOperands(lines[n-1])
	// print("Operands loaded:", operands, "\n")
	row := solveAlgorithm(lines, operands)

	// Initialize row

	// Print final row
	// fmt.Println("Final row:", row)
	// Now we sum it all up
	total := 0
	for _, val := range row {
		total += val
	}
	fmt.Println("Total sum:", total)

	// For part two we want a matrix of characters:
	testResult:= partTwo(getTestLines())
	finalResult:= 0
	for _, val := range testResult{
		finalResult += val
	}
	fmt.Printf("Test output with final result %d\n", finalResult)
	
	finalResult= 0
	partTwoResult := partTwo(lines)
	for _, val := range partTwoResult{
		finalResult += val
	}
	fmt.Printf("Part two result: %d\n", finalResult)
	// fmt.Println("Character matrix:", chars)

}
