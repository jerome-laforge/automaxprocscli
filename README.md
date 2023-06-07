# automaxprocscli
https://github.com/uber-go/automaxprocs/issues/69

# usage
```
automaxprocscli -h

Usage of automaxprocscli:

  -v    display GOMAXPROCS before and after or if no change has occurred
```

# example:
```
- go install github.com/jerome-laforge/automaxprocscli@latest
go: downloading github.com/jerome-laforge/automaxprocscli v0.0.0-20230607074117-ff6b276e3325
go: downloading go.uber.org/automaxprocs v1.5.2
- $ ${GOPATH}/bin/automaxprocscli -v
GOMAXPROCS before: 16 after: 5
- go test -cpu $(${GOPATH}/bin/automaxprocscli) ./...
```
