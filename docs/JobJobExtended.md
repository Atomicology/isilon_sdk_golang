# JobJobExtended

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ControlState** | **string** | State to which the job is transitioning; if control_state is identical to state, the job&#39;s state is stable. | [optional] [default to null]
**CreateTime** | **int32** | The time the job was queued, in seconds since the epoch. | [default to null]
**CurrentPhase** | **int32** | The current phase of the job. | [optional] [default to null]
**Description** | **string** | A text representation of the job. | [optional] [default to null]
**EndTime** | **int32** | The time the job ended, in seconds since the Epoch. | [optional] [default to null]
**Id** | **int32** | The ID of the job. | [default to null]
**Impact** | **string** | The current impact of the job. | [default to null]
**Participants** | **[]int32** | The set of devids working on the job. | [optional] [default to null]
**Paths** | **[]string** | Paths for which the job was queued. | [optional] [default to null]
**Policy** | **string** | Current impact policy of the job. | [default to null]
**Priority** | **int32** | Current priority of the job; lower numbers preempt higher numbers. | [default to null]
**Progress** | **string** | A text representation of the job&#39;s progress. | [optional] [default to null]
**RetriesRemaining** | **int32** | The number of retries remaining if the job fails. | [default to null]
**RunningTime** | **int32** | The number of seconds the job has executed. | [optional] [default to null]
**StartTime** | **int32** | The time the job started, in seconds since the Epoch. | [optional] [default to null]
**State** | **string** | Current state of the job. | [default to null]
**TotalPhases** | **int32** | The total number of phases of the job type. | [default to null]
**Type_** | **string** | The job type. | [default to null]
**WaitingOn** | **int32** | The ID of a job for which this job is waiting. | [optional] [default to null]
**WaitingReason** | **string** | The reason the job is waiting. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


