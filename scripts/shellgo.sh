#!/bin/bash
set -e  # exit on any error

echo "Removing vendor directory..."
rm -rf vendor

echo "Vendoring dependencies..."
go mod vendor

echo "Cleaning build cache..."
go clean -cache

echo "Tidying go.mod..."
go mod tidy

echo "Downloading modules..."
go mod download

echo "Verifying modules..."
go mod verify

echo "Removing old Swagger docs..."
rm -rf docs/

echo "Generating new Swagger docs..."
swag init

echo "Running the application with 'serve' command..."
go run main.go serve