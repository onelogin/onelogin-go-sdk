#!/bin/bash

# Script to run the manual tests for OneLogin Go SDK

# Check if credentials are set
if [ -z "$ONELOGIN_CLIENT_ID" ] || [ -z "$ONELOGIN_CLIENT_SECRET" ] || [ -z "$ONELOGIN_SUBDOMAIN" ]; then
    echo "Please set the following environment variables:"
    echo "  ONELOGIN_CLIENT_ID"
    echo "  ONELOGIN_CLIENT_SECRET"
    echo "  ONELOGIN_SUBDOMAIN"
    exit 1
fi

# Run the test
echo "Running test_user_types.go..."
cd ../../cmd/user_types_test && go run .

echo "Test completed."