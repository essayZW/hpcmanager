.PHONY: build
build:
	go build -o build/main ./*.go
.PHONY: clean
clean:
	rm -rf ./build
.PHONY: proto
proto:
	protoc -I=. -I=../ --go_out=paths=source_relative:. proto/*.proto
