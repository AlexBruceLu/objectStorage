#!/bin/sh

LISTEN_ADDRESS=10.29.1.1:12345 STORAGE_ROOT=storage_data/1 go run ./cmd/data-service/data-server.go &

LISTEN_ADDRESS=10.29.1.2:12345 STORAGE_ROOT=storage_data/2 go run ./cmd/data-service/data-server.go &

LISTEN_ADDRESS=10.29.1.3:12345 STORAGE_ROOT=storage_data/3 go run ./cmd/data-service/data-server.go &

LISTEN_ADDRESS=10.29.1.4:12345 STORAGE_ROOT=storage_data/4 go run ./cmd/data-service/data-server.go &

LISTEN_ADDRESS=10.29.1.5:12345 STORAGE_ROOT=storage_data/5 go run ./cmd/data-service/data-server.go &

LISTEN_ADDRESS=10.29.1.6:12345 STORAGE_ROOT=storage_data/6 go run ./cmd/data-service/data-server.go &

LISTEN_ADDRESS=10.29.2.1:12345 go run ./cmd/api-service/api-server.go &

LISTEN_ADDRESS=10.29.2.2:12345 go run ./cmd/api-service/api-server.go &
