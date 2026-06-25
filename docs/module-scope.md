# RD-02 Module Scope

RD-02 adds function boundaries and error handling around currency conversion.

## Build

- `internal/rates` owns pairs, rates, lookup, and duplicate policy.
- `internal/converter` converts an amount using an existing rate table.
- `internal/importer` reads CSV input and returns a table.
- `cmd/ratedesk` wires flags to the use case.

## Do Not Build

- goroutines or worker pools;
- provider frameworks;
- HTTP, DB, Redis, Kafka, cache;
- broad architecture layers.
