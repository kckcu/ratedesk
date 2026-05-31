# RD-01 Testing Strategy

Write table tests around externally observable behavior:

- valid money inputs normalize to the expected string;
- invalid inputs return errors instead of panics;
- constructors protect invariants;
- CLI glue stays thin and delegates parsing to `internal/money`.

Do not test private parsing steps directly. Prefer examples that match README commands and MR acceptance criteria.
