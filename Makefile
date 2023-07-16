
all: clean build test coverage benchmark

clean:
	rm -f coverage.out
	go clean -testcache -testcache

build:
	go build ./...

#binary:
#	go build -o binname cli/binname/main.go

test:
	go test ./...

coverage:
	go test -cover -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

benchmark: build
	go test -run=Benchmark -bench=. ./...


