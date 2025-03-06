# Technical Write-Up: Money Transfer Service

## Overview

This project is a backend service for transferring money between accounts. It is built using Go and follows clean architecture principles with clearly separated layers: API, service, and database.

## Architecture

The system is organized into the following layers:
- **API Layer:** Handles HTTP requests.
- **Service Layer:** Contains business logic for money transfers.
- **Database Layer:** Uses an in-memory store with thread-safe operations.

## Key Features

- **Concurrency Safety:** Uses sync.RWMutex to ensure safe concurrent access.
- **Structured Logging:** Implements logging using zap.
- **Configuration Management:** Loads settings from YAML and JSON files.
- **Graceful Shutdown:** Handles OS signals for clean shutdowns.
- **Observability:** Provides health and metrics endpoints.
- **Testing:** Includes unit tests and integration tests.
- **Docker Support:** Containerized for easy deployment.
- **Automation:** Common tasks are automated using a Makefile.

## Testing Strategy

Unit tests cover individual components (service and database logic), while integration tests validate end-to-end functionality of the service.

## Production Readiness

This project is designed to be:
- Easy to maintain and extend.
- Reliable in a concurrent environment.
- Ready for deployment via Docker and CI pipelines.

## Future Enhancements

- Add persistent storage (e.g., PostgreSQL)
- Implement advanced monitoring and tracing
- Introduce authentication and rate limiting
