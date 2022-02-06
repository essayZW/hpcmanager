.PHONY: proto
proto:
	protoc -I=. --micro_out=paths=source_relative:. --go_out=paths=source_relative:. proto/*.proto
