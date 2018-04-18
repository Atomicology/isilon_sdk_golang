# Zone

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlternateSystemProvider** | **string** | Specifies an alternate system provider. | [optional] [default to null]
**AuthProviders** | **[]string** | Specifies the list of authentication providers available on this access zone. | [optional] [default to null]
**CacheEntryExpiry** | **int32** | Specifies amount of time in seconds to cache a user/group. | [optional] [default to null]
**CreatePath** | **bool** | Determines if a path is created when a path does not exist. | [optional] [default to null]
**ForceOverlap** | **bool** | Allow for overlapping base path. | [optional] [default to null]
**HomeDirectoryUmask** | **int32** | Specifies the permissions set on automatically created user home directories. | [optional] [default to null]
**IfsRestricted** | [**[]AuthAccessAccessItemFileGroup**](AuthAccessAccessItemFileGroup.md) | Specifies a list of users and groups that have read and write access to /ifs. | [optional] [default to null]
**MapUntrusted** | **string** | Maps untrusted domains to this NetBIOS domain during authentication. | [optional] [default to null]
**Name** | **string** | Specifies the access zone name. | [optional] [default to null]
**NegativeCacheEntryExpiry** | **int32** | Specifies number of seconds the negative cache entry is valid. | [optional] [default to null]
**NetbiosName** | **string** | Specifies the NetBIOS name. | [optional] [default to null]
**Path** | **string** | Specifies the access zone base directory path. | [optional] [default to null]
**SkeletonDirectory** | **string** | Specifies the skeleton directory that is used for user home directories. | [optional] [default to null]
**SystemProvider** | **string** | Specifies the system provider for the access zone. | [optional] [default to null]
**UserMappingRules** | **[]string** | Specifies the current ID mapping rules. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


