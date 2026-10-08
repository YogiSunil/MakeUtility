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