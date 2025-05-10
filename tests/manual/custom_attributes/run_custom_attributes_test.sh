#!/bin/bash

# Set this to your test credentials
# Or use env variables: ONELOGIN_CLIENT_ID, ONELOGIN_CLIENT_SECRET, ONELOGIN_SUBDOMAIN

cd "$(dirname "$0")"
echo "Running custom attributes example..."
go run custom_attributes_example.go