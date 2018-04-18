# EventSettingsSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HeartbeatInterval** | **string** | Interval between heartbeat events. \&quot;daily\&quot;, \&quot;weekly\&quot;, or \&quot;monthly\&quot;. | [optional] [default to null]
**Maintenance** | [***EventSettingsSettingsMaintenance**](EventSettingsSettingsMaintenance.md) | Specifies start and duration of maintenance period during which no alerts will be sent for new eventgroups. | [optional] [default to null]
**RetentionDays** | **int32** | Retention period in days | [optional] [default to null]
**StorageLimit** | **int32** | Storage limit in megabytes per terabyte of available storage | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


