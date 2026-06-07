# RD-01-2 Parse amount

## Outcome

Implement `ParseAmount(input string)` for inputs such as `123 USD` and `123.45 USD`.

## Acceptance

- `123 USD` becomes `123.00 USD`.
- `123.45 USD` stays `123.45 USD`.
- Empty, malformed, and too-precise inputs return errors.
- Tests assert errors without depending on exact message text.
