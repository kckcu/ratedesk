# RD-01-1 Money model

## Outcome

Create `Currency`, `Amount`, constructors, and `Amount.String`.

## Acceptance

- Amount uses integer minor units.
- Currency codes are normalized to uppercase and validated.
- Negative amounts are rejected for this module.
- No `float64` appears in money code.
- Unit tests cover constructor invariants.
