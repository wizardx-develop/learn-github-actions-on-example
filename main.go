package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/wizardx-develop/learn-github-actions-on-example/numbers"
)

func getInputNumber() float64 {
	input := os.Getenv("NUMBER")
	if input == "" {
		log.Fatalf("Number doesn't to be empty")
		return 0.0
	}

	number, err := strconv.Atoi(input)
	if err != nil {
		log.Fatalf("Error while parsing number: %v", err)
		return 0.0
	}
	return float64(number)
}

func main() {
	num := getInputNumber()
	s := numbers.FindPrimeNumbers(num)
	fmt.Printf("%v\n", s)
}
