# RD-02-4 CLI convert command

## Outcome

Wire importer and converter into `ratedesk convert`.

## Acceptance

- `ratedesk convert -rates testdata/rates.csv -amount "10.00 USD" -to EUR` prints a normalized amount.
- Bad flags or invalid data exit non-zero with a useful error.
- CLI remains composition glue.
