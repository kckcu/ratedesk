# RD-02 Review Checklist

- CSV contract is exactly `from,to,rate`.
- File open errors include path context and still preserve the underlying error.
- Domain sentinel errors work with `errors.Is` after wrapping.
- Converter does not read files or parse CSV itself.
- Tests cover success, invalid input, missing rate, duplicate rate, and malformed CSV.
- `make check` is green.
