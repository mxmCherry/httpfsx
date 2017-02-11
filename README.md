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

For `internal/statichandler`, install `statik` generator:

```bash
go get -u github.com/rakyll/statik github.com/rakyll/statik/fs
```

Then, after editing it, do:

```
go generate ./...
```
