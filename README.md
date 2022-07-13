- **Question: What is the max USD value that can be processed in 50ms, 60ms, 90ms?**
  - In 50ms => _$4139.43_
  - In 60ms => _$4675.71_
  - In 90ms => _$6972.290000000001_
  - In 1000ms => _$35471.810000000005_

### START

- go mod tidy
- go run main.go `$totalTransactionTime`
