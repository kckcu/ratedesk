# RD-01-3 CLI glue

## Outcome

Wire the money parser into `cmd/ratedesk`.

## Acceptance

- `go run ./cmd/ratedesk "125.00 USD"` prints `125.00 USD`.
- Invalid input prints an error and exits non-zero.
- CLI remains thin and does not duplicate money logic.
