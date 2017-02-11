# httpfsx

Mobile-friendly HTTP file-system explorer (readonly)

# Installing

```bash
go get -u github.com/mxmCherry/httpfsx
```

# Running

```bash
httpfsx --addr=:1024 --root=$HOME/share
```

# Dev notes

After editing `internal/statichandler/public/*`, run:

```
go generate ./...
```
