package idranges_test

import (
	"testing"

	"veralfre.com/aoc-2025/internal/idranges"
)

func TestInvalidIds(t *testing.T) {
	idRange := idranges.NewIdRange(11, 22)
	invalidIds := idRange.InvalidIdsPart1()
	expectedInvalidIds := []int{11, 22}
	if len(invalidIds) != len(expectedInvalidIds) {
		t.Errorf("Expected %d invalid ids, got %d", len(expectedInvalidIds), len(invalidIds))
		t.Errorf("Invalid ids: %v", invalidIds)
	}

	idRange = idranges.NewIdRange(95, 115)
	invalidIds = idRange.InvalidIdsPart1()
	expectedInvalidIds = []int{99}
	if len(invalidIds) != len(expectedInvalidIds) {
		t.Errorf("Expected %d invalid ids, got %d", len(expectedInvalidIds), len(invalidIds))
		t.Errorf("Invalid ids: %v", invalidIds)
	}
}

func TestProvidedRanges(t *testing.T) {
	// 11-22 has two invalid IDs, 11 and 22.
	// 95-115 has one invalid ID, 99.
	// 998-1012 has one invalid ID, 1010.
	// 1188511880-1188511890 has one invalid ID, 1188511885.
	// 222220-222224 has one invalid ID, 222222.
	// 1698522-1698528 contains no invalid IDs.
	// 446443-446449 has one invalid ID, 446446.
	// 38593856-38593862 has one invalid ID, 38593859.

	testCases := []struct {
		start              int
		end                int
		expectedInvalidIds []int
	}{
		{11, 22, []int{11, 22}},
		{95, 115, []int{99}},
		{998, 1012, []int{1010}},
		{1188511880, 1188511890, []int{1188511885}},
		{222220, 222224, []int{222222}},
		{1698522, 1698528, []int{}},
		{446443, 446449, []int{446446}},
		{38593856, 38593862, []int{38593859}},
	}

	for _, tc := range testCases {
		idRange := idranges.NewIdRange(tc.start, tc.end)
		invalidIds := idRange.InvalidIdsPart1()
		if len(invalidIds) != len(tc.expectedInvalidIds) {
			t.Errorf("For range %d-%d, expected %d invalid ids, got %d", tc.start, tc.end, len(tc.expectedInvalidIds), len(invalidIds))
			t.Errorf("Invalid ids: %v", invalidIds)
		}
	}
}

func TestComputeFromString(t *testing.T) {
	input := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	result := 1227775554

	idranges := idranges.FromString(input)
	totalInvalidIds := 0
	for _, idrange := range idranges {
		invalidIds := idrange.InvalidIdsPart1()
		for _, invalidId := range invalidIds {
			t.Logf("Invalid ID: %d", invalidId)
			totalInvalidIds += invalidId
		}
	}

	if totalInvalidIds != result {
		t.Errorf("Expected total invalid IDs sum %d, got %d", result, totalInvalidIds)
	}

}

func TestInvalidIdsPart2(t *testing.T) {
	idRange := idranges.NewIdRange(11, 22)
	invalidIds := idRange.InvalidIdsPart2()
	expectedInvalidIds := []int{11, 22}
	if len(invalidIds) != len(expectedInvalidIds) {
		t.Errorf("Expected %d invalid ids, got %d", len(expectedInvalidIds), len(invalidIds))
		t.Errorf("Invalid ids: %v", invalidIds)
	}

	idRange = idranges.NewIdRange(95, 115)
	invalidIds = idRange.InvalidIdsPart2()
	expectedInvalidIds = []int{99, 111}
	if len(invalidIds) != len(expectedInvalidIds) {
		t.Errorf("Expected %d invalid ids, got %d", len(expectedInvalidIds), len(invalidIds))
		t.Errorf("Invalid ids: %v", invalidIds)
	}
}

func TestProvidedRangesPart2(t *testing.T) {
	// Part2 catches more invalid IDs than Part1
	// 11-22 has two invalid IDs, 11 and 22.
	// 95-115 has two invalid IDs, 99 and 111.
	// 998-1012 has two invalid IDs, 999 and 1010.
	// 1188511880-1188511890 has one invalid ID, 1188511885.
	// 222220-222224 has one invalid ID, 222222.
	// 1698522-1698528 contains no invalid IDs.
	// 446443-446449 has one invalid ID, 446446.
	// 38593856-38593862 has one invalid ID, 38593859.
	// 565653-565659 has one invalid ID, 565656.
	// 824824821-824824827 has one invalid ID, 824824824.
	// 2121212118-2121212124 has one invalid ID, 2121212121.

	testCases := []struct {
		start              int
		end                int
		expectedInvalidIds []int
	}{
		{11, 22, []int{11, 22}},
		{95, 115, []int{99, 111}},
		{998, 1012, []int{999, 1010}},
		{1188511880, 1188511890, []int{1188511885}},
		{222220, 222224, []int{222222}},
		{1698522, 1698528, []int{}},
		{446443, 446449, []int{446446}},
		{38593856, 38593862, []int{38593859}},
		{565653, 565659, []int{565656}},
		{824824821, 824824827, []int{824824824}},
		{2121212118, 2121212124, []int{2121212121}},
	}

	for _, tc := range testCases {
		idRange := idranges.NewIdRange(tc.start, tc.end)
		invalidIds := idRange.InvalidIdsPart2()
		if len(invalidIds) != len(tc.expectedInvalidIds) {
			t.Errorf("For range %d-%d, expected %d invalid ids, got %d", tc.start, tc.end, len(tc.expectedInvalidIds), len(invalidIds))
			t.Errorf("Invalid ids: %v", invalidIds)
		}
	}
}

func TestComputeFromStringPart2(t *testing.T) {
	input := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	result := 4174379265

	idranges := idranges.FromString(input)
	totalInvalidIds := 0
	for _, idrange := range idranges {
		invalidIds := idrange.InvalidIdsPart2()
		for _, invalidId := range invalidIds {
			t.Logf("Invalid ID: %d", invalidId)
			totalInvalidIds += invalidId
		}
	}

	if totalInvalidIds != result {
		t.Errorf("Expected total invalid IDs sum %d, got %d", result, totalInvalidIds)
	}

}
