#!/bin/bash

# Script to run the manual tests for the OneLogin SDK fixes

echo "=== Running Manual Tests for OneLogin SDK Fixes ==="

# Check if credentials are set
if [ -z "$ONELOGIN_CLIENT_ID" ] || [ -z "$ONELOGIN_CLIENT_SECRET" ] || [ -z "$ONELOGIN_SUBDOMAIN" ]; then
    echo "Please set the following environment variables:"
    echo "  ONELOGIN_CLIENT_ID"
    echo "  ONELOGIN_CLIENT_SECRET"
    echo "  ONELOGIN_SUBDOMAIN"
    exit 1
fi

# Run the test
cd "$(dirname "$0")" # Move to the directory containing this script
go run fixes_manual_test.go

echo "=== Manual Tests Completed ==="