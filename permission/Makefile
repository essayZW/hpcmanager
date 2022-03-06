.PHONY: proto
proto:
	protoc -I=. -I=../ --micro_out=paths=source_relative:. --go_out=paths=source_relative:. proto/*.proto

.PHONY: build
build:
	go build -o build/main ./*.go
.PHONY: clean
clean:
	rm -rf ./build
