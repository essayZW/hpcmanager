.PHONY: proto
proto:
	protoc -I=. --go_out=paths=source_relative:. proto/*.proto
