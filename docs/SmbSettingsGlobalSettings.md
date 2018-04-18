# SmbSettingsGlobalSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessBasedShareEnum** | **bool** | Only enumerate files and folders the requesting user has access to. | [optional] [default to null]
**AuditFileshare** | **string** | Specify level of file share audit events to log. | [optional] [default to null]
**AuditGlobalSacl** | [**[]SmbSettingsGlobalSettingsAuditGlobalSaclItem**](SmbSettingsGlobalSettingsAuditGlobalSaclItem.md) | Specifies a list of permissions to audit. | [optional] [default to null]
**AuditLogon** | **string** | Specify the level of logon audit events to log. | [optional] [default to null]
**DotSnapAccessibleChild** | **bool** | Allow access to .snapshot directories in share subdirectories. | [optional] [default to null]
**DotSnapAccessibleRoot** | **bool** | Allow access to the .snapshot directory in the root of the share. | [optional] [default to null]
**DotSnapVisibleChild** | **bool** | Show .snapshot directories in share subdirectories. | [optional] [default to null]
**DotSnapVisibleRoot** | **bool** | Show the .snapshot directory in the root of a share. | [optional] [default to null]
**EnableSecuritySignatures** | **bool** | Indicates whether the server supports signed SMB packets. | [optional] [default to null]
**GuestUser** | **string** | Specifies the fully-qualified user to use for guest access. | [optional] [default to null]
**IgnoreEas** | **bool** | Specify whether to ignore EAs on files. | [optional] [default to null]
**OnefsCpuMultiplier** | **int32** | Specify the number of OneFS driver worker threads per CPU. | [optional] [default to null]
**OnefsNumWorkers** | **int32** | Set the maximum number of OneFS driver worker threads. | [optional] [default to null]
**RequireSecuritySignatures** | **bool** | Indicates whether the server requires signed SMB packets. | [optional] [default to null]
**ServerSideCopy** | **bool** | Enable Server Side Copy. | [optional] [default to null]
**ServerString** | **string** | Provides a description of the server. | [optional] [default to null]
**Service** | **bool** | Specify whether service is enabled. | [optional] [default to null]
**SrvCpuMultiplier** | **int32** | Specify the number of SRV service worker threads per CPU. | [optional] [default to null]
**SrvNumWorkers** | **int32** | Set the maximum number of SRV service worker threads. | [optional] [default to null]
**SupportMultichannel** | **bool** | Support multichannel. | [optional] [default to null]
**SupportNetbios** | **bool** | Support NetBIOS. | [optional] [default to null]
**SupportSmb2** | **bool** | Support the SMB2 protocol on the server. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


