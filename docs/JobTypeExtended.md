# JobTypeExtended

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Whether the job type is enabled and able to run. | [default to null]
**Policy** | **string** | Default impact policy of this job type. | [default to null]
**Priority** | **int32** | Default priority of this job type; lower numbers preempt higher numbers. | [default to null]
**Schedule** | **string** | The schedule on which this job type is queued, if any. | [optional] [default to null]
**AllowMultipleInstances** | **bool** | Whether multiple instances of this job type may run simultaneously. | [default to null]
**Description** | **string** | Brief description of the job type. | [default to null]
**ExclusionSet** | **string** | The set(s) of mutually-exclusive job types to which this job belongs.  No job in this set may run with any other job in this set.  Obsolete; this value will always be an empty string, as exclusion sets are no longer a job type property. | [default to null]
**Hidden** | **bool** | Whether this job type is normally visible in the UI. | [default to null]
**Id** | **string** | Job type ID. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


