# SnapshotScheduleExtended

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | **string** | Alias name to create for each snapshot. | [optional] [default to null]
**Duration** | **int32** | Time in seconds added to creation time to construction expiration time. | [optional] [default to null]
**Id** | **int32** | The system ID given to the schedule. | [optional] [default to null]
**Name** | **string** | The schedule name. | [optional] [default to null]
**NextRun** | **int32** | Unix Epoch time of next snapshot to be created. | [optional] [default to null]
**NextSnapshot** | **string** | Formatted name (see pattern) of next snapshot to be created. | [optional] [default to null]
**Path** | **string** | The /ifs path snapshotted. | [optional] [default to null]
**Pattern** | **string** | Pattern expanded with strftime to create snapshot name. | [optional] [default to null]
**Schedule** | **string** | The isidate compatible natural language description of the schedule. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


