# go-design-patterns
Design Patterns with Go Programming Language


## Code Actions

The following commands are to run the actions for a better code quality.

### Unit Tests

In the root of the project (go-design-patterns/), run
```bash
go test -v ./...
```

### golangci-lint
[golangci-lint](https://github.com/golangci/golangci-lint) is a fast Go linters runner. It runs linters in parallel, uses caching, supports yaml config, has integrations with all major IDE and has dozens of linters included.

Download the golangci-lint
```bash
$ curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
```

Run the golangci-lint
```bash
$(go env GOPATH)/bin/golangci-lint run -v
```

### gofumpt
[gofumpt](https://github.com/mvdan/gofumpt) this version of gofumpt enforces a stricter format than gofmt and happy subset of the formats that gofmt is happy with.

Install
```bash
go install mvdan.cc/gofumpt@latest
```

Run
```bash
gofumpt -l -w .
```
