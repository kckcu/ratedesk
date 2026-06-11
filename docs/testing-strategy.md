# RD-02 Testing Strategy

Cover success and failure paths:

- rate validation and duplicate insertion;
- missing rate behavior with `errors.Is(err, rates.ErrRateNotFound)`;
- converter rounding policy;
- CSV malformed rows, invalid currencies, invalid rates, and file open failures;
- CLI smoke command with `testdata/rates.csv`.

Assertions should check sentinel errors through `errors.Is` and structured row context through `errors.As`, not fragile full message strings.
