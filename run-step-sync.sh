#!/bin/bash
set -eu

echo "usage: $0 urlfile concurrency qps"

go run main.go -n 10000 -p $1 -c $2 --qps $3 -sync

