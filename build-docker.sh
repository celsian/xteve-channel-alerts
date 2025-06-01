#!/bin/bash
#
# build-docker.sh - Build and publish Docker image for xTeVe Channel Alerts
#
# This script:
# 1. Builds the Docker image for xTeVe Channel Alerts
# 2. Tags it with the current Git commit hash and 'latest'
# 3. Provides an option to push the image to Docker Hub
#
# Usage: ./build-docker.sh [push]
#   - Run without arguments to build locally only
#   - Run with 'push' argument to build and push to Docker Hub

# Exit on any error
set -e

# Configuration - Change these variables to match your Docker Hub username
DOCKER_HUB_USERNAME="yourdockerhubuser"
IMAGE_NAME="xteve-channel-alerts"
FULL_IMAGE_NAME="${DOCKER_HUB_USERNAME}/${IMAGE_NAME}"

echo "=== xTeVe Channel Alerts Docker Build Script ==="
echo

# Get the current git commit hash for versioning
GIT_COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "nogit")
echo "🔍 Current Git commit: ${GIT_COMMIT}"

# Build the Docker image
echo "🔨 Building Docker image..."
docker build -t "${FULL_IMAGE_NAME}:${GIT_COMMIT}" .

# Tag as latest
echo "🏷️ Tagging image as latest..."
docker tag "${FULL_IMAGE_NAME}:${GIT_COMMIT}" "${FULL_IMAGE_NAME}:latest"

echo "✅ Build complete!"
echo "   Created: ${FULL_IMAGE_NAME}:${GIT_COMMIT}"
echo "   Created: ${FULL_IMAGE_NAME}:latest"

# Check if we should push to Docker Hub
if [ "$1" = "push" ]; then
    echo
    echo "🚀 Pushing images to Docker Hub..."
    
    # Check if user is logged in to Docker Hub
    if ! docker info | grep -q "Username"; then
        echo "⚠️ You are not logged in to Docker Hub."
        echo "   Please run 'docker login' first."
        exit 1
    fi
    
    # Push both tags
    docker push "${FULL_IMAGE_NAME}:${GIT_COMMIT}"
    docker push "${FULL_IMAGE_NAME}:latest"
    
    echo "✅ Push complete!"
    echo "   Images are now available at:"
    echo "   - ${FULL_IMAGE_NAME}:${GIT_COMMIT}"
    echo "   - ${FULL_IMAGE_NAME}:latest"
else
    echo
    echo "ℹ️ Images built locally only."
    echo "   To push to Docker Hub, run: ./build-docker.sh push"
fi

echo
echo "=== Done ==="
