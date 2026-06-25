package importer_test

import (
	"errors"
	"io"
	"strings"
	"testing"

	"gitlab.nebinarnitest.ru/mentoring/templates/ratedesk/02-functions-errors-starter/internal/importer"
	"gitlab.nebinarnitest.ru/mentoring/templates/ratedesk/02-functions-errors-starter/internal/rates"
)

func TestImportCSV_HappyPath(t *testing.T) {
	input := "from,to,rate\n" +
		"USD,EUR,0.92\n" +
		"USD,GBP,0.79\n"

	table, err := importer.ImportCSV(strings.NewReader(input))
	if err != nil {
		t.Fatalf("ImportCSV: %v", err)
	}
	if got := table.Len(); got != 2 {
		t.Errorf("Len = %d, want 2", got)
	}
}

func TestImportCSV_BadHeader(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "wrong column names",
			input: "a,b,c\nUSD,EUR,0.92\n",
		},
		{
			name:  "too few columns",
			input: "from,to\nUSD,EUR\n",
		},
		{
			name:  "empty input",
			input: "",
		},
		{
			name:  "too many columns",
			input: "from,to, rate, rate\nUSD,EUR, 0.92, 0.92\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := importer.ImportCSV(strings.NewReader(tt.input))
			if err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !errors.Is(err, importer.ErrMalformedCSV) {
				t.Errorf("expected ErrMalformedCSV, got %v", err)
			}
		})
	}
}

func TestImportCSV_BadRow(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantRow int
		wantErr error
	}{
		{
			name: "invalid currency in row 2",
			input: "from,to,rate\n" +
				"US,EUR,0.92\n",
			wantRow: 2,
			wantErr: importer.ErrInvalidField,
		},
		{
			name: "invalid rate format in row 3",
			input: "from,to,rate\n" +
				"USD,EUR,0.92\n" +
				"USD,GBP,1.2.3\n",
			wantRow: 3,
			wantErr: importer.ErrInvalidField,
		},
		{
			name: "same currency pair in row 2",
			input: "from,to,rate\n" +
				"USD,USD,1.00\n",
			wantRow: 2,
			wantErr: importer.ErrInvalidField,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := importer.ImportCSV(strings.NewReader(tt.input))
			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			var rowErr importer.RowError
			if !errors.As(err, &rowErr) {
				t.Fatalf("expected RowError in chain, got %v", err)
			}

			if rowErr.Row != tt.wantRow {
				t.Errorf("Row = %d, want %d", rowErr.Row, tt.wantRow)
			}

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("expected %v in chain, got %v", tt.wantErr, err)
			}
		})
	}
}

func TestImportCSV_DublicateRate(t *testing.T) {
	input := "from,to,rate\n" +
		"USD,EUR,0.92\n" +
		"USD,EUR,0.95\n"

	_, err := importer.ImportCSV(strings.NewReader(input))
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if !errors.Is(err, rates.ErrDuplicate) {
		t.Errorf("expected ErrDuplicate, got %v", err)
	}

	var rowErr importer.RowError
	if !errors.As(err, &rowErr) {
		t.Fatalf("expected RowError in chain, got %v", err)
	}
	if rowErr.Row != 3 {
		t.Errorf("Row = %d, want 3", rowErr.Row)
	}
}

type fakeReadCloser struct {
	data       *strings.Reader
	readErr    error
	closeErr   error
	closeCount int
}

func (f *fakeReadCloser) Read(p []byte) (int, error) {
	if f.readErr != nil {
		return 0, nil
	}
	return f.data.Read(p)
}

func (f *fakeReadCloser) Close() error {
	f.closeCount++
	return f.closeErr
}

func TestLoadCSV_HappyPath(t *testing.T) {
	fake := &fakeReadCloser{
		data: strings.NewReader("from,to,rate\nUSD,EUR,0.92\n"),
	}
	open := func(p string) (io.ReadCloser, error) { return fake, nil }

	table, err := importer.LoadCSVWith("rates.csv", open)
	if err != nil {
		t.Fatalf("LoadCSVWith: %v", err)
	}
	if got := table.Len(); got != 1 {
		t.Errorf("Len = %d, want 1", got)
	}
	if fake.closeCount != 1 {
		t.Errorf("Close called %d times, want 1", fake.closeCount)
	}
}

func TestLoadCSV_OpenError(t *testing.T) {
	sentinel := errors.New("disk on fire")
	open := func(p string) (io.ReadCloser, error) { return nil, sentinel }

	_, err := importer.LoadCSVWith("rates.csv", open)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), `load "rates.csv"`) {
		t.Errorf("error does not contain file context: %v", err)
	}
	if !errors.Is(err, sentinel) {
		t.Errorf("expected sentinel in chain, got %v", err)
	}
}

func TestLoadCSV_ReadError(t *testing.T) {
	sentinel := errors.New("read failed")
	fake := &fakeReadCloser{
		data:    strings.NewReader(""),
		readErr: sentinel,
	}
	open := func(p string) (io.ReadCloser, error) { return fake, nil }

	_, err := importer.LoadCSVWith("rates.csv", open)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), `load "rates.csv"`) {
		t.Errorf("error does not contain file context: %v", err)
	}
	if fake.closeCount != 1 {
		t.Errorf("Close called %d times, want 1", fake.closeCount)
	}
}

func TestLoadCSV_CloseError(t *testing.T) {
	closeSentinel := errors.New("close failed")
	fake := &fakeReadCloser{
		data:     strings.NewReader("from,to,rate\nUSD,EUR,0.92\n"),
		closeErr: closeSentinel,
	}
	open := func(p string) (io.ReadCloser, error) { return fake, nil }

	_, err := importer.LoadCSVWith("rates.csv", open)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, closeSentinel) {
		t.Errorf("expected close sentinel in chain, got %v", err)
	}
	if !strings.Contains(err.Error(), `load "rates.csv"`) {
		t.Errorf("error does not contain file context: %v", err)
	}
}
