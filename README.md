# Demo Web Server

A simple Go web server designed for testing ECS deployments via Harness CI/CD pipeline.

## Features

- **Root endpoint** (`/`) - serves a welcome message and echoes any path
- **Health check endpoint** (`/health`) - returns HTTP 200 for ECS health checks
- **Environment variable support** for configuration
- **Containerized** with Docker for easy deployment

## Quick Start

### Run Locally

```bash
go run main.go
```

Visit:
- http://localhost:8080/ - Root endpoint
- http://localhost:8080/health - Health check
- http://localhost:8080/any-path - Path echoing

### Run with Docker

```bash
# Build the image
docker build -t demo-web-server .

# Run the container
docker run -p 8080:8080 demo-web-server

# Run with custom version
docker run -p 8080:8080 -e APP_VERSION=1.0.0 demo-web-server
```

## Configuration

| Environment Variable | Default | Description |
|---------------------|---------|-------------|
| `PORT` | `8080` | Server port |
| `APP_VERSION` | `dev` | Application version displayed in responses |

## Endpoints

- `GET /` - Returns welcome message with version
- `GET /health` - Health check endpoint (returns "OK")
- `GET /*` - Echoes the requested path with version

## Deployment

This application is containerized and ready for:
- AWS ECS deployments
- Harness CI/CD pipelines
- Any container orchestration platform

The Docker image uses a multi-stage build for optimal size and includes health check support for production deployments.
