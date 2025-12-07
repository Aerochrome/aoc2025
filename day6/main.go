package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	OperatorAdd      OperatorType = "+"
	OperatorMultiply OperatorType = "*"
)

func main() {
	operations := getInput()
	fmt.Println(task1(operations))
}

func task1(operations []Operation) int {
	result := 0
	for _, operation := range operations {
		result += operation.Calculate()
	}
	return result
}

type OperatorType string

func (o OperatorType) validate() error {
	switch o {
	case OperatorAdd:
		return nil
	case OperatorMultiply:
		return nil
	default:
		return fmt.Errorf("invalid operator: %s", o)
	}
}

type Operation struct {
	Operator OperatorType
	Operands []int
}

func (o Operation) Calculate() int {
	switch o.Operator {
	case OperatorAdd:
		result := 0
		for _, operand := range o.Operands {
			result += operand
		}
		return result
	case OperatorMultiply:
		result := 1
		for _, operand := range o.Operands {
			result *= operand
		}
		return result
	default:
		log.Fatal("Invalid operator: ", o.Operator)
		return 0
	}
}

func getInput() []Operation {
	filePath := "day6/input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	operations := make([]Operation, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		parts := strings.Fields(line)

		for idx, part := range parts {
			if !isOperator(part) {
				operand, err := strconv.Atoi(part)
				if err != nil {
					log.Fatal("Invalid operand: ", part)
				}

				if len(operations) > idx {
					operations[idx].Operands = append(operations[idx].Operands, operand)
				} else {
					operations = append(operations, Operation{
						Operator: "",
						Operands: []int{operand},
					})
				}
			} else {
				operator := OperatorType(part)
				if err := operator.validate(); err != nil {
					log.Fatal("Invalid operator: ", part)
				}

				if len(operations) > idx {
					if operations[idx].Operator != "" {
						log.Fatal("Operator already set: ", part)
					}

					operations[idx].Operator = operator
				} else {
					log.Fatal("Operator found before any operands: ", part)
				}
			}
		}
	}

	return operations
}

func isOperator(in string) bool {
	return in == "+" || in == "*"
}
