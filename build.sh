#!/bin/bash
# Build script for NoBl programming language

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${YELLOW}Building NoBl programming language...${NC}"

# Ensure we're in the project root
cd "$(dirname "$0")"

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo -e "${RED}Error: Go is not installed or not in PATH${NC}"
    exit 1
fi

# Clean any previous build
echo -e "${YELLOW}Cleaning previous builds...${NC}"
rm -f NoBl

# Build the project
echo -e "${YELLOW}Compiling...${NC}"
go build -o NoBl

# Check if build succeeded
if [ $? -eq 0 ]; then
    echo -e "${GREEN}Build successful!${NC}"
    echo -e "${YELLOW}Binary created: ${NC}$(pwd)/NoBl"
    
    # Create a version file with timestamp
    echo "NoBl $(date +'%Y-%m-%d %H:%M:%S')" > version.txt
    echo -e "${YELLOW}Version file created: ${NC}$(pwd)/version.txt"
    
    # Make scripts executable
    chmod +x start_web_repl.sh
    echo -e "${YELLOW}Made start_web_repl.sh executable${NC}"
    
    echo ""
    echo -e "${GREEN}NoBl is ready to use!${NC}"
    echo -e "Run the REPL with: ${YELLOW}./NoBl${NC}"
    echo -e "Run a script with: ${YELLOW}./NoBl path/to/script.nobl${NC}"
    echo -e "Start web server with: ${YELLOW}./NoBl --web${NC} or ${YELLOW}./start_web_repl.sh${NC}"
else
    echo -e "${RED}Build failed${NC}"
    exit 1
fi
