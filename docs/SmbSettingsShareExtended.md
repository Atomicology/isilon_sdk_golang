# SmbSettingsShareExtended

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessBasedEnumeration** | **bool** | Only enumerate files and folders the requesting user has access to. | [optional] [default to null]
**AccessBasedEnumerationRootOnly** | **bool** | Access-based enumeration on only the root directory of the share. | [optional] [default to null]
**AllowDeleteReadonly** | **bool** | Allow deletion of read-only files in the share. | [optional] [default to null]
**AllowExecuteAlways** | **bool** | Allows users to execute files they have read rights for. | [optional] [default to null]
**CaTimeout** | **int32** | Persistent open timeout for the share. | [optional] [default to null]
**CaWriteIntegrity** | **string** | Specify the level of write-integrity on continuously available shares. | [optional] [default to null]
**ChangeNotify** | **string** | Specify level of change notification alerts on the share. | [optional] [default to null]
**CreatePermissions** | **string** | Set the create permissions for new files and directories in share. | [optional] [default to null]
**CscPolicy** | **string** | Client-side caching policy for the shares. | [optional] [default to null]
**DirectoryCreateMask** | **int32** | Unix umask or mode bits. | [optional] [default to null]
**DirectoryCreateMode** | **int32** | Unix umask or mode bits. | [optional] [default to null]
**FileCreateMask** | **int32** | Unix umask or mode bits. | [optional] [default to null]
**FileCreateMode** | **int32** | Unix umask or mode bits. | [optional] [default to null]
**FileFilterExtensions** | **[]string** | Specifies the list of file extensions. | [optional] [default to null]
**FileFilterType** | **string** | Specifies if filter list is for deny or allow. Default is deny. | [optional] [default to null]
**FileFilteringEnabled** | **bool** | Enables file filtering on the share. | [optional] [default to null]
**HideDotFiles** | **bool** | Hide files and directories that begin with a period &#39;.&#39;. | [optional] [default to null]
**HostAcl** | **[]string** | An ACL expressing which hosts are allowed access. A deny clause must be the final entry. | [optional] [default to null]
**ImpersonateGuest** | **string** | Specify the condition in which user access is done as the guest account. | [optional] [default to null]
**ImpersonateUser** | **string** | User account to be used as guest account. | [optional] [default to null]
**MangleByteStart** | **int32** | Specifies the wchar_t starting point for automatic byte mangling. | [optional] [default to null]
**MangleMap** | **[]string** | Character mangle map. | [optional] [default to null]
**NtfsAclSupport** | **bool** | Support NTFS ACLs on files and directories. | [optional] [default to null]
**Oplocks** | **bool** | Allow oplock requests. | [optional] [default to null]
**StrictCaLockout** | **bool** | Specifies if persistent opens would do strict lockout on the share. | [optional] [default to null]
**StrictFlush** | **bool** | Handle SMB flush operations. | [optional] [default to null]
**StrictLocking** | **bool** | Specifies whether byte range locks contend against SMB I/O. | [optional] [default to null]
**Zone** | **string** | Name of the access zone in which to update settings | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


