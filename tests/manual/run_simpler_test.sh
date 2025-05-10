#!/bin/bash

# Script to run the simpler test for OneLogin Go SDK

echo "Running simpler test..."
cd "$(dirname "$0")" # Move to the directory containing this script

# Export the variables for the test
if [ -z "$ONELOGIN_CLIENT_ID" ] || [ -z "$ONELOGIN_CLIENT_SECRET" ] || [ -z "$ONELOGIN_SUBDOMAIN" ]; then
    echo "Please set the following environment variables:"
    echo "  ONELOGIN_CLIENT_ID"
    echo "  ONELOGIN_CLIENT_SECRET"
    echo "  ONELOGIN_SUBDOMAIN"
    exit 1
fi

# Run the test
cd ../../cmd/simple_test && go run .

echo "Test completed."