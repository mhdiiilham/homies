.PHONY:

run:
	go run cmd/rest/main.go --cfg=config.local.yaml

test:
	go test -cover -v -race ./...
