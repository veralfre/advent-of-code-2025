package forklifts_test

import (
	"testing"
	"fmt"
	"veralfre.com/aoc-2025/internal/forklifts"
)


func TestRollsOfPaper(t *testing.T) {
	rollsMap := [][]rune{
		[]rune("..@@.@@@@."),
		[]rune("@@@.@.@.@@"),
		[]rune("@@@@@.@.@@"),
		[]rune("@.@@@@..@."),
		[]rune("@@.@@@@.@@"),
		[]rune(".@@@@@@@.@"),
		[]rune(".@.@.@.@@@"),
		[]rune("@.@@@.@@@@"),
		[]rune(".@@@@@@@@."),
		[]rune("@.@.@@@.@."),
	}
	result := 13
	grid := forklifts.NewForkliftGrid(rollsMap)
	stepResult := grid.CountAccessibile(4)
	if stepResult != result {
		t.Errorf("Expected %d but got %d", result, stepResult)
	}
	fmt.Println("Printing the grid:")
	grid.PrintAccessible()
}


func TestRemoveAccessibleRolls(t *testing.T) {
	rollsMap := [][]rune{
		[]rune("..@@.@@@@."),
		[]rune("@@@.@.@.@@"),
		[]rune("@@@@@.@.@@"),
		[]rune("@.@@@@..@."),
		[]rune("@@.@@@@.@@"),
		[]rune(".@@@@@@@.@"),
		[]rune(".@.@.@.@@@"),
		[]rune("@.@@@.@@@@"),
		[]rune(".@@@@@@@@."),
		[]rune("@.@.@@@.@."),
	}
	// Sum on results must be 43
	results:= []int{13,12,7,5,2,1,1,1,1}
	grid := forklifts.NewForkliftGrid(rollsMap)
	totalRemoved := 0
	removed := 0
	i:= 0
	for {
		removed = grid.RemoveAccessibleRolls()
		if removed == 0 {
			break
		}
		grid.PrintGrid()
		fmt.Println()
		totalRemoved += removed
		expectedRemoved := results[i]
		if removed != expectedRemoved {
			t.Errorf("Step %d: Expected to remove %d but got %d", i, expectedRemoved, removed)
		}
		i++
	}
	
	

	expectedTotal := 43
	if totalRemoved != expectedTotal {
		t.Errorf("Expected total removed %d but got %d", expectedTotal, totalRemoved)
	}

}

func TestRemoveAllAccessibleRolls(t *testing.T) {
	rollsMap := [][]rune{
		[]rune("..@@.@@@@."),
		[]rune("@@@.@.@.@@"),
		[]rune("@@@@@.@.@@"),
		[]rune("@.@@@@..@."),
		[]rune("@@.@@@@.@@"),
		[]rune(".@@@@@@@.@"),
		[]rune(".@.@.@.@@@"),
		[]rune("@.@@@.@@@@"),
		[]rune(".@@@@@@@@."),
		[]rune("@.@.@@@.@."),
	}
	// Sum on results must be 43
	grid := forklifts.NewForkliftGrid(rollsMap)
	totalRemoved := grid.RemoveAllAccessibleRolls()
	expectedTotal := 43
	if totalRemoved != expectedTotal {
		t.Errorf("Expected total removed %d but got %d", expectedTotal, totalRemoved)
	}

}