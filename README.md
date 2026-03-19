# hurl_demo

A minimal demo showing [hurl](https://hurl.dev) testing an HTTP API.

## Prerequisites

```bash
brew install go
brew install hurl
```

## Run the server

```bash
go run main.go
```

The server starts on `http://localhost:8080`.

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/health` | Health check |
| GET | `/items` | List all items |
| GET | `/items/{id}` | Get item by ID |
| POST | `/items` | Create an item |
| DELETE | `/items/{id}` | Delete an item |

## Run the tests

```bash
hurl --test tests/health.hurl tests/items.hurl
```

## Reports

```bash
# HTML report
hurl --test --report-html report/ tests/*.hurl

# JUnit XML
hurl --test --report-junit report.xml tests/*.hurl
```
