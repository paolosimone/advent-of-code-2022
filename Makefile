day_pad := $(shell printf "%02d" $(day))

.PHONY: run
run:
	go run main.go $(day)

.PHONY: test
test:
	go test ./...

.PHONY: test-day
test-day:
	go test ./advent/days/day$(day_pad)/day$(day_pad)_test.go

.PHONY: format
format: 
	golangci-lint run --fix
	golangci-lint run