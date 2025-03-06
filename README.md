# Money Transfer Service

Welcome to our **Money Transfer Service**!

We've built a straightforward backend service that helps you check account balances and move money between accounts. It's all crafted with **Go** and follows modern design principles to keep the code clean, readable, and easy to extend.

---

## What Can You Do With This Service?

This service focuses on two core functions:
1. **View Your Balance** - Check how much money is available in any account
2. **Send Money** - Transfer funds between accounts with ease

The best part? Our transfer system is rock-solid even during high traffic, thanks to Go's powerful concurrency features that prevent conflicts when multiple transfers happen simultaneously.

---

## How We've Structured Things

We've organized the project in layers, creating a clean separation of concerns:

```
[API Layer] ⟶ Receives and responds to HTTP requests
    ↓
[Service Layer] ⟶ Contains all the business logic
    ↓
[Database Layer] ⟶ Manages account data in memory
```

Here's a breakdown of our project folders:

- `cmd/server/` - The starting point of our application
- `internal/api/` - Where all HTTP endpoints live
- `internal/service/` - Houses the transfer and account management logic
- `internal/model/` - Defines data structures like Account and TransferRequest
- `pkg/config/` - Handles configuration loading
- `pkg/database/` - Manages in-memory account data with thread-safe locks
- `test/` - Contains both unit and integration tests
- `scripts/` - Helpful utilities for manual API testing
- `configs/` - Configuration files (server port, initial account balances)

---

## Getting Started

### Step 1: Grab the Code
```bash
git clone https://github.com/Anand078/moneytransfer.git
cd moneytransfer
```

### Step 2: Check the Settings (Optional)
Feel free to customize these files:
- `configs/config.yaml` - Set which port the service runs on
- `configs/initial_balances.json` - Configure starting account balances

### Step 3: Launch the Service
```bash
make run
```
That's it! The server will start and listen for requests on your configured port.

---

## Using the API

Here are the endpoints you can access:

| Method | Endpoint | What It Does |
|--------|----------|--------------|
| GET | `/accounts` | Shows all accounts and their balances |
| POST | `/transfer` | Moves money between accounts |
| GET | `/health` | Simple health check (great for monitoring) |
| GET | `/metrics` | Prometheus metrics (for performance tracking) |

---

## Running Tests

Want to make sure everything's working properly? Just run:

```bash
make test
```

This will execute:
- Unit tests (checking the service logic)
- Integration tests (starting the server and making actual requests)

---

## Configuration Files

| File | Purpose |
|------|---------|
| `configs/config.yaml` | Sets the server port |
| `configs/initial_balances.json` | Defines starting account balances |

---

## Docker Support

Prefer running in a container? We've got you covered:

```bash
make docker
docker run -p 8080:8080 money-transfer:latest
```

---

## Production-Ready Features

We've incorporated several real-world best practices:

- Structured logging with zap
- Configuration-driven setup using YAML and JSON
- Graceful shutdown (handles termination signals properly)
- Health check and metrics endpoints for monitoring
- Comprehensive test coverage (unit + integration)
- Docker support for easy deployment
- Automation through a Makefile

Need help or have questions? Feel free to open an issue on GitHub!