# Manual Tests for OneLogin Go SDK

This directory contains manual tests for verifying functionality of the SDK.

## Running Tests

### Test User Types

This test verifies the fixes for:
1. Field type corrections (ManagerADID, ExternalID, MemberOf, RoleIDs)
2. Custom attributes implementation
3. Custom attributes API path

To run the test:

```bash
# Set your OneLogin API credentials
export ONELOGIN_CLIENT_ID="your_client_id"
export ONELOGIN_CLIENT_SECRET="your_client_secret"
export ONELOGIN_SUBDOMAIN="your_subdomain"

# Run the test
go run test_user_types.go
```

## Expected Output

The test will:
1. Create a test user with the fixed field types
2. Fetch custom attributes using the corrected API path
3. List users and display the field types for verification

If all field types are correct and API paths work as expected, the test has passed.

## Notes

- These tests require valid OneLogin API credentials
- Some operations (like creating a user) may incur charges or impact your OneLogin environment
- You may want to modify the tests to use a sandbox environment if available