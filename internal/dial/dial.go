package dial

import (
	// "log"
	"strconv"
)

type Dial struct {
	currPosition  int
	steps         int
	startPosition int
	endPosition   int
}

func NewDial(steps int) *Dial {
	return &Dial{
		steps:         steps,
		startPosition: 0,
		endPosition:   steps - 1,
		currPosition:  steps / 2,
	}
}

func (d *Dial) Reset() {
	d.currPosition = d.steps / 2
}

func (d *Dial) Turn(direction string, steps int) int {
	// A turn over happens when we cross position 0
	turnOvers := 0
	startedAtZero := d.currPosition == 0

	if direction == "R" {
		// For right turns: count how many times we pass through 0
		// From position X, we cross 0 when we reach position >= steps
		stepsToZero := d.steps - d.currPosition
		if steps >= stepsToZero {
			turnOvers = 1 + (steps-stepsToZero)/d.steps
		}
		d.currPosition = (d.currPosition + steps) % d.steps
	} else {
		// For left turns: count how many times we pass through 0 going backwards
		// From position X, we cross 0 when we go past position 0
		if steps > d.currPosition {
			turnOvers = (steps - d.currPosition + d.steps - 1) / d.steps
		}
		newPosition := d.currPosition - steps
		d.currPosition = ((newPosition % d.steps) + d.steps) % d.steps
	}

	// Don't count it as a turnover if we land exactly on position 0
	// (that's already counted as password++)
	if d.currPosition == 0 && turnOvers > 0 {
		turnOvers--
	}

	// If we started at 0 going left, we don't count the immediate "crossing"
	if startedAtZero && direction == "L" && turnOvers > 0 {
		turnOvers--
	}

	return turnOvers
}

func (d *Dial) GetCurrentPosition() int {
	return d.currPosition
}

func (d *Dial) GetSimplePassword(directions []string) int {
	password := 0
	for _, direction := range directions {
		dir := string(direction[0])
		steps, _ := strconv.Atoi(direction[1:])
		d.Turn(dir, steps)
		if d.currPosition == 0 {
			password++
		}
	}
	return password
}
func (d *Dial) GetComplexPassword(directions []string) int {
	password := 0
	turnOvers := 0
	for _, direction := range directions {
		dir := string(direction[0])
		steps, _ := strconv.Atoi(direction[1:])
		turnOvers += d.Turn(dir, steps)
		if d.currPosition == 0 {
			password++
		}
	}
	return password + turnOvers
}
