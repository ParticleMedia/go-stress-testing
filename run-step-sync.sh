#!/bin/bash
set -eu

echo "usage: $0 concurrency qps"

go run main.go -c $1 -n 10000 -p curl/din --qps $2 -sync

