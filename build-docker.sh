#!/bin/bash
#
# build-docker.sh - Build and publish Docker image for xTeVe Channel Alerts
#
# This script:
# 1. Builds a multi-architecture Docker image for xTeVe Channel Alerts (amd64 and arm64)
# 2. Tags it with the current Git commit hash and 'latest'
# 3. Provides an option to push the image to Docker Hub
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

# Set up buildx for multi-architecture builds if not already set up
if ! docker buildx inspect multi-arch-builder &>/dev/null; then
    echo "🔧 Setting up multi-architecture builder..."
    docker buildx create --name multi-arch-builder --use
fi

# Ensure the builder is ready
docker buildx inspect multi-arch-builder --bootstrap

# Check if we should push to Docker Hub
if [ "$1" = "push" ]; then
    echo
    echo "🚀 Building and pushing multi-architecture images to Docker Hub..."

    # Check if user is logged in to Docker Hub using a more reliable method
    if ! docker buildx ls &>/dev/null; then
        echo "⚠️ You are not logged in to Docker Hub."
        echo "   Please run 'docker login' first."
        exit 1
    fi

    # Build and push both architectures
    echo "🔨 Building for amd64 and arm64 platforms..."
    docker buildx build \
        --platform linux/amd64,linux/arm64 \
        --tag "${FULL_IMAGE_NAME}:${GIT_COMMIT}" \
        --tag "${FULL_IMAGE_NAME}:latest" \
        --push \
        .

    echo "✅ Build and push complete!"
    echo "   Images are now available at:"
    echo "   - ${FULL_IMAGE_NAME}:${GIT_COMMIT}"
    echo "   - ${FULL_IMAGE_NAME}:latest"
else
    echo "🔨 Building multi-architecture images locally..."
    
    # For local builds, we can only build for the current architecture
    # but we'll configure it as a multi-arch build for testing
    docker buildx build \
        --platform linux/amd64,linux/arm64 \
        --tag "${FULL_IMAGE_NAME}:${GIT_COMMIT}" \
        --tag "${FULL_IMAGE_NAME}:latest" \
        --load \
        .

    echo "✅ Build complete!"
    echo "   Created: ${FULL_IMAGE_NAME}:${GIT_COMMIT}"
    echo "   Created: ${FULL_IMAGE_NAME}:latest"
    echo
    echo "ℹ️ Images built locally only."
    echo "   To push to Docker Hub, run: ./build-docker.sh push"
fi

echo
echo "=== Done ==="
