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
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter algebraic equation: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Invalid input: %v", err)
	}

	fmt.Printf("The result: %d\n", evalAlgebra(input))
}

func evalAlgebra(expr string) int {
	parts := strings.Fields(expr)
	if len(parts) != 5 || parts[3] != "=" {
		log.Fatalf("invalid expression format: %s", expr)
	}

	left1, op, left2, _, right := parts[0], parts[1], parts[2], parts[3], parts[4]

	toInt := func(s string) (int, bool) {
		if s == "x" {
			return 0, true
		}
		val, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("invalid number: %s", s)
		}
		return val, false
	}

	a, aIsX := toInt(left1)
	b, bIsX := toInt(left2)
	c, _ := toInt(right) // right is always int, no "x" case

	switch op {
	case "+":
		if aIsX {
			return c - b
		}
		if bIsX {
			return c - a
		}
	case "-":
		if aIsX {
			return c + b
		}
		if bIsX {
			return a - c
		}
	default:
		log.Fatalf("unsupported operator: %s", op)
	}

	return 0
}
