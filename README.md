# About

Answers for AlpineIQ [code test](https://gist.github.com/shahzilabid/787cd00cbf0af0b5c4adf8b40234abd4)

* To run, `go run main.go`

* To run tests, `go test -v -bench=. ./...`. This takes about 50 seconds. There are benchmark tests for concurrent and non-current execution. In these test, real world delays are simulated using `time.Sleep`.
