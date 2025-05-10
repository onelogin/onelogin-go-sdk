# Changelog

## [4.2.0]

### Fixed
- Corrected User model field types to match the API documentation (Issue #91):
  - Changed `ManagerADID` from string to int32
  - Changed `ExternalID` from int32 to string
  - Changed `MemberOf` from string to []string
  - Added missing `RoleIDs` field as []int32
  - Updated `CustomAttributes` to use map[string]interface{}
- Fixed custom attribute endpoints to use the correct paths
- Fixed pagination for GetUsers method (Issue #79):
  - Improved query parameter handling for different data types
  - Added proper support for array types in query parameters
  - Added new GetUsersWithPagination method that returns cursor information

### Changed
- Removed region requirement for API authentication
- Simplified API URL construction to use only the subdomain
- Added validation for new field types
- Added GroupID to UserQuery for filtering users by group membership

### Added
- Context support for all API methods (Issue #50):
  - Added WithContext variants of all API methods
  - Allows timeout and cancellation control
  - Maintains backward compatibility with existing methods
- Added PagedResponse structure for pagination metadata
- Added tests for all new functionality

## V4.0.0
- New implementation not compatible with previous versions of the SDK
