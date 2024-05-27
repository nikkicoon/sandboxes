#!/bin/sh
curl -d '{"BPM":50}' -H "Content-Type: application/json" -X POST http://localhost:8080
