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
	rotations := getInput()

	fmt.Println(task1(rotations))
	fmt.Println(task2(rotations))
}

func task1(rotations []Rotation) int {

	currentPosition := 50
	timesAtZero := 0

	for _, rotation := range rotations {
		newPosition, _ := getNewPosition(currentPosition, rotation.Direction, rotation.Steps)

		if newPosition == 0 {
			timesAtZero++
		}

		currentPosition = newPosition
	}

	return timesAtZero
}

func task2(rotations []Rotation) int {

	currentPosition := 50
	timesAtZero := 0

	for _, rotation := range rotations {
		newPosition, timesZeroDuring := getNewPosition(currentPosition, rotation.Direction, rotation.Steps)

		if newPosition == 0 {
			timesAtZero++
		}

		timesAtZero += timesZeroDuring
		currentPosition = newPosition
	}

	return timesAtZero
}

func getNewPosition(currentPosition int, direction Direction, steps int) (int, int) {
	timesZeroDuring := 0

	for i := range steps {
		newPosition := takeOneStep(currentPosition, direction)

		if newPosition == 0 {
			if i+1 < steps {
				// only if not end pos
				timesZeroDuring++
			}
		}

		if newPosition < 0 {
			currentPosition = 99
			continue
		}

		if newPosition > 99 {
			currentPosition = 0
			if i+1 < steps {
				// only if not end pos
				timesZeroDuring++
			}
			continue
		}

		currentPosition = newPosition
	}

	return currentPosition, timesZeroDuring
}

func takeOneStep(currentPosition int, direction Direction) int {
	switch direction {
	case DirectionLeft:
		return currentPosition - 1
	default:
		return currentPosition + 1
	}
}

type Direction string

const (
	DirectionLeft  Direction = "L"
	DirectionRight Direction = "R"
)

type Rotation struct {
	Direction Direction
	Steps     int
}

func getInput() []Rotation {
	filePath := "day1/input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	re := regexp.MustCompile(`([A-Z])([0-9]+)`)

	result := make([]Rotation, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		// TODO
		matches := re.FindStringSubmatch(line)
		if len(matches) != 3 {
			log.Fatal("Invalid line: ", line)
		}

		direction := matches[1]
		steps, err := strconv.Atoi(matches[2])
		if err != nil {
			log.Fatal("Invalid steps: ", matches[2])
		}

		result = append(result, Rotation{
			Direction: Direction(direction),
			Steps:     steps,
		})
	}

	return result
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
