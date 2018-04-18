# FsaResultExtended

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pinned** | **bool** | True if the result is pinned to prevent automatic removal. | [default to null]
**BeginTime** | **int32** | Unix Epoch time of start of results collection job. | [default to null]
**ContentPath** | **string** | Path to results database. | [optional] [default to null]
**DeleteLink** | **string** | Resource to call with DELETE to remove results.. | [optional] [default to null]
**EndTime** | **int32** | Unix Epoch time of end of results collection job. | [default to null]
**FsaState** | **string** | State of the result set. | [default to null]
**Id** | **int32** | The system generated result set ID. | [default to null]
**JobState** | **[]string** | State information about the FSA Job. | [default to null]
**PropertiesLink** | **string** | Resource to call to get result properties. | [default to null]
**Size** | **int32** | Size of the result set database in bytes. | [default to null]
**Version** | **int32** | FSA version used to create result set. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


