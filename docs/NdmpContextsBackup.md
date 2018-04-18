# NdmpContextsBackup

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextExpirationTime** | **int32** | Context expiration time | [optional] [default to null]
**ContextId** | **string** | Context ID | [optional] [default to null]
**Id** | **string** | Unique display id. | [optional] [default to null]
**Path** | **string** | The directory of the backup. This is not applicable to restore contexts. | [optional] [default to null]
**Sessions** | [**[]NdmpContextsBackupSession**](NdmpContextsBackupSession.md) |  | [optional] [default to null]
**Snapid** | **int32** | Snapshot ID reserved for the context. This is not applicable to restore contexts. | [optional] [default to null]
**StartTime** | **int32** | Context creation time | [optional] [default to null]
**Status** | **string** | Whether the context is active. | [optional] [default to null]
**TotalSessions** | **int32** | The number of sessions in the context | [optional] [default to null]
**Type_** | **string** | Type of context; It can be bre, backup, and restore | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


