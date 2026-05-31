# RD-01 Review Checklist

- Money is represented as integer minor units, not `float64`.
- Currency validation rejects non-ISO-like codes.
- Parser rejects garbage, empty input, and too many decimal digits.
- `cmd/ratedesk` does not contain domain arithmetic.
- Tests describe behavior and include failure cases.
- `make check` is green.
