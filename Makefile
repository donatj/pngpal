BIN=pngpal

HEAD=$(shell git describe --tags 2> /dev/null  || git rev-parse --short HEAD)

all: build

build: darwin64 linux64 freebsd64 windows64

clean:
	-rm -f $(BIN)
	-rm -rf release

darwin64:
	env GOOS=darwin GOARCH=amd64 go clean -i ./...
	env GOOS=darwin GOARCH=amd64 go build -o release/darwin64/$(BIN) ./cmd/pngpal

linux64:
	env GOOS=linux GOARCH=amd64 go clean -i ./...
	env GOOS=linux GOARCH=amd64 go build -o release/linux64/$(BIN) ./cmd/pngpal

freebsd64:
	env GOOS=freebsd GOARCH=amd64 go clean -i ./...
	env GOOS=freebsd GOARCH=amd64 go build -o release/freebsd64/$(BIN) ./cmd/pngpal

windows64:
	env GOOS=windows GOARCH=amd64 go clean -i ./...
	env GOOS=windows GOARCH=amd64 go build -o release/windows64/$(BIN).exe ./cmd/pngpal

.PHONY: release
release: clean build
	zip -9 release/$(BIN).darwin_amd64.$(HEAD).zip release/darwin64/$(BIN)
	zip -9 release/$(BIN).linux_amd64.$(HEAD).zip release/linux64/$(BIN)
	zip -9 release/$(BIN).freebsd_amd64.$(HEAD).zip release/freebsd64/$(BIN)
	zip -9 release/$(BIN).windows_amd64.$(HEAD).zip release/windows64/$(BIN).exe