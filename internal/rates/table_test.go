package rates_test

import (
	"errors"
	"testing"

	"gitlab.nebinarnitest.ru/mentoring/templates/ratedesk/02-functions-errors-starter/internal/money"
	"gitlab.nebinarnitest.ru/mentoring/templates/ratedesk/02-functions-errors-starter/internal/rates"
)

func TestTable_AddAndLookup(t *testing.T) {
	usd := mustCurrency(t, "USD")
	eur := mustCurrency(t, "EUR")

	pair, err := rates.NewPair(usd, eur)
	if err != nil {
		t.Fatalf("NewPair: %v", err)
	}
	rate, err := rates.NewRate(pair, 92, 100)
	if err != nil {
		t.Fatalf("NewRate: %v", err)
	}

	table, err := rates.NewTable()
	if err != nil {
		t.Fatalf("NewTable: %v", err)
	}

	if err := table.Add(rate); err != nil {
		t.Fatalf("Add: %v", err)
	}

	if got := table.Len(); got != 1 {
		t.Errorf("Len = %d, want 1", got)
	}

	got, err := table.Lookup(pair)
	if err != nil {
		t.Fatalf("Lookup: %v", err)
	}
	if got != rate {
		t.Errorf("Lookup = %+v, want %+v", got, rate)
	}
}

func TestTable_AddDuplicate(t *testing.T) {
	usd := mustCurrency(t, "USD")
	eur := mustCurrency(t, "EUR")

	pair, err := rates.NewPair(usd, eur)
	if err != nil {
		t.Fatalf("NewPair: %v", err)
	}
	rate, err := rates.NewRate(pair, 92, 100)
	if err != nil {
		t.Fatalf("NewRate: %v", err)
	}

	table, err := rates.NewTable(rate)
	if err != nil {
		t.Fatalf("NewTable: %v", err)
	}

	err = table.Add(rate)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, rates.ErrDuplicate) {
		t.Errorf("expected ErrDuplicate, got %v", err)
	}
}

func TestTable_LookupNotFound(t *testing.T) {
	usd := mustCurrency(t, "USD")
	eur := mustCurrency(t, "EUR")

	pair, err := rates.NewPair(usd, eur)
	if err != nil {
		t.Fatalf("NewPair: %v", err)
	}

	table, err := rates.NewTable()
	if err != nil {
		t.Fatalf("NewTable: %v", err)
	}

	_, err = table.Lookup(pair)
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
