# APIForge

APIForge is a small command-line tool written in Go that tests REST API endpoints. You list your endpoints in a JSON file and run one command. APIForge sends the requests, checks the status codes, measures how long each one takes, and saves a report.

## Features

- Reads the endpoints to test from a JSON file
- Tests all endpoints at the same time using goroutines
- Checks the expected status code and a maximum response time
- Prints a readable PASS/FAIL table in the terminal
- Saves the results to a JSON report
- Exits with an error code when a test fails, so it can be used in scripts