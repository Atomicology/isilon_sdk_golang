# SnapshotSnapshotExtended

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | **string** | Alias name to create for this snapshot. If null, remove any alias. | [optional] [default to null]
**Expires** | **int32** | The Unix Epoch time the snapshot will expire and be eligible for automatic deletion. | [optional] [default to null]
**Name** | **string** | The user or system supplied snapshot name. This will be null for snapshots pending delete. | [optional] [default to null]
**Created** | **int32** | The Unix Epoch time the snapshot was created. | [default to null]
**HasLocks** | **bool** | True if the snapshot has one or more locks present see, see the locks subresource of a snapshot for a list of locks. | [default to null]
**Id** | **int32** | The system ID given to the snapshot. This is useful for tracking the status of delete pending snapshots. | [default to null]
**Path** | **string** | The /ifs path snapshotted. | [optional] [default to null]
**PctFilesystem** | **float32** | Percentage of /ifs used for storing this snapshot. | [default to null]
**PctReserve** | **float32** | Percentage of configured snapshot reserved used for storing this snapshot. | [default to null]
**Schedule** | **string** | The name of the schedule used to create this snapshot, if applicable. | [optional] [default to null]
**ShadowBytes** | **int32** | The amount of shadow bytes referred to by this snapshot. | [default to null]
**Size** | **int32** | The amount of storage in bytes used to store this snapshot. | [default to null]
**State** | **string** | Snapshot state. | [default to null]
**TargetId** | **int32** | The ID of the snapshot pointed to if this is an alias. | [optional] [default to null]
**TargetName** | **string** | The name of the snapshot pointed to if this is an alias. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


