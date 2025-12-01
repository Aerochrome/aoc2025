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
}

func task1(rotations []Rotation) int {

	currentPosition := 50
	timesAtZero := 0

	for _, rotation := range rotations {
		var newPosition int

		switch rotation.Direction {
		case DirectionLeft:
			newPosition = currentPosition - rotation.Steps
		case DirectionRight:
			newPosition = currentPosition + rotation.Steps
		}

		currentPosition = newPosition % 100

		if currentPosition == 0 {
			timesAtZero++
		}

	}

	return timesAtZero
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
	filePath := "day1/example.txt"

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
