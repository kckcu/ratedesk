package converter

import (
	"gitlab.nebinarnitest.ru/mentoring/templates/ratedesk/02-functions-errors-starter/internal/money"
	"gitlab.nebinarnitest.ru/mentoring/templates/ratedesk/02-functions-errors-starter/internal/rates"
)

type Converter struct {
	table *rates.Table
}

func New(table *rates.Table) *Converter {
	return &Converter{table: table}
}

func (c *Converter) Convert(amount money.Amount, to money.Currency) (money.Amount, error) {
	if amount.Currency() == to {
		return amount, nil
	}
	pair, err := rates.NewPair(amount.Currency(), to)
	if err != nil {
		return money.Amount{}, err
	}

	rate, err := c.table.Lookup(pair)
	if err != nil {
		return money.Amount{}, err
	}

	raw := amount.Minor() * rate.Value
	quotient := raw / rate.Scale
	remainder := raw % rate.Scale

	if remainder*2 >= rate.Scale {
		quotient++
	}

	return money.NewAmount(quotient, to)

}
