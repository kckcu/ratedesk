# RD-02-1 Rates table

## Outcome

Implement `Pair`, `Rate`, and `Table`.

## Acceptance

- Same-currency pairs are rejected.
- Non-positive rates are rejected.
- Missing lookup returns `ErrRateNotFound`.
- Duplicate pairs are rejected.
- Tests cover validation and lookup behavior.
