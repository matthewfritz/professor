build: clean
	go build -o bin/

clean:
	rm bin/*.*

test:
	go test ./...

test-fresh:
	go test -count=1 ./...

test-verbose:
	go test -v ./...

test-verbose-fresh:
	go test -v -count=1 ./...

.PHONY: build clean test test-fresh test-verbose test-verbose-fresh
