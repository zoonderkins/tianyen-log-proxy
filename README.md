### README.md

```markdown
# Project Name

This project provides a simple logging API using the Gin framework in Go. Logs are forwarded to a Loki endpoint with CORS support enabled.

## Prerequisites

- Go 1.22
- Make

## Setup

Follow these steps to set up and run the project.

### 1. Clone the Repository

```sh
git clone https://gitlab.com/t9963/log-proxy-server
cd log-proxy-server
```

### 2. Initialize the Project

Initialize the Go module for the project:

```sh
go mod init gitlab.com/t9963/log-proxy-server
```

### 3. Install Dependencies

Install the required Go packages:

```sh
go mod tidy
go get github.com/gin-gonic/gin
go get github.com/go-resty/resty/v2
go get github.com/joho/godotenv
go get github.com/gin-contrib/cors
```

### 4. Create .env File

Create a `.env` file in the root directory of the project and add your configuration values:

```plaintext
LOKI_ENDPOINT=https://loki.xxxx.com/loki/api/v1/push
BASIC_AUTH_USERNAME=xxxx
BASIC_AUTH_PASSWORD=xxxx
```

### 5. Build the Project

Use the `Makefile` to build the project for different platforms:

```sh
make build
```

This command will create binaries for Linux (64-bit) and Mac OS (ARM64) in the `bin` directory.

### 6. Run the Application

To run the application, use the built binary or simply run it with `go run` for development:

```sh
go run main.go
```

Or, if you want to run the production-ready binary:

```sh
./bin/proxy-linux-amd64   # For Linux
./bin/proxy-darwin-arm64  # For Mac OS
```

### 7. Usage

The logging API is available at `/log`. You can send a POST request with the log data in JSON format.

Example:

```json
POST /log
Content-Type: application/json

{
  "level": "info",
  "logger": "meta-tag-demosite",
  "service": "pd0001",
  "application": "lucky-dart",
  "message": "This is a test log message",
  "user_agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36",
  "locationurlPath": "/home",
  "userId": "user123",
  "additionalInfo": {
    "key1": "value1",
    "key2": "value2"
  }
}
```

## Cleaning Up

To clean up the build artifacts, run:

```sh
make clean
```
