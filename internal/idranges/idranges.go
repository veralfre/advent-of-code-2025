package idranges

import (
	"strconv"
	"strings"
	// "log"
)

type IdRange struct {
	start int
	end   int
}

func NewIdRange(start, end int) *IdRange {
	return &IdRange{
		start: start,
		end:   end,
	}
}

func (r *IdRange) Start() int {
	return r.start
}

func (r *IdRange) End() int {
	return r.end
}

func (r *IdRange) IsInRange(id int) bool {
	return id >= r.start && id <= r.end
}

func (r *IdRange) CanMerge(other *IdRange) bool {
	// two ranges can merge if they overlap or touch
	return r.start <= other.end+1 && other.start <= r.end+1
}

func (r *IdRange) Merge(other *IdRange) *IdRange {
	newStart := r.start
	if other.start < newStart {
		newStart = other.start
	}
	newEnd := r.end
	if other.end > newEnd {
		newEnd = other.end
	}
	return NewIdRange(newStart, newEnd)
}
// for all ids in the range, if there is any id with duplicate digits, return false
// 55 is not valid
// 99 is not valid
// 6464 is not valid
// 123123 is not valid
func (r *IdRange) InvalidIdsPart1() []int {
	invalidIds := []int{}
	for id := r.start; id <= r.end; id++ {
		idStr := strconv.Itoa(id)
		idLen := len(idStr)
		for i := 1; i < idLen; i++ {
			substr := string(idStr[0:i])
			for len(substr) < idLen {
				substr += substr
			}
			if substr == idStr {
				invalidIds = append(invalidIds, id)
				break
			}
		}

	}
	return invalidIds
}

func (r *IdRange) InvalidIdsPart2() []int {
	invalidIds := []int{}
	for id := r.start; id <= r.end; id++ {
		idStr := strconv.Itoa(id)
		idLen := len(idStr)
		for i := 1; i < idLen; i++ {
			substr := string(idStr[0:i])
			acc := ""
			for len(acc) < idLen {
				acc += substr
			}
			if acc == idStr {
				invalidIds = append(invalidIds, id)
				break
			}
		}

	}
	return invalidIds
}

func FromString(rangeStr string) []IdRange {
	// we must split by ,
	// then by -
	// split by , in order to get the ranges
	outputRanges := []IdRange{}
	ranges := strings.Split(rangeStr, ",")
	for _, singleRange := range ranges {
		parts := strings.Split(singleRange, "-")
		lb, _ := strconv.Atoi(parts[0])
		ub, _ := strconv.Atoi(parts[1])
		outputRanges = append(outputRanges, IdRange{start: lb, end: ub})
	}

	return outputRanges
}

//  External API 

func ComputeTotalInvalidIdsPart1(ranges []IdRange) int {
	totalInvalidIds := 0
	for _, idRange := range ranges {
		invalidIds := idRange.InvalidIdsPart1()
		for _, invalidId := range invalidIds {
			totalInvalidIds += invalidId
		}
	}
	return totalInvalidIds
}

func ComputeTotalInvalidIdsPart2(ranges []IdRange) int {
	totalInvalidIds := 0
	for _, idRange := range ranges {
		invalidIds := idRange.InvalidIdsPart2()
		for _, invalidId := range invalidIds {
			totalInvalidIds += invalidId
		}
	}
	return totalInvalidIds
}

