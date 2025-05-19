# Changelog

## [4.5.1]

### Fixed
- Fixed issue with empty arrays in Role updates by ensuring empty arrays are properly serialized in JSON
- Added documentation example for removing all users from a role

## [4.5.0]

### Improved
- Added context support for roles API methods:
  - `CreateRoleWithContext`
  - `GetRolesWithContext`
  - `GetRoleByIDWithContext`
  - `UpdateRoleWithContext`
  - `DeleteRoleWithContext`
- Improved `UpdateRole` function to use pointer parameter for consistency with other API methods
- Removed unused queryParams parameter from DeleteRole function
- Updated Role model documentation to match the actual structure

## [4.4.0]

### Added
- Complete implementation of User Mappings API endpoints:
  - Added strongly typed methods for all User Mappings operations
  - Added proper model structures for User Mappings
  - Added query parameters support for filtering User Mappings
  - Added examples for User Mappings API usage
- Improved handling of API responses that return minimal data (ID only)
- Added comprehensive test suite for User Mappings functionality

### Fixed
- Fixed HTTP response handling for endpoints that return only an ID
- Improved error reporting for API responses

## [4.3.0]

### Added
- Helper methods for creating and updating custom attributes with proper payload structure:
  - Added `CreateCustomAttribute(name, shortname string)` method
  - Added `UpdateCustomAttribute(id int, name, shortname string)` method
  - These methods properly wrap parameters in the required "user_field" object
- Added example code for custom attributes with create, read, update, delete examples

### Fixed
- Fixed custom attribute creation/update by properly wrapping parameters in the "user_field" object
- Fixed handling of HTTP 204 No Content responses for operations like delete that return no content on success

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
