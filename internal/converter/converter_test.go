package converter_test

import (
	"errors"
	"testing"

	"gitlab.nebinarnitest.ru/mentoring/templates/ratedesk/02-functions-errors-starter/internal/converter"
	"gitlab.nebinarnitest.ru/mentoring/templates/ratedesk/02-functions-errors-starter/internal/money"
	"gitlab.nebinarnitest.ru/mentoring/templates/ratedesk/02-functions-errors-starter/internal/rates"
)

func TestConvert_SameCurrency(t *testing.T) {
	usd, err := money.NewCurrency("USD")
	if err != nil {
		t.Fatalf("NewCurrency: %v", err)
	}
	amount, err := money.NewAmount(1000, usd)
	if err != nil {
		t.Fatalf("NewAmmount: %v", err)
	}
	table, err := rates.NewTable()
	if err != nil {
		t.Fatalf("NewTable: %v", err)
	}
	conv := converter.New(table)

	got, err := conv.Convert(amount, usd)

	if err != nil {
		t.Fatalf("Convert returned error: %v", err)
	}
	if got != amount {
		t.Errorf("got %v, want %v", got, amount)
	}
}

func TestConvert_Math(t *testing.T) {
	usd := mustCurrency(t, "USD")
	eur := mustCurrency(t, "EUR")

	tests := []struct {
		name        string
		amountMinor int64
		rateValue   int64
		rateScale   int64
		wantMinor   int64
	}{
		{
			name:        "no rounding, scale 100",
			amountMinor: 10000,              // 100.00 USD
			rateValue:   85, rateScale: 100, // 0.85
			wantMinor: 8500, // 85.00 EUR (100.00 * 0.85 = 85.0000)
		},
		{
			name:        "round down, scale 100",
			amountMinor: 9999,               // 99.99 USD
			rateValue:   92, rateScale: 100, // 0.92
			wantMinor: 9199, // 91.99 EUR (91.9908 → вниз)
		},
		{
			name:        "round up, scale 100",
			amountMinor: 5055,               // 50.55 USD
			rateValue:   49, rateScale: 100, // 0.49
			wantMinor: 2477, // 24.77 EUR (24.7695 → вверх)
		},
		{
			name:        "half boundary, scale 100",
			amountMinor: 1010,              // 10.10 USD
			rateValue:   5, rateScale: 100, // 0.05
			wantMinor: 51, // 0.51 EUR (0.505 → вверх по half-up)
		},
		{
			name:        "scale 1000, round up",
			amountMinor: 3333,                  // 33.33 USD
			rateValue:   1555, rateScale: 1000, // 1.555
			wantMinor: 5183, // 51.83 EUR (51.82815 → вверх)
		},
		{
			name:        "scale 10000, round up",
			amountMinor: 3333,                    // 33.33 USD
			rateValue:   19999, rateScale: 10000, // 1.9999
			wantMinor: 6666, // 66.66 EUR (66.6566667 → вверх)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			amount, err := money.NewAmount(tt.amountMinor, usd)
			if err != nil {
				t.Fatalf("NewAmount: %v", err)
			}
			pair, err := rates.NewPair(usd, eur)
			if err != nil {
				t.Fatalf("NewPair: %v", err)
			}
			rate, err := rates.NewRate(pair, tt.rateValue, tt.rateScale)
			if err != nil {
				t.Fatalf("NewRate: %v", err)
			}
			table, err := rates.NewTable(rate)
			if err != nil {
				t.Fatalf("newTable: %v", err)
			}
			conv := converter.New(table)

			got, err := conv.Convert(amount, eur)

			if err != nil {
				t.Fatalf("Convert returned error: %v", err)
			}
			if got.Minor() != tt.wantMinor {
				t.Errorf("got minor = %d, want %d", got.Minor(), tt.wantMinor)
			}
			if got.Currency() != eur {
				t.Errorf("got currency = %v, want %v", got.Currency(), eur)

			}
		})
	}
}

func TestConvert_RateNotFound(t *testing.T) {
	usd := mustCurrency(t, "USD")
	eur := mustCurrency(t, "EUR")

	table, err := rates.NewTable()
	if err != nil {
		t.Fatalf("NewTable: %v", err)
	}
	conv := converter.New(table)

	amount, err := money.NewAmount(10000, usd)
	if err != nil {
		t.Fatalf("NewAmount: %v", err)
	}

	_, err = conv.Convert(amount, eur)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, rates.ErrRateNotFound) {
		t.Errorf("expected ErrRateNotFound, got %v", err)
	}
}

func mustCurrency(t *testing.T, code string) money.Currency {
	t.Helper()
	c, err := money.NewCurrency(code)
	if err != nil {
		t.Fatalf("NewCurrency(%q): %v", code, err)
	}
	return c
}
