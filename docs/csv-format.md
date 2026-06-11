# RD-02 CSV Format

The importer accepts CSV with a required header:

```text
from,to,rate
USD,EUR,0.92
```

Rules:

- `from` and `to` are three-letter currency codes.
- `rate` is a positive decimal multiplier.
- duplicate pairs are rejected.
- row and field context should be present in returned errors.
