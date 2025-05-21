#!/bin/bash
# Render.com deployment helper

# This is a simple script that updates the app.json file for deployment
# Usage: ./deploy_render.sh [your-render-username]

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${YELLOW}Preparing NoBl for Render.com deployment...${NC}"

# Ask for username if not provided
if [ -z "$1" ]; then
    read -p "Enter your Render.com username: " USERNAME
else
    USERNAME=$1
fi

# Update the app.json file with the correct repository URL
cat > app.json << EOL
{
  "name": "NoBl Programming Language",
  "description": "Web-based REPL interface for the NoBl programming language",
  "repository": "https://github.com/${USERNAME}/nobl",
  "keywords": ["go", "NoBl", "programming-language", "repl"],
  "buildpacks": [
    {
      "url": "heroku/go"
    }
  ]
}
EOL

# Create render.yaml file
cat > render.yaml << EOL
services:
  - type: web
    name: nobl
    env: go
    buildCommand: go build -o NoBl
    startCommand: ./NoBl --web
    envVars:
      - key: PORT
        value: 10000
    healthCheckPath: /health
EOL

echo -e "${GREEN}Deployment files updated for Render.com!${NC}"
echo -e "${YELLOW}Next steps:${NC}"
echo -e "1. Push your code to GitHub (if you haven't already)"
echo -e "2. Go to Render.com and create a new Web Service"
echo -e "3. Connect to your GitHub repository"
echo -e "4. The service will automatically deploy using the render.yaml configuration"
