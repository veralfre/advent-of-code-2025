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

// for all ids in the range, if there is any id with duplicate digits, return false
// 55 is not valid
// 99 is not valid
// 6464 is not valid
// 123123 is not valid
func (r *IdRange) InvalidIds() []int {
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

func ComputeTotalInvalidIds(ranges []IdRange) int {
	totalInvalidIds := 0
	for _, idRange := range ranges {
		invalidIds := idRange.InvalidIds()
		for _, invalidId := range invalidIds {
			totalInvalidIds += invalidId
		}
	}
	return totalInvalidIds
}
