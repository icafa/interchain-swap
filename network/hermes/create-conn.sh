#!/bin/bash
set -e

# Load shell variables
. ./network/hermes/variables.sh

### Configure the clients and connection
echo "Initiating connection handshake..."
hermes -c ./network/hermes/config.toml create connection test-1 test-2
sleep 5 

hermes -c ./network/hermes/config.toml create channel --port-a transfer --port-b transfer test-1 connection-0
