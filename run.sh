#!/bin/bash
set -eu

echo "usage: $0 concurrency qps"

go run main.go -c $1 -n 10000 -p curl/send_req.sh --qps $2

