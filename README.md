# APIForge

APIForge is a small command-line tool written in Go that tests REST API endpoints. You list your endpoints in a JSON file and run one command. APIForge sends the requests, checks the status codes, measures how long each one takes, and saves a report.

## Features

- Reads the endpoints to test from a JSON file
- Tests all endpoints at the same time using goroutines
- Checks the expected status code and a maximum response time
- Prints a readable PASS/FAIL table in the terminal
- Saves the results to a JSON report
- Exits with an error code when a test fails, so it can be used in scripts

## Installation

You need Go 1.21 or newer.

```
git clone https://github.com/YogiSunil/MakeUtility.git
cd MakeUtility
go build -o apiforge .
```

This creates an `apiforge` program in the project folder.

## Usage

```
apiforge test endpoints.json
```

| Flag | Short | Default | What it does |
| --- | --- | --- | --- |
| `--output` | `-o` | `results.json` | Where to save the JSON report |
| `--timeout` | | `10s` | How long to wait for each endpoint |
| `--verbose` | `-v` | off | Also show the URL and status code for every endpoint |

If any endpoint fails, APIForge exits with status 1.

## Endpoints file

```json
{
  "endpoints": [
    {
      "name": "users",
      "url": "https://jsonplaceholder.typicode.com/users",
      "method": "GET",
      "expected_status": 200,
      "max_latency_ms": 1500
    }
  ]
}
```

| Field | Required | Description |
| --- | --- | --- |
| `url` | yes | Full http or https address |
| `name` | no | Label used in the results. Defaults to the URL |
| `method` | no | GET, POST, PUT, PATCH, DELETE or HEAD. Defaults to GET |
| `expected_status` | no | Status code that counts as a pass. Defaults to 200 |
| `max_latency_ms` | no | The test fails if the response takes longer. 0 means no limit |

## Example output

```
$ apiforge test endpoints.json

APIForge Results
------------------------------------------------
GET    httpbin get              PASS    195ms
GET    users                    PASS     92ms
GET    missing page             PASS    195ms
------------------------------------------------
Passed: 3 | Failed: 0

Report saved to results.json
```