#!/bin/bash

# Build clickhouse-operator
# Do not forget to update version

# Source configuration
CUR_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

echo -n "Service..."
"${CUR_DIR}"/go_build_service.sh

echo -n "Consumer..."
"${CUR_DIR}"/go_build_consumer.sh
