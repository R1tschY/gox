#!/bin/bash

# Copyright Richard Liebscher 2026
# SPDX-License-Identifier: MIT

# Exit immediately if a command exits with a non-zero status.
set -e

# Ensure we're in the project root
cd "$(dirname "$0")"

# Check if the working directory is clean
if [ -n "$(git status --porcelain)" ]; then
    echo "Error: Working directory is not clean. Commit or stash your changes before releasing."
    exit 1
fi

# Run tests
echo "Running tests..."
go test ./...

# Get new version from argument or prompt
NEW_VERSION=$1
if [ -z "$NEW_VERSION" ]; then
    read -r -p "Enter new version (e.g., v0.1.0): " NEW_VERSION
fi

# Basic semantic versioning validation (with 'v' prefix)
if [[ ! "$NEW_VERSION" =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    echo "Error: Version must follow semantic versioning with a 'v' prefix (e.g., v0.1.0)."
    exit 1
fi

# Check if tag already exists
if git rev-parse "$NEW_VERSION" >/dev/null 2>&1; then
    echo "Error: Tag $NEW_VERSION already exists."
    exit 1
fi

# Create tag
echo "Creating tag $NEW_VERSION..."
git tag -a "$NEW_VERSION" -m "Release $NEW_VERSION"

# Ask for confirmation before pushing
if [ -z "$1" ]; then
    # Interactive mode confirmation
    read -r -p "Push tag $NEW_VERSION to origin? (y/N): " CONFIRM
    if [ "$CONFIRM" != "y" ] && [ "$CONFIRM" != "Y" ]; then
        echo "Release cancelled. Tag $NEW_VERSION created locally."
        echo "To push manually: git push origin $NEW_VERSION"
        exit 0
    fi
fi

# Push tag
echo "Pushing tag $NEW_VERSION to origin..."
git push origin "$NEW_VERSION"

# Trigger indexing on pkg.go.dev via proxy.golang.org
echo "Triggering indexing on pkg.go.dev..."
MODULE_NAME=$(go list -m)
curl -s "https://proxy.golang.org/${MODULE_NAME}/@v/${NEW_VERSION}.info" > /dev/null
