# RD-02-3 CSV importer

## Outcome

Implement `ImportCSV(io.Reader)` and `LoadCSV(path)`.

## Acceptance

- CSV header is `from,to,rate`.
- Bad rows include row and field context.
- `LoadCSV` closes files with `defer` after successful open.
- `errors.Is` works for malformed CSV and invalid fields.
