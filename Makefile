.PHONY: lint
lint:
	golangci-lint run

.PHONY: dev
dev:
	air
