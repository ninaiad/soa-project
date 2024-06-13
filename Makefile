.PHONY: run
run:
	export GATEWAY_DB_VOLUME=./src/gateway/.database/postgres/data && \
	export POSTS_DB_VOLUME=./src/posts/.database/postgres/data && \
	export CLICKHOUSE_DB_VOLUME=./src/statistics/.database/ && \
	docker compose build ateway gateway-db posts posts-db statistics clickhouse zookeeper kafka && \
	docker compose up gateway gateway-db posts posts-db statistics clickhouse zookeeper kafka

.PHONY: stop
stop:
	docker compose down

.PHONY: test-e2e
test-e2e:
	export GATEWAY_DB_VOLUME=./tests/gateway_db/.database/postgres/data && \
	export POSTS_DB_VOLUME=./tests/posts_db/.database/postgres/data && \
	export CLICKHOUSE_DB_VOLUME=./tests/statistics_db/.database/ && \
	docker compose build gateway gateway-db posts posts-db statistics clickhouse zookeeper kafka e2e-tests && \
	docker compose up gateway gateway-db posts posts-db statistics clickhouse zookeeper kafka e2e-tests

.PHONY: test-posts
test-posts:
	export GATEWAY_DB_VOLUME=./tests/gateway_db/.database/postgres/data && \
	export POSTS_DB_VOLUME=./tests/posts_db/.database/postgres/data && \
	export CLICKHOUSE_DB_VOLUME=./tests/statistics_db/.database/ && \
	docker compose build posts posts-db posts-tests && \
	docker compose up posts posts-db posts-tests
