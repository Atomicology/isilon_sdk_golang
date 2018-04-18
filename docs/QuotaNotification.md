# QuotaNotification

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionAlert** | **bool** | Send alert when rule matches. | [optional] [default to null]
**ActionEmailAddress** | **string** | Email a specific email address when rule matches. | [optional] [default to null]
**ActionEmailOwner** | **bool** | Email quota domain owner when rule matches. | [optional] [default to null]
**EmailTemplate** | **string** | Path of optional /ifs template file used for email actions. | [optional] [default to null]
**Holdoff** | **int32** | Time to wait between detections for rules triggered by user actions. | [optional] [default to null]
**Schedule** | **string** | Schedule for rules that repeatedly notify. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


