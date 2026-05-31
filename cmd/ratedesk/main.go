package main

import (
	"fmt"
	"gitlab.nebinarnitest.ru/mentoring/templates/ratedesk/01-basics-starter/internal/money"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, `usage: ratedesk "125.00 USD"`)
		os.Exit(2)
	}

	input := os.Args[1]

	amount, err := money.ParseAmount(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(amount)
}
