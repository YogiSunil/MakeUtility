# APIForge — Project Proposal

## Problem
Software developers frequently need to test REST API endpoints while building and maintaining applications. Testing endpoints individually can become repetitive, especially when a project has many routes. Developers also need to know whether endpoints are responding correctly and how long each request takes.


## Proposed Solution
I plan to develop APIForge, a command-line utility written in Go that allows developers to test multiple REST API endpoints from one configuration file. The program will send HTTP requests, check response status codes, measure response times, and display results in the terminal. It will also save the results to a JSON file for later review.

