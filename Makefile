all: test

test:
	golangci-lint run
