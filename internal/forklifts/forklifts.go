package forklifts

import (
	"fmt"
)

const AccessibleThreshold int = 4

type ForkliftGrid struct {
	Grid       [][]rune
	Width      int
	Height     int
	accessible [][]rune
}

func NewForkliftGrid(grid [][]rune) *ForkliftGrid {
	accessible := make([][]rune, len(grid))
	for i := range accessible {
		accessible[i] = make([]rune, len(grid[0]))
	}
	return &ForkliftGrid{
		Grid:       grid,
		Width:      len(grid[0]),
		Height:     len(grid),
		accessible: accessible,
	}
}

func (fg *ForkliftGrid) RemoveAccessibleRolls() int {
	// We will create a temporary grid to hold the updated state
	// this way we do not incur in spurious reads
	tmp := make([][]rune, fg.Height)
	for i := range tmp {
		tmp[i] = make([]rune, fg.Width)
	}
	fmt.Println("Sizes of tmp grid:", len(tmp), len(tmp[0]))
	removed := 0
	for y := 0; y < fg.Height; y++ {
		for x := 0; x < fg.Width; x++ {
			if fg.Grid[y][x] == '@' {
				neighbours := fg.getNeighbours(y, x)
				if neighbours < AccessibleThreshold {
					removed++
					tmp[y][x] = '.' // Mark as removed
				} else {
					tmp[y][x] = fg.Grid[y][x] // Keep original
				}
			} else {
				tmp[y][x] = fg.Grid[y][x] // Keep original
			}
		}
	}
	fg.Grid = tmp
	return removed
}

func (fg *ForkliftGrid) RemoveAllAccessibleRolls() int {
	totalRemoved:= 0 
	for {
		removed := fg.RemoveAccessibleRolls()
		if removed == 0 {
			break
		}
		totalRemoved += removed
	}
	return totalRemoved
}


func (fg *ForkliftGrid) PrintAccessible() {
	for y := 0; y < fg.Height; y++ {
		print("|")
		for x := 0; x < fg.Width; x++ {
			print(string(fg.accessible[y][x]))
		}
		print("|")
		println()
	}
}

func (fg *ForkliftGrid) PrintGrid() {
	for y := 0; y < fg.Height; y++ {
		print("|")
		for x := 0; x < fg.Width; x++ {
			print(string(fg.Grid[y][x]))
		}
		print("|")
		println()
	}
}

func (fg *ForkliftGrid) CountAccessibile(neighbourThreshold int) int {
	count := 0
	for y := 0; y < fg.Height; y++ {
		for x := 0; x < fg.Width; x++ {
			fg.accessible[y][x] = fg.Grid[y][x]
			if fg.Grid[y][x] == '@' {
				neighbours := fg.getNeighbours(y, x)
				if neighbours < neighbourThreshold {
					fg.accessible[y][x] = 'x'
					count++
				}
			}
		}
	}
	return count
}

//  internal

func (fg *ForkliftGrid) getNeighbours(col, row int) int {
	neighbours := 0
	directions := [][2]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1},
	}
	for _, dir := range directions {
		nx := row + dir[0]
		ny := col + dir[1]
		if nx >= 0 && nx < fg.Width && ny >= 0 && ny < fg.Height {
			if fg.Grid[ny][nx] == '@' {
				neighbours++
			}
		}
	}
	return neighbours
}
