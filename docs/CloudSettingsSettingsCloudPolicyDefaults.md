# CloudSettingsSettingsCloudPolicyDefaults

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveSnapshotFiles** | **bool** | Specifies if files with snapshots should be archived. | [optional] [default to null]
**Cache** | [***CloudSettingsSettingsCloudPolicyDefaultsCache**](CloudSettingsSettingsCloudPolicyDefaultsCache.md) | Specifies default cloudpool cache settings for new filepool policies. | [optional] [default to null]
**Compression** | **bool** | Specifies if files should be compressed. | [optional] [default to null]
**DataRetention** | **int32** | Specifies the minimum amount of time archived data will be retained in the cloud after deletion. | [optional] [default to null]
**Encryption** | **bool** | Specifies if files should be encrypted. | [optional] [default to null]
**FullBackupRetention** | **int32** | (Used with NDMP backups only.  Not applicable to SyncIQ.)  The minimum amount of time cloud files will be retained after the creation of a full NDMP backup. | [optional] [default to null]
**IncrementalBackupRetention** | **int32** | (Used with SyncIQ and NDMP backups.)  The minimum amount of time cloud files will be retained after the creation of a SyncIQ backup or an incremental NDMP backup. | [optional] [default to null]
**WritebackFrequency** | **int32** | The minimum amount of time to wait before updating cloud data with local changes. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


