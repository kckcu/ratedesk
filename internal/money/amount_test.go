package money

import "testing"

func TestParseAmountSpec(t *testing.T) {

	tests := []struct {
		name string
		raw  string
		want string
	}{
		{name: "whole amount", raw: "123 USD", want: "123.00 USD"},
		{name: "minor units", raw: "123.45 USD", want: "123.45 USD"},
		{name: "one decimal place", raw: "12.5 USD", want: "12.50 USD"},
		{name: "leading zero in fraction", raw: "12.05 USD", want: "12.05 USD"},
		{name: "zero amount", raw: "0 USD", want: "0.00 USD"},
		{name: "zero with decimals", raw: "0.00 USD", want: "0.00 USD"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseAmount(tt.raw)
			if err != nil {
				t.Fatalf("ParseAmount() error = %v", err)
			}
			if got.String() != tt.want {
				t.Fatalf("String() = %q, want %q", got.String(), tt.want)
			}
		})
	}
}

func TestInvalidAmountSpec(t *testing.T) {
	tests := []struct {
		name string
		raw  string
	}{
		{name: "too many fractional digits", raw: "123.456 USD"},
		{name: "empty string", raw: ""},
		{name: "invalid currency length", raw: "125 US"},
		{name: "lowercase currency", raw: "125 usd"},
		{name: "negative amount", raw: "-125 USD"},
		{name: "garbage string", raw: "what is this"},
		{name: "no currency", raw: "125.00"},
		{name: "too many dots", raw: ".12."},
		{name: "no currency code", raw: "12."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseAmount(tt.raw)
			if err == nil {
				t.Fatalf("expected error for input %q, but got nil", tt.raw)
			}
		})
	}
}

func TestParseAmountsSpec(t *testing.T) {
	t.Run("nil input", func(t *testing.T) {
		got := ParseAmounts(nil)
		if got != nil {
			t.Errorf("ParseAmounts(nil) = %v, want nil", got)
		}
	})
	t.Run("empty input", func(t *testing.T) {
		got := ParseAmounts([]string{})
		if len(got) != 0 {
			t.Errorf("ParseAmounts([]) len = %d, want 0", len(got))
		}
	})
	t.Run("valid inputs preserve order", func(t *testing.T) {
		inputs := []string{"1 USD", "2.50 EUR"}
		got := ParseAmounts(inputs)
		if len(got) != 2 {
			t.Fatalf("ParseAmounts len = %d, want 2", len(got))
		}
		if got[0].String() != "1.00 USD" {
			t.Errorf("got[0] = %q, want %q", got[0].String(), "1.00 USD")
		}
		if got[1].String() != "2.50 EUR" {
			t.Errorf("got[1] = %q, want %q", got[1].String(), "2.50 EUR")
		}
	})

}
