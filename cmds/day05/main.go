package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"veralfre.com/aoc-2025/internal/idranges"
)

func main() {
	fileName := flag.String("filename", "./input.txt", "Combinations File")
	flag.Parse()

	data, err := os.ReadFile(*fileName)
	if err != nil {
		panic(err)
	}

	ranges := []*idranges.IdRange{}
	ingredients := []int{}

	mode := 0 // 0 for ranges, 1 for ingredients
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			mode = 1
			continue
		}
		if mode == 0 {
			// RANGES PARSING
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				panic("Invalid range line: " + line)
			}
			start, err1 := strconv.Atoi(parts[0])
			end, err2 := strconv.Atoi(parts[1])
			if err1 != nil || err2 != nil {
				panic("Invalid range numbers: " + line)
			}

			// Well for part two we need to merge them first ..
			newRange := idranges.NewIdRange(start, end)
			for {
				mergedSomething := false
				for i := 0; i < len(ranges); i++ {
					if ranges[i].CanMerge(newRange) {
						newRange = ranges[i].Merge(newRange)

						// remove the old one (no duplicates)
						ranges = append(ranges[:i], ranges[i+1:]...)
						i-- // because we modified the slice
						mergedSomething = true
					}
				}

				if !mergedSomething {
					break
				}
			}

			ranges = append(ranges, newRange)
		} else {

			// INGREDIENTS PARSING
			id, err := strconv.Atoi(line)
			if err != nil {
				panic("Invalid ingredient id: " + line)
			}
			ingredients = append(ingredients, id)
		}
	}

	// PART 1

	// let's sort the ranges
	sort.Slice(ranges, func(i, j int) bool {
		// We sort by start and if they match by end
		if ranges[i].Start() == ranges[j].Start() {
			return ranges[i].End() < ranges[j].End()
		} else {
			return ranges[i].Start() < ranges[j].Start()
		}
	})

	totalFresh := 0
	for _, id := range ingredients {
		isValid := false
		for _, r := range ranges {
			if r.IsInRange(id) {
				isValid = true
				break
			}
		}
		if isValid {
			totalFresh++
		}
	}

	fmt.Println("Total fresh ingredients:", totalFresh)
	// PART 2 , one could have simply used an hashmap instead of ranges but okay..
	totalAllowedIngredients := 0
	for _, r := range ranges {
		// fmt.Println("Range from", r.Start(), "to", r.End())
		totalAllowedIngredients += r.End() - r.Start() + 1
	}
	fmt.Println("Total allowed ingredients in ranges:", totalAllowedIngredients)

}
