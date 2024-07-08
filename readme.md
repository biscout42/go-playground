# Learning golang and experiments

## Mutext understanding and running tests

There is a `./internal/util/mutex.go` and two tests in `p1` and `p2` packages. If you run 
```shell
go test -mod=vendor -v -cover -race ./...
```

you will see, that tests from same package respect the mutex, but not tests from different package. The question is why? 