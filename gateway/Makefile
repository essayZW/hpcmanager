.PHONY: build
build:
	go build -o build/main ./*.go
.PHONY: clean
clean:
	rm -rf ./build
.PHONY: proto
proto:
	protoc -I=. -I=../ --go_out=paths=source_relative:. proto/*.proto
.PHONY: dev
dev:
	env HPCMANAGER_ENV=dev go run main.go -port 8080
.PHONY: docker
docker:
	sudo docker build . -t hpcmanager/gateway:1.0
