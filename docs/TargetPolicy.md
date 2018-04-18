# TargetPolicy

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailoverFailbackState** | **string** | The condition of this policy with respect to sync failover/failback. | [default to null]
**Id** | **string** | The system ID given to this sync policy. | [default to null]
**LastJobState** | **string** | The state of the last job run for this policy. | [default to null]
**LastSourceCoordinatorIp** | **string** | The IP address from which a SyncIQ coordinator daemon most recently connected to this cluster to update it about the progress of a job for this policy. | [default to null]
**LastUpdateFromSource** | **int32** | The last time this cluster was updated with sync information from the source cluster for this policy, in unix epoch seconds.  Null if no such update has occurred yet. | [optional] [default to null]
**LegacyPolicy** | **bool** | Was this policy defined by a OneFS version earlier than 6.0? (Pre-6.0 policies did not have the target policy concept and canceling from the target side will not be available.) | [default to null]
**Name** | **string** | User-assigned name of this sync policy. | [default to null]
**SourceClusterGuid** | **string** | Unique identifier for the source cluster. | [default to null]
**SourceHost** | **string** | Hostname or IP address of sync source cluster. | [default to null]
**TargetPath** | **string** | Absolute filesystem path on the target cluster for the sync destination. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


