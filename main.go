package main

import (
	"fmt"

	"github.com/wizardx-develop/learn-github-actions-on-example/numbers"
)

func main() {
	s := numbers.FindPrimeNumbers(10)
	fmt.Printf("%v\n", s)
}
