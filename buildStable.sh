#!/bin/bash

set -e

STABLE_BUILD_DIR="./bin/bot_stable"
CURRENT_BUILD_DIR="./bin/bot_current"

cleanup() {
    echo "Cleaning up old builds..."
    rm "$STABLE_BUILD_DIR" "$CURRENT_BUILD_DIR"
}

echo "Finding the latest '-stable' tag..."
LATEST_STABLE_TAG=$(git tag --sort=-creatordate | head -n 2 | tail -n 1)
echo $LATEST_STABLE_TAG

if [[ -z "$LATEST_STABLE_TAG" ]]; then
    echo "Error: No '-stable' tags found."
    exit 1
fi

echo "Latest '-stable' tag found: $LATEST_STABLE_TAG"

echo "Building the stable version..."
git checkout "$LATEST_STABLE_TAG"
go build -o "$STABLE_BUILD_DIR"

echo "Building the current version..."
git checkout -
go build -o "$CURRENT_BUILD_DIR"

# Display build locations
echo "Builds completed:"
echo "Stable build: $STABLE_BUILD_DIR"
echo "Current build: $CURRENT_BUILD_DIR"

# Optionally, run your test script to compare them
# ./your_test_script.sh "$STABLE_BUILD_DIR/app" "$CURRENT_BUILD_DIR/app"
