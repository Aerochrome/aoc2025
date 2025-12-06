package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	freshIndex, ids := getInput()
	fmt.Println(task1(freshIndex, ids))
	fmt.Println(task2(freshIndex))
}

func task1(rangeList []Range, ids []int) int {
	totalFreshIds := 0

	for _, id := range ids {
		for _, r := range rangeList {
			if r.IsInRange(id) {
				totalFreshIds++
				break
			}
		}
	}

	return totalFreshIds
}

func task2(rangeList []Range) int {
	optimizedRangeList := optimizeRanges(rangeList)

	totalFreshIds := 0
	for _, r := range optimizedRangeList {
		totalFreshIds += r.NumberOfIdsInRange()
	}

	return totalFreshIds
}

func optimizeRanges(rangeList []Range) []Range {
	optimizedRanges := make(map[Range]struct{})
	for _, r := range rangeList {
		optimizedRanges[r] = struct{}{}
	}

	keepLooping := true

	for keepLooping {
		cleanedRangeList := make(map[Range]struct{})

		// Delete ranges that fully contained by other ranges and also deduplicate
		for r := range optimizedRanges {
			isContainedByOther := false
			for other := range optimizedRanges {
				if other == r {
					continue
				}

				if other.Contains(r) {
					isContainedByOther = true
					break
				}
			}

			if !isContainedByOther {
				cleanedRangeList[r] = struct{}{}
			}
		}

		overlappingRanges := make(map[string][]Range)
		nonOverlappingRanges := make(map[Range]struct{})

		// Now separate ranges that overlap with other ranges
		for r := range cleanedRangeList {
			isNonOverlapping := true
			for other := range cleanedRangeList {
				if other == r {
					continue
				}
				if r.OverlapsWith(other) {
					overlapKey := buildOverlapKey(r, other)
					overlappingRanges[overlapKey] = []Range{r, other}
					isNonOverlapping = false
					break
				}
			}

			if isNonOverlapping {
				nonOverlappingRanges[r] = struct{}{}
			}
		}

		keepLooping = false

		newOptimizedRanges := make(map[Range]struct{})
		for r := range nonOverlappingRanges {
			_, ok := optimizedRanges[r]
			if !ok {
				keepLooping = true
			}

			newOptimizedRanges[r] = struct{}{}
		}

		for _, ranges := range overlappingRanges {
			r := ranges[0]
			other := ranges[1]
			mergedRange := r.Merge(other)

			_, ok := optimizedRanges[mergedRange]
			if !ok {
				keepLooping = true
			}

			newOptimizedRanges[mergedRange] = struct{}{}
		}

		optimizedRanges = newOptimizedRanges
	}

	optimizedRangeList := make([]Range, 0)
	for r := range optimizedRanges {
		optimizedRangeList = append(optimizedRangeList, r)
	}

	return optimizedRangeList
}

func buildOverlapKey(r Range, other Range) string {
	firstRange := r
	secondRange := other

	if r.Start > other.Start {
		firstRange = other
		secondRange = r
	}

	return fmt.Sprintf("%d-%d,%d-%d", firstRange.Start, firstRange.End, secondRange.Start, secondRange.End)
}

type Range struct {
	Start int
	End   int
}

func (r Range) IsInRange(id int) bool {
	return id >= r.Start && id <= r.End
}

func (r Range) NumberOfIdsInRange() int {
	return r.End - r.Start + 1
}

func (r Range) Contains(other Range) bool {
	return r.Start <= other.Start && r.End >= other.End
}

func (r Range) OverlapsWith(other Range) bool {
	if r.Contains(other) {
		return true
	}

	if other.Contains(r) {
		return true
	}

	return (r.Start >= other.Start && r.Start <= other.End) || (r.End >= other.Start && r.End <= other.End)
}

func (r Range) Merge(other Range) Range {
	if r.Contains(other) {
		return r
	}

	if other.Contains(r) {
		return other
	}

	if r == other {
		return r
	}

	if r.Start < other.Start {
		return Range{Start: r.Start, End: other.End}
	}

	return Range{Start: other.Start, End: r.End}
}

func (r Range) OverlappingIdCount(other Range) int {
	if !r.OverlapsWith(other) {
		return 0
	}

	if r.Contains(other) {
		return other.NumberOfIdsInRange()
	}

	if other.Contains(r) {
		return r.NumberOfIdsInRange()
	}

	if r.Start < other.Start {
		return r.End - other.Start + 1
	}

	return other.End - r.Start + 1
}

var rangeRegex = regexp.MustCompile(`([0-9]+)-([0-9]+)`)

func parseRange(in string) Range {
	matches := rangeRegex.FindStringSubmatch(in)
	if len(matches) != 3 {
		log.Fatal("Invalid range: ", in)
	}

	start, err := strconv.Atoi(matches[1])
	if err != nil {
		log.Fatal("Invalid start: ", matches[1])
	}

	end, err := strconv.Atoi(matches[2])
	if err != nil {
		log.Fatal("Invalid end: ", matches[2])
	}

	return Range{
		Start: start,
		End:   end,
	}
}

func getInput() ([]Range, []int) {
	filePath := "day5/input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rangeList := make([]Range, 0)
	ids := make([]int, 0)

	rangeMode := true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "" {
			rangeMode = false
			continue
		}

		if rangeMode {
			rangeList = append(rangeList, parseRange(line))
		} else {
			id, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal("Invalid id: ", line)
			}

			ids = append(ids, id)
		}
	}

	return rangeList, ids
}
