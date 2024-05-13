build-linux: clean-linux
	go build -o bin/linux/

build-osx: clean-osx
	go build -o bin/osx/

build-win: clean-win
	go build -o bin/windows/

clean-linux:
	rm -rf bin/linux
	mkdir -p bin/linux

clean-osx:
	rm -rf bin/osx
	mkdir -p bin/osx

clean-win:
	rm -rf bin/windows
	mkdir -p bin/windows

test:
	go test ./...

test-fresh:
	go test -count=1 ./...

test-verbose:
	go test -v ./...

test-verbose-fresh:
	go test -v -count=1 ./...

.PHONY: build-linux build-osx build-win clean-linux clean-osx clean-win test test-fresh test-verbose test-verbose-fresh
