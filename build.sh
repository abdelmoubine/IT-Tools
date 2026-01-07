#!/usr/bin/env bash
set -e
echo "Building IT Support Toolkit (windows/amd64)..."
CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -o ittool_win64.exe main.go
echo "Built ittool_win64.exe"
CGO_ENABLED=1 GOOS=windows GOARCH=386 go build -o ittool_win32.exe main.go
echo "Built ittool_win32.exe"