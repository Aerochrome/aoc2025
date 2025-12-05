package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	coordSystem := getInput()
	fmt.Println(task1(coordSystem))
	fmt.Println(task2(coordSystem))
}

func task1(coordSystem CoordSystem) int {
	accessiblePaperRolls := 0

	for coord, coordType := range coordSystem {
		if coordType != CoordTypePaperRoll {
			continue
		}

		if coordSystem.PaperRollsInVicinity(coord) > 3 {
			// Not accessible by forklift
			continue
		}

		accessiblePaperRolls++
	}

	return accessiblePaperRolls
}

func task2(coordSystem CoordSystem) int {
	accessiblePaperRolls := 0

	for {
		accessibleBefore := accessiblePaperRolls

		for coord, coordType := range coordSystem {
			if coordType != CoordTypePaperRoll {
				continue
			}

			if coordSystem.PaperRollsInVicinity(coord) > 3 {
				// Not accessible by forklift
				continue
			}

			accessiblePaperRolls++

			coordSystem[coord] = CoordTypeWalk
		}

		if accessibleBefore == accessiblePaperRolls {
			break
		}
	}

	return accessiblePaperRolls
}

type Coordinate struct {
	x int
	y int
}

type CoordSystem map[Coordinate]CoordType
type CoordType string

func CordTypeFromRune(input rune) CoordType {
	switch input {
	case '@':
		return CoordTypePaperRoll
	case '.':
		return CoordTypeWalk
	default:
		log.Fatal("Invalid coord type: ", input)
		return CoordType("")
	}
}

func (cs CoordSystem) PaperRollsInVicinity(coord Coordinate) int {
	paperRolls := 0

	coordsToCheck := []Coordinate{
		{x: coord.x - 1, y: coord.y - 1},
		{x: coord.x - 1, y: coord.y},
		{x: coord.x - 1, y: coord.y + 1},
		{x: coord.x, y: coord.y - 1},
		{x: coord.x, y: coord.y + 1},
		{x: coord.x + 1, y: coord.y - 1},
		{x: coord.x + 1, y: coord.y},
		{x: coord.x + 1, y: coord.y + 1},
	}

	for _, coordToCheck := range coordsToCheck {
		coordType, ok := cs[coordToCheck]
		if !ok {
			// not in system
			continue
		}

		if coordType == CoordTypePaperRoll {
			paperRolls++
		}
	}

	return paperRolls
}

const (
	CoordTypePaperRoll CoordType = "paper_roll"
	CoordTypeWalk      CoordType = "walkable"
)

func getInput() CoordSystem {
	filePath := "day4/input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	coordSystem := make(CoordSystem)
	y := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		// TODO
		for x, char := range line {
			coord := Coordinate{x: x, y: y}
			coordSystem[coord] = CordTypeFromRune(char)
		}

		y++
	}

	return coordSystem
}
