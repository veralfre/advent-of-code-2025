package joltage

import (
	"strings"
	// "log"
)
// type Battery struct {
// 	joltage int 
// }

// type Bank struct {
// 	joltages []int
// }

// type Banks struct{
// 	banks []Bank
// }

type Banks struct {
	batteries [][]int
}



func NewBanks(batteries [][]int) *Banks {
	return &Banks{
		batteries: batteries,
	}
}

func NewBanksFromString(input string) *Banks{
	lines := strings.Split(input,"\n")
	batteries := [][]int{}
	for _, line := range lines{
		bankStrs := strings.Split(line,"")
		bank := []int{}
		for _, bStr := range bankStrs{
			joltage := int(bStr[0] - '0')
			bank = append(bank, joltage)
		}
		batteries = append(batteries, bank)
	}
	// log.Printf("Battery size: %dx%d", len(batteries), len(batteries[0]))
	return &Banks{
		batteries: batteries,
	}
}

func MaxJoltagePerBank(bank []int) int{
	return MaxJoltagePerBankWithBatteryNum(2, bank)
}

func MaxJoltagePerBankWithBatteryNum(batteryNum int, bank []int) int {
	max := 0 
	n := len(bank)
	stack := []int{bank[0]}
	for i, battery := range bank[1:]{
		for len(stack)>0 && stack[len(stack)-1] < battery && len(stack)-1 + (n-i)> batteryNum{
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, battery)
		if len(stack) > batteryNum{
			stack = stack[:batteryNum]
		}
		// log.Printf("[%d] Stack is %v", battery, stack)
	}
	for _, val := range stack{
		max = max * 10 + val
	}
	return max
}



func (banks *Banks) TotalOutputJoltage(batteryNum int) int{
	total := 0
	for _, bank:= range banks.batteries{
		stepResult := MaxJoltagePerBankWithBatteryNum(batteryNum, bank)
		// log.Printf("Bank %v produced joltage %d", bank, stepResult)
		total += stepResult
	}
	return total
}
