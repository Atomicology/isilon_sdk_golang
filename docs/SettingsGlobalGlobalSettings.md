# SettingsGlobalGlobalSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllocRetries** | **int32** | Specifies the number of times to retry an ID allocation before failing. | [optional] [default to null]
**GidRangeEnabled** | **bool** | If true, allocates GIDs from a fixed range. | [optional] [default to null]
**GidRangeMax** | **int32** | Specifies the ending number for a fixed range from which GIDs are allocated. | [optional] [default to null]
**GidRangeMin** | **int32** | Specifies the starting number for a fixed range from which GIDs are allocated. | [optional] [default to null]
**GidRangeNext** | **int32** | Specifies the next GID to allocate. | [optional] [default to null]
**GroupUid** | **int32** | Specifies the user iD for a group when requested by the kernel. | [optional] [default to null]
**LoadProviders** | **[]string** | Specifies which providers are loaded by the authentication daemon (lsassd). | [optional] [default to null]
**MinMappedRid** | **int32** | Starts the RID in the local domain to map a UID and a GID. | [optional] [default to null]
**NullGid** | **int32** | Specifies an alternative GID when the kernel is unable to retrieve a GID for a persona. | [optional] [default to null]
**NullUid** | **int32** | Specifies an alternative UID when the kernel is unable to retrieve a UID for a persona. | [optional] [default to null]
**OnDiskIdentity** | **string** | Specifies the type of identity that is stored on disk. | [optional] [default to null]
**RpcBlockTime** | **int32** | Specifies the minimum amount of time in milliseconds to wait before performing an oprestart. | [optional] [default to null]
**RpcMaxRequests** | **int32** | Specifies the maximum number of outstanding RPC requests. | [optional] [default to null]
**RpcTimeout** | **int32** | Specifies the maximum amount of time in seconds to wait for an idmap response. | [optional] [default to null]
**SendNtlmv2** | **bool** | If true, sends NTLMv2 responses. | [optional] [default to null]
**SpaceReplacement** | **string** | Specifies the space replacement character. | [optional] [default to null]
**SystemGidThreshold** | **int32** | Specifies the minimum GID to attempt to look up in the idmap database. | [optional] [default to null]
**SystemUidThreshold** | **int32** | Specifies the minimum UID to attempt to look up in the idmap database. | [optional] [default to null]
**UidRangeEnabled** | **bool** | If true, allocates UIDs from a fixed range. | [optional] [default to null]
**UidRangeMax** | **int32** | Specifies the ending number for a fixed range from which UIDs are allocated. | [optional] [default to null]
**UidRangeMin** | **int32** | Specifies the starting number for a fixed range from which UIDs are allocated. | [optional] [default to null]
**UidRangeNext** | **int32** | Specifies the next UID to allocate. | [optional] [default to null]
**UnknownGid** | **int32** | Specifies the GID for the unknown (anonymous) group. | [optional] [default to null]
**UnknownUid** | **int32** | Specifies the UID for the unknown (anonymous) user. | [optional] [default to null]
**UserObjectCacheSize** | **int32** | Specifies the maximum size (in bytes) of the security object cache in the authentication daemon. | [optional] [default to null]
**Workgroup** | **string** | Specifies the NetBIOS workgroup or domain. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


