# Report Service (gRPC + Go)

A simple Go service that exposes a gRPC API to generate reports for users. It stores everything in memory and runs a periodic job every 10 seconds.


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

## Getting Started

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
Make sure `$GOPATH/bin` is in your `PATH` so the plugins can run.
>⚠️ Note: gRPC code is already generated and committed. You only need to run `protoc` if you modify `proto/report.proto`.
### 5. Run the server
```bash
go run main.go
```
Once it’s running, you should see logs like:
```bash
gRPC server listening on :50051
[2025-06-16T20:10:00+05:30] Generated report for user user-101 (report ID: ...)
```
### Testing gRPC Calls
#### Using grpcurl
+ You can test the service using [`grpcurl`](https://github.com/fullstorydev/grpcurl)

##### Health Check
```bash
grpcurl -plaintext localhost:50051 report.ReportService/HealthCheck
```
##### Generate Report
```bash
grpcurl -plaintext -d '{"user_id": "test-user"}' localhost:50051 report.ReportService/GenerateReport
```
Predefined Cron Users
Every 10 seconds, reports are generated for:

```go
predefinedUsers := []string{"user-101", "user-202", "user-303"}
```
These show up in your terminal with timestamps so you know which report was generated when.

### Used Libraries
- `google.golang.org/grpc` – gRPC transport

- `google.golang.org/protobuf` – Protobuf types

- `github.com/robfig/cron/v3` – Cron job scheduling

- `github.com/google/uuid` – For generating unique report IDs

## BONUS

## Q: How would I design and scale this service to handle 10,000 concurrent gRPC requests per second across multiple data centers?


### 1. More copies of the service

* I'd make sure the report generation logic is **stateless** — all data goes into a proper DB or cache.
* Then I'd run **lots of instances** (pods) of the service in Kubernetes or whatever orchestration platform we're using.
* Set up **auto-scaling** based on traffic/load.

---

### 2. Balance the incoming traffic

* Use something like **Envoy** or a cloud-native load balancer to distribute requests to all the service instances.
* Let Kubernetes (or Consul) handle **service discovery**.
* If needed, the client can do round-robin with gRPC’s built-in client-side load balancing.

---

### 3. Replace in-memory storage

* That in-memory map isn’t going to cut it. I’d move reports to **Postgres**, **Redis**, or something similar.
* For high scale, **shard** writes by user/report ID.
* Use Redis to cache frequently accessed reports.

---

### 4. Make it fault-tolerant

* Add **timeouts**, **retries**, and **backoff** logic on every gRPC call.
* Use **readiness/liveness probes** so Kubernetes can restart unhealthy pods.
* If report generation becomes heavy, throw a **queue like Kafka** in the mix.

---

### 5. Multi-region setup

* Deploy the service to **multiple regions** with DB replication.
* Add a **global load balancer** to route users to the closest data center.
* Use **latency-aware routing** for better failover if one region goes down.

---

### 6. Keep an eye on things

* Track **metrics** (requests/sec, error rates) using Prometheus.
* Add **tracing** (OpenTelemetry + Jaeger) so we know where things are slow.
* Ship **structured logs** to something like Datadog or ELK.


## Architecture Summary

![alt text](<Untitled Diagram.drawio.svg>)
