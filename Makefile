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

.PHONY: new-day
new-day:
	mkdir -p ./advent/days/day$(day_pad)

	cp ./advent/days/day00/day00.go ./advent/days/day$(day_pad)/day$(day_pad).go
	cp ./advent/days/day00/day00_test.go ./advent/days/day$(day_pad)/day$(day_pad)_test.go
	grep -rl 00 ./advent/days/day$(day_pad) | xargs sed -i 's/00/$(day_pad)/g'
	touch ./advent/days/day$(day_pad)/input.txt

	perl -i -p0 -e 's|(?<=import \()(.*?)(?=\))|$$1"github.com/paolosimone/advent-of-code-2022/advent/days/day$(day_pad)"\n|s;' ./advent/advent.go
	perl -i -p0 -e 's/(?<=DayLoader\{)(.*?)(?=\})/$$1day$(day_pad).Load,\n/s;' ./advent/advent.go
	golangci-lint run --fix