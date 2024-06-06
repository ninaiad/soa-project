#!/bin/bash

MAIN_DB_VOLUME='./tests/main_db/.database/postgres/data' \
POSTS_DB_VOLUME='./tests/posts_db/.database/postgres/data' \
CLICKHOUSE_DB_VOLUME='./tests/statistics_db/.database/' \
  docker compose up \
  main \
  main-db \
  posts \
  posts-db \
  statistics \
  clickhouse \
  zookeeper \
  kafka
