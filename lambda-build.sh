#!/bin/bash

# Set the name of your Lambda handler script here
handler_script="lambda_function.py"

# Create a temporary directory
temp_dir="lambda_build"
mkdir -p "$temp_dir"

# Install packages from requirements.txt into the temporary directory
pip install -r requirements.txt -t "$temp_dir"

# Copy the Lambda handler script into the temporary directory
cp "$handler_script" "$temp_dir"

# Change to the temporary directory
cd "$temp_dir"

# Zip the contents
zip -r ../lambda_function.zip .

# Change back to the original directory
cd ..

# Clean up the temporary directory
rm -rf "$temp_dir"

echo "Lambda deployment package is ready: lambda_function.zip"
