.PHONY: run
run:
	go run main.go $(day)

.PHONY: test
test:
	go test ./...

.PHONY: test-day
test-day:
	echo TODO

.PHONY: format
format: 
	golangci-lint run --fix
	golangci-lint run