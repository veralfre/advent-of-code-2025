package joltage_test

import (
	"testing"
	"veralfre.com/aoc-2025/internal/joltage"
)


// 987654321111111
// 811111111111119
// 234234234234278
// 818181911112111
func TestMaxJoltage(t *testing.T) {
	banks := [][]int{
		{9,8,7,6,5,4,3,2,1,1,1,1,1,1,1},
		{8,1,1,1,1,1,1,1,1,1,1,1,1,1,9},
		{2,3,4,2,3,4,2,3,4,2,3,4,2,7,8},
		{8,1,8,1,8,1,9,1,1,1,2,1,1,1,1},
	}
	results:= []int{98,89,78,92}
	for i, bank := range banks{
		stepResult := joltage.MaxJoltagePerBank(bank)
		if stepResult != results[i]{
			t.Errorf("Expected %d but got %d", results[i], stepResult)
		} 
	}
}

func TestTotalOutputJoltage(t *testing.T) {
	banks := [][]int{
		{9,8,7,6,5,4,3,2,1,1,1,1,1,1,1},
		{8,1,1,1,1,1,1,1,1,1,1,1,1,1,9},
		{2,3,4,2,3,4,2,3,4,2,3,4,2,7,8},
		{8,1,8,1,8,1,9,1,1,1,2,1,1,1,1},
	}
	banksObj := joltage.NewBanks(banks)
	expectedTotal := 357
	if banksObj.TotalOutputJoltage(2) != expectedTotal{
		t.Errorf("Expected total output joltage %d but got %d", expectedTotal, banksObj.TotalOutputJoltage(2))
	}
}