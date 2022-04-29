.PHONY: proto
proto:
	protoc -I=. -I=../ --micro_out=paths=source_relative:. --go_out=paths=source_relative:. proto/*.proto

.PHONY: build
build:
	go build -o build/main ./*.go
.PHONY: clean
clean:
	rm -rf ./build
.PHONY: dev
dev:
	env HPCMANAGER_ENV=dev go run main.go
.PHONY: docker
docker:
	sudo docker build . -t hpcmanager/fee:1.0
