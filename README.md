# gRPC Backoff Retry Example

This repository demonstrates a simple gRPC implementation with a **client** that performs exponential backoff retries for transient errors and a **server** that sometimes responds with errors to trigger retry logic. The project uses the [cenkalti/backoff](https://github.com/cenkalti/backoff) library for managing retries.

## Features

- **Client**:
  - Implements exponential backoff with retryable error handling.
  - Handles `Unavailable` and `DeadlineExceeded` errors as retryable.
- **Server**:
  - Randomly simulates transient errors (`Unavailable`).
  - Introduces random delays to test timeout and retry mechanisms.

---

## Getting Started

### Prerequisites

- [Go](https://go.dev/) (1.19 or later)
- Protocol Buffers Compiler (`protoc`)
- gRPC and related Go libraries:
  - `google.golang.org/grpc`
  - `github.com/cenkalti/backoff/v4`

### Installation

1. Clone the repository:
   ```bash
   $ git clone https://github.com/your-repo/grpc-backoff-retry-example.git
   $ cd grpc-backoff-retry-example
   ```

### Install Dependencies
```sh
$ go mod tidy
```