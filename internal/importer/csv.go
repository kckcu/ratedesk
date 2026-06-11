package importer

import (
	"encoding/csv"
	"errors"
	"fmt"
	"gitlab.nebinarnitest.ru/mentoring/templates/ratedesk/02-functions-errors-starter/internal/money"
	"gitlab.nebinarnitest.ru/mentoring/templates/ratedesk/02-functions-errors-starter/internal/rates"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	ErrMalformedCSV = errors.New("malformed csv")
	ErrInvalidField = errors.New("invalid field")
)

type RowError struct {
	Row int
	Err error
}

func (e RowError) Error() string {
	return fmt.Sprintf("row %d: %v", e.Row, e.Err)
}

func (e RowError) Unwrap() error {
	return e.Err
}

func ImportCSV(r io.Reader) (*rates.Table, error) {
	reader := csv.NewReader(r)

	header, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("не получилось прочитать файл: %w", ErrMalformedCSV)
	}

	if len(header) != 3 || header[0] != "from" || header[1] != "to" || header[2] != "rate" {
		return nil, fmt.Errorf("неверные названия заголовков csv файла: %w", ErrMalformedCSV)
	}

	table, _ := rates.NewTable()
	rowNum := 1

	for {
		rowNum++
		record, err := reader.Read()

		if err == io.EOF {
			return table, nil
		}

		if err != nil {
			return nil, RowError{Row: rowNum, Err: ErrMalformedCSV}
		}

		fromCurrency, err := money.NewCurrency(record[0])
		if err != nil {
			return nil, RowError{Row: rowNum, Err: ErrInvalidField}
		}

		toCurrency, err := money.NewCurrency(record[1])
		if err != nil {
			return nil, RowError{Row: rowNum, Err: ErrInvalidField}
		}

		pair, err := rates.NewPair(fromCurrency, toCurrency)
		if err != nil {
			return nil, RowError{Row: rowNum, Err: ErrInvalidField}
		}
		scale := 1
		numbers := strings.Split(record[2], ".")
		if len(numbers) > 2 {
			return nil, RowError{Row: rowNum, Err: ErrInvalidField}
		}
		if len(numbers) == 2 {
			for range len(numbers[1]) {
				scale *= 10
			}
		}
		valueStr := strings.Join(numbers, "")
		value, err := strconv.Atoi(valueStr)
		if err != nil {
			return nil, RowError{Row: rowNum, Err: ErrInvalidField}
		}
		rate, err := rates.NewRate(pair, int64(value), int64(scale))
		if err != nil {
			return nil, RowError{Row: rowNum, Err: err}
		}
		err = table.Add(rate)
		if err != nil {
			return nil, RowError{Row: rowNum, Err: err}
		}
	}
}

type OpenFunc func(path string) (io.ReadCloser, error)

func LoadCSVWith(path string, open OpenFunc) (table *rates.Table, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("load %q: %w", path, err)
		}
	}()

	f, err := open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		if cerr := f.Close(); cerr != nil {
			err = errors.Join(err, cerr)
		}
	}()

	return ImportCSV(f)
}

func LoadCSV(path string) (*rates.Table, error) {
	return LoadCSVWith(path, func(p string) (io.ReadCloser, error) {
		return os.Open(p)
	})
}
