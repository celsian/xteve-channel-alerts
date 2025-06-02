#!/bin/bash
#
# build-docker.sh - Build and publish Docker image for xTeVe Channel Alerts
#
# This script:
# 1. Builds the Docker image for xTeVe Channel Alerts for the current architecture
# 2. Creates a manifest list with support for multiple architectures
# 3. Tags it with the current Git commit hash and 'latest'
# 4. Provides an option to push the image and manifest to Docker Hub
#
# Usage: ./build-docker.sh [push]
#   - Run without arguments to build locally only
#   - Run with 'push' argument to build and push to Docker Hub

# Exit on any error
set -e

# Configuration - Change these variables to match your Docker Hub username
DOCKER_HUB_USERNAME="celsian"
IMAGE_NAME="xteve-channel-alerts"
FULL_IMAGE_NAME="${DOCKER_HUB_USERNAME}/${IMAGE_NAME}"

echo "=== xTeVe Channel Alerts Docker Build Script ==="
echo

# Get the current git commit hash for versioning
GIT_COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "nogit")
echo "🔍 Current Git commit: ${GIT_COMMIT}"

# Build the Docker image for the current architecture
echo "🔨 Building Docker image for current architecture..."
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
    echo "🚀 Pushing image to Docker Hub..."

    # Check if user is logged in to Docker Hub
    if ! docker info | grep -q "Username"; then
        echo "⚠️ You are not logged in to Docker Hub."
        echo "   Please run 'docker login' first."
        exit 1
    fi

    # Push the image
    docker push "${FULL_IMAGE_NAME}:${GIT_COMMIT}"
    docker push "${FULL_IMAGE_NAME}:latest"
    
    echo "📋 Creating multi-architecture manifest..."
    # Create a manifest list
    docker manifest create "${FULL_IMAGE_NAME}:latest" "${FULL_IMAGE_NAME}:latest" --amend
    
    # Annotate the manifest for different architectures
    echo "🔧 Annotating manifest for multiple architectures..."
    docker manifest annotate "${FULL_IMAGE_NAME}:latest" "${FULL_IMAGE_NAME}:latest" --os linux --arch arm64
    docker manifest annotate "${FULL_IMAGE_NAME}:latest" "${FULL_IMAGE_NAME}:latest" --os linux --arch amd64
    
    # Push the manifest
    echo "📤 Pushing manifest to Docker Hub..."
    docker manifest push "${FULL_IMAGE_NAME}:latest"

    echo "✅ Push complete!"
    echo "   Multi-architecture image is now available at:"
    echo "   - ${FULL_IMAGE_NAME}:latest (arm64, amd64)"
    echo "   Single-architecture image is available at:"
    echo "   - ${FULL_IMAGE_NAME}:${GIT_COMMIT}"
else
    echo
    echo "ℹ️ Images built locally only."
    echo "   To push to Docker Hub, run: ./build-docker.sh push"
fi

echo
echo "=== Done ==="
