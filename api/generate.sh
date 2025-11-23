#!/bin/bash

protoc --go_out=. --go-grpc_out=. \
    add.proto \
    sub.proto \
    mul.proto \
    div.proto \
    coordinator.proto
