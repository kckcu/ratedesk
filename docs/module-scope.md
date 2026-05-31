# RD-01 Module Scope

RD-01 is intentionally small. The only production boundary is between the CLI and the money package.

## Build

- `cmd/ratedesk` parses CLI arguments and prints the normalized result.
- `internal/money` owns validation, parsing, formatting, and invariants.

## Do Not Build

- rate lookup
- CSV files
- conversion math
- concurrency
- service packages
- infrastructure code
