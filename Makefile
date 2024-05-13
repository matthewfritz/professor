build: clean
	go build -o bin/

clean:
	rm bin/*.*

test:
	go test -v ./...

.PHONY: build test
