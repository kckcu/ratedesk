# RD-02-2 Converter

## Outcome

Implement amount conversion using an existing rate table.

## Acceptance

- Converter does not read CSV or know file paths.
- Missing rate preserves `rates.ErrRateNotFound`.
- Rounding policy is documented by tests.
