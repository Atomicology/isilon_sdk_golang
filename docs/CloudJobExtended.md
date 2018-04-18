# CloudJobExtended

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletionTime** | **int32** | The time at which the job was completed (if applicable) | [optional] [default to null]
**CreateTime** | **int32** | The time at which the job was created | [optional] [default to null]
**Description** | **string** | A brief description of the job contents | [optional] [default to null]
**EffectiveState** | **string** | The effective state of the job (e.g,. the combination of operation_state and job_state) | [optional] [default to null]
**Files** | [***CloudJobFiles**](CloudJobFiles.md) | The files and filters addressed by this job | [optional] [default to null]
**Id** | **int32** | The job&#39;s ID | [optional] [default to null]
**JobEngineJob** | [***CloudJobJobEngineJob**](CloudJobJobEngineJob.md) | Information about the related job engine job if there is one | [optional] [default to null]
**JobState** | **string** | The current state of the job | [optional] [default to null]
**OperationState** | **string** | The current state of the operation associated with the job | [optional] [default to null]
**StateChangeTime** | **int32** | The last time at which the job state changed | [optional] [default to null]
**Type_** | **string** | The type of cloud action to be performed by this job. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


