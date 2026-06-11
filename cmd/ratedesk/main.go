package main

import (
	"errors"
	"fmt"
	"gitlab.nebinarnitest.ru/mentoring/templates/ratedesk/02-functions-errors-starter/internal/money"
	"os"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) == 1 {
		amount, err := money.ParseAmount(args[0])
		if err != nil {
			return err
		}
		fmt.Println(amount.String())
		return nil
	}
	if len(args) > 0 && args[0] == "convert" {
		return errors.New("RD-02 TODO: implement convert command")
	}
	return fmt.Errorf(`usage: ratedesk "125.00 USD" OR ratedesk convert -rates testdata/rates.csv -amount "10.00 USD" -to EUR`)
}
