## Project Overview

This project is a distributed calculator implemented using a microservices architecture in Go. It uses gRPC for communication between services.

The system is composed of a central **Coordinator** service and several worker services for basic arithmetic operations (Add, Subtract, Multiply, Divide).

A user can send a mathematical expression (e.g., "5 + 3 * 2") to the Coordinator's `Evaluate` endpoint. The Coordinator then performs the following steps:
1.  **Tokenizes** the input string into numbers and operators.
2.  Uses the **Shunting-yard algorithm** to convert the tokens from infix notation to Reverse Polish Notation (RPN).
3.  Builds an **Abstract Syntax Tree (AST)** from the RPN expression.
4.  Traverses the AST and **dispatches** the operations to the corresponding worker microservices (`add`, `sub`, `mul`, `div`) via gRPC calls.
5.  Returns the final result to the client.

### Core Technologies
- **Go**: Primary programming language.
- **gRPC**: Framework for inter-service communication.
- **Protobuf**: Interface Definition Language for gRPC services.
- **Microservices**: The application is split into a Coordinator and multiple arithmetic operation services.

## Building and Running

### Prerequisites
- Go (version 1.24 or later)
- Protoc compiler (for regenerating gRPC code)

### 1. Generate gRPC Code
The gRPC and Protobuf code is generated from the `.proto` files located in the `api/` directory. A convenience script is provided.

To regenerate the code, run:
```bash
./api/generate.sh
```

### 2. Configure Environment
The addresses of the worker services are configured via environment variables. You can copy the example file and modify it if needed.
```bash
cp .env.example .env
```
The default values should work for local development.

### 3. Run the Services
A shell script is provided to start all the necessary servers in the correct order.

To start all services:
```bash
./server.sh
```
This will launch the `add`, `sub`, `mul`, `div`, and `coordinator` servers in the background.

## Development Conventions

### Code Structure
- **`api/`**: Contains all the Protobuf (`.proto`) definitions and the generated Go code for gRPC services and messages.
- **`cmd/`**: Contains the `main` package for each server executable.
- **`internal/`**: Contains the core application logic.
    - **`clients/`**: gRPC client implementations for communicating with the worker services.
    - **`coordinator/`**: The implementation of the Coordinator gRPC service.
    - **`dispatcher/`**: Logic to dispatch operations to the correct worker service.
    - **`parser/`**: The mathematical expression parser, including the tokenizer, shunting-yard algorithm, and AST builder.

### Testing
Tests are located alongside the code they are testing, using the `_test.go` naming convention (e.g., `parser/eval_test.go`). You can run tests for a specific package using `go test`.

For example, to run tests for the parser:
```bash
go test ./internal/parser/...
```
To run all tests in the project:
```bash
go test ./...
```
