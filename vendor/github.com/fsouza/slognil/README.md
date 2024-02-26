# slognil

Discard handler for slog. Want do discard all logs? Use this handler, without
bothering with levels (useful for tests):

```go
logger := slognil.NewLogger()
```
