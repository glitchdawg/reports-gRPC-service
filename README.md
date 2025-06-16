# 📊 Report Service (gRPC + Go)

A lightweight Go service exposing a gRPC API to generate reports for users, with periodic job execution and in-memory storage.

---

## Features

- gRPC service with two endpoints:
  - `GenerateReport(UserID) → ReportID`
  - `HealthCheck() → Status`
- In-memory report storage (map-based)
- Cron job to auto-generate reports every 10 seconds for predefined users
- gRPC server with logging and concurrency-safe report handling
- Scalable design ready for production extension

---

## 🚀 Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/yourusername/reportservice.git
cd reportservice
```
### 2. Install dependencies
```bash
go mod tidy
```
### 3. Install protoc and plugins
- Download protoc from: https://github.com/protocolbuffers/protobuf/releases

- Install Go plugins:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
Ensure `$GOPATH/bin`is in your `PATH`.

>⚠️ Note: gRPC code is already generated and committed. You only need to run `protoc` if you modify `proto/report.proto`.
### 5. Run the server
```bash
go run main.go
```
You’ll see log output like:

```bash
gRPC server listening on :50051
[2025-06-16T20:10:00+05:30] Generated report for user user-101 (report ID: ...)
```
### 🧪 Testing gRPC Calls
#### Using grpcurl
Install `grpcurl`:
https://github.com/fullstorydev/grpcurl

##### Health Check
```bash
grpcurl -plaintext localhost:50051 report.ReportService/HealthCheck
```
##### Generate Report
```bash
grpcurl -plaintext -d '{"user_id": "test-user"}' localhost:50051 report.ReportService/GenerateReport
```
🔁 Predefined Cron Users
Every 10 seconds, reports are generated for:

```go
predefinedUsers := []string{"user-101", "user-202", "user-303"}
```
Logged to stdout with timestamps.

### Dependencies
- `google.golang.org/grpc` – gRPC transport

- `google.golang.org/protobuf` – Protobuf types

- `github.com/robfig/cron/v3` – Cron job scheduling

- `github.com/google/uuid` – For generating unique report IDs
