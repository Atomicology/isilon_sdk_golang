# StatisticsHistoryStat

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Devid** | **int32** | Devid of node of statistic or 0 for cluster scoped statistics. | [default to null]
**Error_** | **string** | Key specific error string, if applicable. | [optional] [default to null]
**ErrorCode** | **int32** | Key specific error number, if applicable. | [optional] [default to null]
**Key** | **string** | Key name of statistic. | [default to null]
**Resolution** | **int32** | The interval for which these results were figured (averaged against.) | [default to null]
**Values** | [**[]StatisticsHistoryStatValue**](StatisticsHistoryStatValue.md) | Time-series values. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


