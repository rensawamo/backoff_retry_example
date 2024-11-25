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
- [Buf](https://buf.build/)
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

### Generate gRPC code from .proto files:
```sh
$ make buf-gen
```

### Running the Server
1. Start the gRPC server:
```sh
$ go run server/main.go
```

2. The server will start listening on localhost:50051.

### Running the Client
1. Run the gRPC client:

```sh
$ go run client/main.go
```

2. The client sends a request to the server and retries upon encountering transient errors (Unavailable or DeadlineExceeded).

3. Logs will show the retry attempts and the final result.

### Example Output

Server
```sh
gRPC server is running on port 50051...
Received request: Hello, gRPC with Backoff!
Simulating transient error: Unavailable
Received request: Hello, gRPC with Backoff!
Simulating delay: 2s
Sending response: Hello from gRPC server!
```

Client
```sh
Retryable error: rpc error: code = Unavailable desc = simulated transient error. Retrying...
Retryable error: rpc error: code = DeadlineExceeded desc = operation timed out. Retrying...
Response received: Reply:"Hello from gRPC server!"
```



