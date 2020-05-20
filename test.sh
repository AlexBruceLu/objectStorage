#!/usr/bin/env bash

#ifconfig enp0s31f6:1 10.29.1.1/16
#ifconfig enp0s31f6:2 10.29.1.2/16
#ifconfig enp0s31f6:3 10.29.1.3/16
#ifconfig enp0s31f6:4 10.29.1.4/16
#ifconfig enp0s31f6:5 10.29.1.5/16
#ifconfig enp0s31f6:6 10.29.1.6/16
#ifconfig enp0s31f6:7 10.29.2.1/16
#ifconfig enp0s31f6:8 10.29.2.2/16

export RABBITMQ_SERVER=amqp://test:test@127.0.0.1:5672
LISTEN_ADDRESS=10.29.1.1:12345 STORAGE_ROOT=/tmp/1 go run cmd/object-storage-data/main.go &
LISTEN_ADDRESS=10.29.1.2:12345 STORAGE_ROOT=/tmp/2 go run cmd/object-storage-data/main.go &
LISTEN_ADDRESS=10.29.1.3:12345 STORAGE_ROOT=/tmp/3 go run cmd/object-storage-data/main.go &
LISTEN_ADDRESS=10.29.1.4:12345 STORAGE_ROOT=/tmp/4 go run cmd/object-storage-data/main.go &
LISTEN_ADDRESS=10.29.1.5:12345 STORAGE_ROOT=/tmp/5 go run cmd/object-storage-data/main.go &
LISTEN_ADDRESS=10.29.1.6:12345 STORAGE_ROOT=/tmp/6 go run cmd/object-storage-data/main.go &

LISTEN_ADDRESS=10.29.2.1:12345  go run cmd/object-storage-api/main.go &
LISTEN_ADDRESS=10.29.2.2:12345  go run cmd/object-storage-api/main.go &