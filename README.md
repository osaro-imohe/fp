- **Question: What is the max USD value that can be processed in 50ms, 60ms, 90ms?**
  - In 50ms => _$3637.98_
  - In 60ms => _$4362.01_
  - In 90ms => _$6870.48_
  - In 1000ms => _$35289.19999999999_

### START

- go mod tidy
- go run main.go `$totalTransactionTime`
