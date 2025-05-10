# Changelog

## [Unreleased]

### Fixed
- Corrected User model field types to match the API documentation:
  - Changed `ManagerADID` from string to int32
  - Changed `ExternalID` from int32 to string
  - Changed `MemberOf` from string to []string
  - Added missing `RoleIDs` field as []int32
  - Updated `CustomAttributes` to use map[string]interface{}
- Fixed custom attribute endpoints to use the correct paths

### Changed
- Removed region requirement for API authentication
- Simplified API URL construction to use only the subdomain
- Added validation for new field types

### Added
- Test utilities for verifying these fixes

## V4.0.0
- New implementation not compatible with previous versions of the SDK
