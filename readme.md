# Learning golang and experiments

## Mutext and running tests - different packages different process

There is a `./internal/util/mutex.go` and two tests in `p1` and `p2` packages. If you run 
```shell
go test -mod=vendor -v -cover -race ./...
```

you will see, that tests from same package respect the mutex, but not tests from different package. The question is why?
During `go test ./...` , the `go` command will build separate binaries and execute them in parallel. By default, they will be run in parallel according number of CPUs (controlled by the `-p` flag).