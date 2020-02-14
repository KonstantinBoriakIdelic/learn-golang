#!/usr/bin/env bash
# Stop the process if something fails
set -xe

# Fetch dependencies
go get "github.com/gin-gonic/gin"

# Create the application binary that eb uses
GOOS=linux GOARCH=amd64 go build -o bin/application -ldflags="-s -w"