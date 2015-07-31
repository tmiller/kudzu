#!/bin/bash
env GOOS=linux GOARCH=amd64 go build -o build/kudzu.linux_amd64 *.go
env GOOS=darwin GOARCH=amd64 go build -o build/kudzu.darwin_amd64 *.go
