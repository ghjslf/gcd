package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("expected 2 arguments after executable file path")
	}

	a, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	b, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Greatest common divisor of numbers %d and %d is %d\n", a, b, gcd(a, b))
}

func gcd(a int, b int) int {
	for a != 0 && b != 0 {
		if a > b {
			a = a % b
		} else {
			b = b % a
		}
	}

	var gcd int

	if a > b {
		gcd = a
	} else {
		gcd = b
	}

	return gcd
}
