# RateDesk RD-02 Starter

This repository is the starter template for RD-02: Functions and Errors.

It starts from the completed RD-01 money model and adds package skeletons for rates, conversion, and CSV import.

## Target CLI

After the module is complete:

```bash
go run ./cmd/ratedesk convert -rates testdata/rates.csv -amount "10.00 USD" -to EUR
```

## CSV Contract

CSV files use this header:

```text
from,to,rate
USD,EUR,0.92
```

## Commands

```bash
make fmt
make test
make vet
make check
```

## Out of Scope

- goroutines and worker pools;
- HTTP, DB, Redis, Kafka, cache;
- broad provider interfaces;
- retry frameworks.
