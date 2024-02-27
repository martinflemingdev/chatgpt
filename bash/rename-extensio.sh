#!/bin/bash

# Navigate to the directory containing the .txt files
# cd /path/to/directory

# Loop through all .txt files in the current directory
for file in *.txt; do
    # Check if the file variable contains a valid filename
    # This is necessary to handle the case where there are no .txt files
    if [[ -f "$file" ]]; then
        # Use mv to rename the file by changing its extension to .json
        mv -- "$file" "${file%.txt}.json"
    fi
done
