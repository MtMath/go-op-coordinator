#!/bin/bash

echo "Starting AddServer..."
go run ./cmd/add_server &
ADD_PID=$!

echo "Starting SubServer..."
go run ./cmd/sub_server &
SUB_PID=$!

echo "Starting MulServer..."
go run ./cmd/mul_server &
MUL_PID=$!

echo "Starting DivServer..."
go run ./cmd/div_server &
DIV_PID=$!

sleep 1

echo "Starting CoordinatorServer..."
go run ./cmd/coordinator_server &
COORD_PID=$!

echo "Everything is running!"
echo "CTRL+C to Stop"
wait
