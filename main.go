package main

import (
	"fmt"

	"github.com/wizardx-develop/learn-github-actions-on-example/numbers"
)

func main() {
	s := numbers.FindPrimeNumbers(10)
	fmt.Printf("%v\n", s)

	s = numbers.FindPrimeNumbers(35)
	fmt.Printf("%v\n", s)

	s = numbers.FindPrimeNumbers(100)
	fmt.Printf("%v\n", s)
}
