#!/bin/bash

# Build the current version
go build -o bin/bot_current main.go

# Checkout and build the last stable version
last_stable_tag=$(git describe --tags $(git rev-list --tags='*-stable' --max-count=1))
git checkout $last_stable_tag

# Build the last stable version
go build -o bin/bot_stable main.go

# Return to the current version
git checkout -

# Confirm the binaries are ready
echo "Built current version as ./bin/bot_current and stable version as ./bin/bot_stable"
