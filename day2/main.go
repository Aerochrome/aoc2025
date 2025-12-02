package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	stepRanges := getInput()

	fmt.Println(task1(stepRanges))
	fmt.Println(task2(stepRanges))
}

func task1(stepRanges []StepRange) int {
	invalidIDSum := 0

	for _, stepRange := range stepRanges {
		for id := range stepRange.All {
			if isValidID(id) {
				continue
			}

			invalidIDSum += id
		}
	}

	return invalidIDSum
}

func task2(stepRanges []StepRange) int {
	validIDSum := 0

	for _, stepRange := range stepRanges {
		for id := range stepRange.All {
			if findStringRepetition(strconv.Itoa(id)) {
				validIDSum += id
			}
		}
	}

	return validIDSum
}

type StepRange struct {
	Start int
	End   int
}

func NewStepRangeFromInput(input string) StepRange {
	input = strings.TrimSpace(input)
	if input == "" {
		log.Fatal("Invalid input: ", input)
	}

	parts := strings.Split(input, "-")
	if len(parts) != 2 {
		log.Fatal("Invalid input: ", input)
	}

	start, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatal("Invalid start: ", parts[0])
	}

	end, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal("Invalid end: ", parts[1])
	}

	if start > end {
		log.Fatal("Start is greater than end: ", start, " > ", end)
	}

	return StepRange{
		Start: start,
		End:   end,
	}
}

func (r *StepRange) All(yield func(int) bool) {
	for i := r.Start; i <= r.End; i++ {
		if !yield(i) {
			return
		}
	}
}

func isValidID(id int) bool {
	idStr := strconv.Itoa(id)
	if len(idStr)%2 != 0 {
		return true
	}

	firstHalf := idStr[:len(idStr)/2]
	secondHalf := idStr[len(idStr)/2:]

	return firstHalf != secondHalf
}

func findStringRepetition(input string) bool {
	startAt := len(input) / 2

	for i := startAt; i > 0; i-- {
		if len(input)%i != 0 {
			// not divisible, so no repetition possible
			continue
		}

		prevBuf := ""
		currBuf := ""

		repetitionFound := true

		for _, char := range input {
			currBuf += string(char)

			if len(currBuf) == i {
				if prevBuf != "" && prevBuf != currBuf {
					repetitionFound = false
					break
				}

				prevBuf = currBuf
				currBuf = ""
			}
		}

		if repetitionFound {
			return true
		}
	}

	return false
}

func getInput() []StepRange {
	filePath := "day2/input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	stepRanges := make([]StepRange, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		for part := range strings.SplitSeq(line, ",") {
			stepRanges = append(stepRanges, NewStepRangeFromInput(part))
		}
	}

	return stepRanges
}
