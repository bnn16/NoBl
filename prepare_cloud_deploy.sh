#!/bin/bash
# Script to prepare NoBl for general cloud deployment

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${YELLOW}Preparing NoBl for cloud deployment...${NC}"

# Build the application
echo -e "${YELLOW}Building NoBl...${NC}"
./build.sh

# Create necessary directories
mkdir -p deploy
mkdir -p deploy/bin

# Copy the executable
echo -e "${YELLOW}Copying executable...${NC}"
cp NoBl deploy/bin/

# Copy necessary files
echo -e "${YELLOW}Copying configuration files...${NC}"
cp Procfile deploy/
cp app.json deploy/
cp -r examples.nobl deploy/
cp -r test.nobl deploy/

# Create a README for deployment
cat > deploy/README.md << 'EOL'
# NoBl Cloud Deployment Package

This package contains everything needed to deploy the NoBl web REPL to a cloud platform.

## Contents

- `bin/NoBl`: The compiled NoBl executable
- `Procfile`: Configuration for cloud platforms (especially Heroku)
- `app.json`: Application metadata
- Example NoBl programs

## Deployment Instructions

1. Upload this entire directory to your cloud provider
2. Set the PORT environment variable as required by your provider
3. Start the application with: `./bin/NoBl --web`

The application will automatically detect the PORT environment variable and bind to it.
EOL

echo -e "${GREEN}Deployment package prepared in the 'deploy' directory${NC}"
echo -e "${YELLOW}You can now upload the contents of the 'deploy' directory to your cloud provider.${NC}"
echo -e "${YELLOW}Ensure your cloud provider sets the PORT environment variable.${NC}"
