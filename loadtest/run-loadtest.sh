#!/bin/bash

set -e

host="$1"
shift

while ! curl -s "http://${host}:1323/health" > /dev/null; do
  >&2 echo "App is unavailable - waiting"
  sleep 1
done

>&2 echo "App is up - executing load test"
k6 run /mnt/loadtest/loadtest.js
