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
	batteryBanks := getInput()
	fmt.Println(task1(batteryBanks))
	fmt.Println(task2(batteryBanks))
}

func task1(batteryBanks []BatteryBank) int {
	joltageSum := 0
	for _, batteryBank := range batteryBanks {
		joltageSum += batteryBank.FindMaximumJoltage()
	}
	return joltageSum
}

func task2(batteryBanks []BatteryBank) int {
	joltageSum := 0

	for _, batteryBank := range batteryBanks {
		joltageSum += batteryBank.FindMaximumJoltageVariable(12)
	}

	return joltageSum
}

type BatteryBank []Battery
type Battery int

func (bb BatteryBank) FindMaximumJoltage() int {
	highestTenth := Battery(0)
	highestJoltage := Battery(0)

	for idx, battery := range bb {
		if battery <= highestTenth {
			// If we tried for example every combination with 2, 1 as a tenth will be defo lower
			continue
		}

		highestTenth = battery

		for _, otherBattery := range bb[idx+1:] {
			joltage := (battery * 10) + otherBattery
			if joltage > highestJoltage {
				highestJoltage = joltage
			}
		}
	}

	return int(highestJoltage)
}

func (bb BatteryBank) FindMaximumJoltageVariable(digitLen int) int {
	startIdx := 0

	resultDigits := ""

	for i := range digitLen {
		suffixLen := digitLen - (i + 1)
		highestNumberIdx := bb.FindMaximumNumberPosition(startIdx, suffixLen)
		resultDigits += strconv.Itoa(int(bb[startIdx+highestNumberIdx]))

		startIdx += highestNumberIdx + 1
	}

	result, err := strconv.Atoi(resultDigits)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func (bb BatteryBank) FindMaximumNumberPosition(startIdx int, suffixLen int) int {
	endIdx := len(bb) - suffixLen

	highestNumber := Battery(0)
	highestNumberIdx := 0

	for idx, battery := range bb[startIdx:endIdx] {
		if battery > highestNumber {
			highestNumber = battery
			highestNumberIdx = idx
		}
	}

	return highestNumberIdx
}

func getInput() []BatteryBank {
	filePath := "day3/input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := make([]BatteryBank, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		batteryBank := make(BatteryBank, 0, len(line))
		for _, battery := range line {
			batteryInt, err := strconv.Atoi(string(battery))
			if err != nil {
				log.Fatal(err)
			}

			batteryBank = append(batteryBank, Battery(batteryInt))
		}

		result = append(result, batteryBank)
	}

	return result
}
