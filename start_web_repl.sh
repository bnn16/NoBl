#!/bin/bash
# Run NoBl web REPL

# Default port
PORT=8080

# Check if a port was provided
if [ ! -z "$1" ]; then
  PORT=$1
fi

# Build the NoBl executable if it doesn't exist
if [ ! -f "./NoBl" ]; then
  echo "Building NoBl..."
  go build
fi

# Start the web server
echo "Starting NoBl web REPL on http://localhost:$PORT"
./NoBl --web $PORT
