#!/bin/bash

# Array of Git repository URLs
REPOS=("https://github.com/user/repo1.git" "https://github.com/user/repo2.git" ...)

# Directory where repositories will be cloned
CLONE_DIR="/path/to/clone/directory"

# Cloning each repository
for repo in "${REPOS[@]}"
do
    git clone "$repo" "$CLONE_DIR/$(basename "$repo" .git)"
done

# Navigating to the directory containing all the repositories
cd "$CLONE_DIR"

# Updating each repository
for dir in */
do
    cd "$dir"
    git pull
    cd ..
done
