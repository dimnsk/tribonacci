#!/usr/bin/env bash

cd /app/app

shopt -s extglob

go get

echo "Running tests..."
go test
echo "Tests have been completed"
echo "----"

echo "Running application..."
go run !(*_test).go
