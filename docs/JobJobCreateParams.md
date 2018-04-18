# JobJobCreateParams

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowDup** | **bool** | Whether or not to queue the job if one of the same type is already running or queued. | [optional] [default to null]
**AvscanParams** | [***JobJobAvscanParams**](JobJobAvscanParams.md) |  | [optional] [default to null]
**ChangelistcreateParams** | [***JobJobChangelistcreateParams**](JobJobChangelistcreateParams.md) |  | [optional] [default to null]
**DomainmarkParams** | [***JobJobDomainmarkParams**](JobJobDomainmarkParams.md) |  | [optional] [default to null]
**Paths** | **[]string** | For jobs which take paths, the IFS paths to pass to the job. | [default to null]
**Policy** | **string** | Impact policy of this job instance. | [optional] [default to null]
**PrepairParams** | [***JobJobPrepairParams**](JobJobPrepairParams.md) |  | [optional] [default to null]
**Priority** | **int32** | Priority of this job instance; lower numbers preempt higher numbers. | [optional] [default to null]
**SmartpoolstreeParams** | [***JobJobSmartpoolstreeParams**](JobJobSmartpoolstreeParams.md) |  | [optional] [default to null]
**SnaprevertParams** | [***JobJobSnaprevertParams**](JobJobSnaprevertParams.md) |  | [optional] [default to null]
**Type_** | **string** | Job type to queue. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


