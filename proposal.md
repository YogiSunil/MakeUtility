# APIForge — Project Proposal

## Problem
Software developers frequently need to test REST API endpoints while building and maintaining applications. Testing endpoints individually can become repetitive, especially when a project has many routes. Developers also need to know whether endpoints are responding correctly and how long each request takes.


## Proposed Solution
I plan to develop APIForge, a command-line utility written in Go that allows developers to test multiple REST API endpoints from one configuration file. The program will send HTTP requests, check response status codes, measure response times, and display results in the terminal. It will also save the results to a JSON file for later review.


## Main Features

- Load API endpoints from a JSON configuration file.
- Send HTTP requests and record response status codes.
- Measure the response time of each request.
- Identify successful and failed tests.
- Display readable results in the terminal.
- Save test results in a JSON report.
- Support concurrent endpoint testing using Go goroutines.


## Technologies
The project will use Go as the primary programming language. I plan to use the Cobra package for command-line functionality, Go's standard `net/http` package for API requests, and `encoding/json` for reading configuration files and saving reports.


## Testing
I will create table-driven tests to verify endpoint validation and response evaluation. I will also implement a benchmark test to measure the performance of report processing or another repeatable operation.

## Expected Outcome
My goal is to create a practical developer utility that reduces repetitive API testing work and provides useful information about endpoint reliability and response times. I also want to improve my understanding of Go, external packages, error handling, concurrency, and automated testing.

