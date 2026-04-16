# AI-Ops Desktop - Backend

Backend service for AI-Ops Desktop, built with Go.

## Technology Stack
- **Language:** Go 1.26+
- **Logging:** `log/slog` (structured logging)
- **Metrics:** `github.com/shirou/gopsutil/v3`
- **HTTP Router:** Standard library `net/http`
- **Configuration:** Environment variables + config struct

## Development

```bash
# Install dependencies
go mod download

# Run the server
go run ./cmd/api

# Build the binary
go build -o aiops-desktop ./cmd/api

# Run tests
go test ./...
```

## Project Structure

```
backend/
├── cmd/
│   └── api/
│       └── main.go           # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go         # Configuration management
│   ├── httpapi/
│   │   └── middleware.go     # HTTP middleware (CORS)
│   ├── logger/
│   │   └── logger.go         # Structured logger setup
│   └── metrics/
│       ├── collector.go      # System metrics collection
│       ├── collector_loop.go # Periodic collection loop
│       ├── handler.go        # HTTP handlers
│       ├── stream_handler.go # SSE streaming handler
│       ├── store.go          # In-memory metrics store
│       └── types.go          # Data models and types
├── go.mod
└── go.sum
```

## Configuration

Environment variables:
- `PORT`: Server port (default: "8080")
- `LOG_LEVEL`: Logging level (default: "info")
- `ENVIRONMENT`: Environment name (default: "development")

## API Endpoints

### Health Check
```
GET /health
Response: { "status": "healthy" }
```

### Metrics
```
GET /metrics
Response: Latest metrics snapshot (JSON)
```

### Metrics Stream (SSE)
```
GET /metrics/stream
Response: Server-Sent Events stream of metrics snapshots
```

## Components

### Metrics Collector
Collects real-time system metrics using `gopsutil`:
- CPU usage percentage
- Memory usage in bytes
- Disk usage in bytes
- Network TX/RX bytes

### Collector Loop
Runs periodic metrics collection:
- Configurable interval (default: 2 seconds)
- Context-based cancellation
- Error handling and logging

### Metrics Store
In-memory storage for metrics:
- Latest snapshot retrieval
- Configurable history retention
- Thread-safe operations

### SSE Streaming
Real-time metrics delivery:
- Standard SSE protocol
- Automatic client disconnection detection
- 2-second update interval

## Development Notes

- The server starts the metrics collector loop in a background goroutine
- All HTTP handlers use structured logging
- CORS is enabled for development
- The collector continues running even if no clients are connected
- Store maintains a configurable history of snapshots

## Building for Production

When building for production with Electron:

1. Use a predictable port or IPC for communication
2. Implement graceful shutdown
3. Configure appropriate log levels
4. Consider adding health check endpoints for monitoring

## Troubleshooting

**Port Already in Use:** Change the `PORT` environment variable

**Permission Errors:** Some metrics collection may require elevated privileges on certain systems

**High CPU Usage:** Adjust the collector interval in `main.go`

**Memory Growth:** The store has a configurable maximum history to prevent unbounded growth

## Future Enhancements

- SQLite persistence for alerts and actions (Sprint 6)
- Alert rule engine (Sprint 3)
- Safe action executor (Sprint 4)
- AI provider integration (Sprint 5)

---

**Last Updated:** April 16, 2026
**Status:** Sprint 2 Complete - Live Metrics Pipeline
**Go Version:** 1.26+
