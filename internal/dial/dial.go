package dial

import (
	"strconv"
)

type Dial struct {
	rawPosition   int // Keep track of raw position (can be negative or > steps)
	steps         int
	startPosition int
	endPosition   int
}

func NewDial(steps int) *Dial {
	return &Dial{
		steps:         steps,
		startPosition: 0,
		endPosition:   steps - 1,
		rawPosition:   steps / 2,
	}
}

func (d *Dial) Reset() {
	d.rawPosition = d.steps / 2
}

func (d *Dial) Turn(direction string, steps int) int {
	// Track turnovers by comparing hundreds digits (position / d.steps)
	prevPosition := d.rawPosition

	if direction == "R" {
		d.rawPosition += steps
	} else {
		d.rawPosition -= steps
	}

	// Count how many times we crossed position 0 by comparing hundreds digits
	// Use floor division to match Python's // operator
	prevHundreds := floorDiv(prevPosition, d.steps)
	currHundreds := floorDiv(d.rawPosition, d.steps)
	turnOvers := abs(currHundreds - prevHundreds)

	// Special handling for left turns
	if direction == "L" {
		// Add 1 if we land exactly on a multiple of d.steps
		if d.rawPosition%d.steps == 0 {
			turnOvers++
		}
		// Subtract 1 if we started on a multiple of d.steps
		if prevPosition%d.steps == 0 {
			turnOvers--
		}
	}

	return turnOvers
}

func floorDiv(a, b int) int {
	// Python-style floor division
	if (a < 0) != (b < 0) && a%b != 0 {
		return a/b - 1
	}
	return a / b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (d *Dial) GetCurrentPosition() int {
	// Return the modulo position for display purposes
	pos := d.rawPosition % d.steps
	if pos < 0 {
		pos += d.steps
	}
	return pos
}

func (d *Dial) GetSimplePassword(directions []string) int {
	password := 0
	for _, direction := range directions {
		dir := string(direction[0])
		steps, _ := strconv.Atoi(direction[1:])
		d.Turn(dir, steps)
		if d.rawPosition%d.steps == 0 {
			password++
		}
	}
	return password
}

func (d *Dial) GetComplexPassword(directions []string) int {
	totalScore := 0
	for _, direction := range directions {
		dir := string(direction[0])
		steps, _ := strconv.Atoi(direction[1:])
		totalScore += d.Turn(dir, steps)
	}
	return totalScore
}
