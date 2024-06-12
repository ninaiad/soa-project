#!/bin/bash

docker compose up \
  main \
  main-db \
  posts \
  posts-db \
  statistics \
  clickhouse \
  zookeeper \
  kafka