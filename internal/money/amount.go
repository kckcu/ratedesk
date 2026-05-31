package money

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

var (
	ErrInvalidAmount   = errors.New("invalid amount")
	ErrInvalidCurrency = errors.New("invalid currency")
)

type Currency string

type Amount struct {
	minor    int64
	currency Currency
}

func NewCurrency(code string) (Currency, error) {
	if utf8.RuneCountInString(code) != 3 {
		return "", ErrInvalidCurrency
	}
	for _, v := range code {
		if !(v >= 'A' && v <= 'Z') {
			return "", ErrInvalidCurrency
		}
	}
	return Currency(code), nil
}

func NewAmount(minor int64, currency Currency) (Amount, error) {
	if minor < 0 {
		return Amount{}, ErrInvalidAmount
	}
	return Amount{minor: minor, currency: currency}, nil
}

func ParseAmount(input string) (Amount, error) {
	pars := strings.Split(input, " ")
	if len(pars) != 2 {
		return Amount{}, ErrInvalidAmount
	}

	currency, err := NewCurrency(pars[1])
	if err != nil {
		return Amount{}, err
	}

	c := strings.Split(pars[0], ".")
	if len(c) > 2 {
		return Amount{}, ErrInvalidAmount
	}

	if len(c) == 1 {
		whole, err := strconv.ParseInt(c[0], 10, 64)
		if err != nil {
			return Amount{}, ErrInvalidAmount
		}
		return NewAmount(whole*100, currency)
	}

	if c[0] == "" || c[1] == "" {
		return Amount{}, ErrInvalidAmount
	}
	if len(c[1]) > 2 {
		return Amount{}, ErrInvalidAmount
	}
	if strings.Contains(c[1], "+") || strings.Contains(c[0], "+") {
		return Amount{}, ErrInvalidAmount
	}

	firstZero := c[1][0] == '0'
	whole, err := strconv.ParseInt(c[0], 10, 64)
	if err != nil {
		return Amount{}, ErrInvalidAmount
	}
	frac, err := strconv.ParseInt(c[1], 10, 64)
	if err != nil {
		return Amount{}, ErrInvalidAmount
	}
	if frac < 0 {
		return Amount{}, ErrInvalidAmount
	}
	if frac <= 9 && !firstZero {
		return NewAmount(whole*100+frac*10, currency)
	}
	return NewAmount(whole*100+frac, currency)
}

func ParseAmounts(inputs []string) []Amount {
	if len(inputs) == 0 {
		return nil
	}
	result := make([]Amount, 0, len(inputs))
	for _, s := range inputs {
		a, err := ParseAmount(s)
		if err == nil {
			result = append(result, a)
		}
	}
	return result
}

func (a Amount) Minor() int64 {
	return a.minor
}

func (a Amount) Currency() Currency {
	return a.currency
}

func (a Amount) String() string {
	whole := a.minor / 100
	frac := a.minor % 100
	return fmt.Sprintf("%d.%02d %s", whole, frac, a.currency)
}
