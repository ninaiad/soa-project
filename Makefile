.PHONY: run
run:
	docker compose build gateway gateway-db posts posts-db statistics clickhouse zookeeper kafka && \
	docker compose up gateway gateway-db posts posts-db statistics clickhouse zookeeper kafka

.PHONY: stop
stop:
	docker compose down

.PHONY: test-e2e
test-e2e:
	docker compose build gateway gateway-db posts posts-db statistics clickhouse zookeeper kafka e2e-tests && \
	docker compose up gateway gateway-db posts posts-db statistics clickhouse zookeeper kafka e2e-tests

.PHONY: test-posts
test-posts:
	docker compose build posts posts-db posts-tests && \
	docker compose up posts posts-db posts-tests

.PHONY: lint
lint:
	gofmt -w .
	cd src/gateway && golangci-lint run ./internal/... ./cmd/...
	cd src/posts && golangci-lint run ./internal/... ./cmd/...
	cd src/statistics && golangci-lint run ./internal/...
	cd tests/posts && golangci-lint run .
