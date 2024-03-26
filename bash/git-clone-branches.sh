#!/bin/bash

# Repository URL
REPO_URL="https://github.com/example/repo.git"

# List of branches to clone
BRANCHES=(branch1 branch2 branch3)

# Loop through each branch in the list
for BRANCH in "${BRANCHES[@]}"; do
    echo "Cloning branch $BRANCH..."
    git clone --branch "$BRANCH" --single-branch "$REPO_URL" "$BRANCH"
done

echo "Cloning completed."
