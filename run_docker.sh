#!/bin/bash
# Script to build and run NoBl in Docker

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${YELLOW}Building and running NoBl in Docker...${NC}"

# Check if Docker is installed
if ! command -v docker &> /dev/null; then
    echo -e "${RED}Error: Docker is not installed or not in PATH${NC}"
    exit 1
fi

# Build the Docker image
echo -e "${YELLOW}Building Docker image...${NC}"
docker build -t nobl .

# Run the Docker container
echo -e "${YELLOW}Running Docker container...${NC}"
docker run -p 8080:8080 --rm nobl

echo -e "${GREEN}NoBl is running at http://localhost:8080${NC}"
