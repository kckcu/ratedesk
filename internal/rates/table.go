package rates

import (
	"errors"
	"gitlab.nebinarnitest.ru/mentoring/templates/ratedesk/02-functions-errors-starter/internal/money"
)

var (
	ErrInvalidRate  = errors.New("invalid rate")
	ErrRateNotFound = errors.New("rate not found")
	ErrDuplicate    = errors.New("duplicate rate")
	ErrInvalidPair  = errors.New("invalid pair")
)

type Pair struct {
	From money.Currency
	To   money.Currency
}

type Rate struct {
	Pair  Pair
	Value int64
	Scale int64
}

type Table struct {
	rates map[Pair]Rate
}

func NewPair(from, to money.Currency) (Pair, error) {
	if from == to {
		return Pair{}, ErrInvalidPair
	}
	p := Pair{From: from, To: to}
	return p, nil
}

func isPowerOf10(n int64) bool {
	if n <= 0 {
		return false
	}
	for n > 1 {
		if n%10 != 0 {
			return false
		}
		n /= 10
	}
	return true
}
func NewRate(pair Pair, value, scale int64) (Rate, error) {
	if value <= 0 {
		return Rate{}, ErrInvalidRate
	}
	trueScale := isPowerOf10(scale)
	if !trueScale {
		return Rate{}, ErrInvalidRate
	}

	r := Rate{Pair: pair, Value: value, Scale: scale}
	return r, nil
}

func NewTable(items ...Rate) (*Table, error) {
	t := &Table{rates: make(map[Pair]Rate)}
	for _, r := range items {
		if err := t.Add(r); err != nil {
			return nil, err
		}
	}
	return t, nil
}

func (t *Table) Add(rate Rate) error {
	if _, exists := t.rates[rate.Pair]; exists {
		return ErrDuplicate
	}
	t.rates[rate.Pair] = rate
	return nil
}

func (t *Table) Lookup(pair Pair) (Rate, error) {
	if rate, exists := t.rates[pair]; exists {
		return rate, nil
	}
	return Rate{}, ErrRateNotFound
}

func (t *Table) Len() int {
	return len(t.rates)
}
