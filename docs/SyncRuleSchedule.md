# SyncRuleSchedule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Begin** | **string** | Start time (inclusive) for this schedule, during its specified days.  Format is \&quot;hh:mm\&quot; (24h format hour, and minute).  A null value indicates the beginning of the day (\&quot;00:00\&quot;). | [optional] [default to null]
**End** | **string** | End time (inclusive) for this schedule, during its specified days.  Format is \&quot;hh:mm\&quot; (three-letter weekday name abbreviation, 24h format hour, and minute).  A null value indicates the end of the day (\&quot;23:59\&quot;). | [optional] [default to null]
**Friday** | **bool** | If true, this rule is in effect on Friday.  If false, or unspecified, it is not. | [optional] [default to null]
**Monday** | **bool** | If true, this rule is in effect on Monday.  If false, or unspecified, it is not. | [optional] [default to null]
**Saturday** | **bool** | If true, this rule is in effect on Saturday.  If false, or unspecified, it is not. | [optional] [default to null]
**Sunday** | **bool** | If true, this rule is in effect on Sunday.  If false, or unspecified, it is not. | [optional] [default to null]
**Thursday** | **bool** | If true, this rule is in effect on Thursday.  If false, or unspecified, it is not. | [optional] [default to null]
**Tuesday** | **bool** | If true, this rule is in effect on Tuesday.  If false, or unspecified, it is not. | [optional] [default to null]
**Wednesday** | **bool** | If true, this rule is in effect on Wednesday.  If false, or unspecified, it is not. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


