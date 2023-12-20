.PHONY: lint
lint:
	golangci-lint run

.PHONY: dev
dev:
	go run cmd/web/*.go
