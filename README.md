# ratedesk

CLI tool that parses and normalises monetary amounts.

## Requirements

Go 1.22 or later.

## Build

```bash
go build ./...
```

## Run

```bash
go run ./cmd/ratedesk/ "125.00 USD"
```

After building:

```bash
./ratedesk "125.00 USD"
# 125.00 USD
```

Invalid input prints to stderr and exits with a non-zero code:

```bash
./ratedesk "abc USD"
# error: invalid amount
```

## Test

```bash
go test ./...
```

## Check

Runs fmt, test, and vet:

```bash
make check
```