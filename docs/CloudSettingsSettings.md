# CloudSettingsSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudPolicyDefaults** | [***CloudSettingsSettingsCloudPolicyDefaults**](CloudSettingsSettingsCloudPolicyDefaults.md) | The default filepool policy values for cloudpools. | [optional] [default to null]
**RetryCoefficientArchive** | **string** | Coefficients in the quadratic function for determining the rest period between successive archive attempts. | [optional] [default to null]
**RetryCoefficientCacheInvalidation** | **string** | Coefficients in the quadratic function for determining the rest period between successive cache invalidation attempts. | [optional] [default to null]
**RetryCoefficientCloudGarbageCollection** | **string** | Coefficients in the quadratic function for determining the rest period between successive cloud garbage collection attempts. | [optional] [default to null]
**RetryCoefficientLocalGarbageCollection** | **string** | Coefficients in the quadratic function for determining the rest period between successive local garbage collection attempts. | [optional] [default to null]
**RetryCoefficientReadAhead** | **string** | Coefficients in the quadratic function for determining the rest period between successive read ahead attempts. | [optional] [default to null]
**RetryCoefficientRecall** | **string** | Coefficients in the quadratic function for determining the rest period between successive recall attempts. | [optional] [default to null]
**RetryCoefficientWriteback** | **string** | Coefficients in the quadratic function for determining the rest period between successive writeback attempts. | [optional] [default to null]
**SleepTimeoutArchive** | [***CloudSettingsSettingsSleepTimeoutCloudGarbageCollection**](CloudSettingsSettingsSleepTimeoutCloudGarbageCollection.md) | Amount of time to wait between successive file archive operations. | [optional] [default to null]
**SleepTimeoutCacheInvalidation** | [***CloudSettingsSettingsSleepTimeoutCloudGarbageCollection**](CloudSettingsSettingsSleepTimeoutCloudGarbageCollection.md) | Amount of time to wait between successive file cache_invalidation operations. | [optional] [default to null]
**SleepTimeoutCloudGarbageCollection** | [***CloudSettingsSettingsSleepTimeoutCloudGarbageCollection**](CloudSettingsSettingsSleepTimeoutCloudGarbageCollection.md) | Amount of time to wait between successive file cloud garbage collection operations. | [optional] [default to null]
**SleepTimeoutLocalGarbageCollection** | [***CloudSettingsSettingsSleepTimeoutCloudGarbageCollection**](CloudSettingsSettingsSleepTimeoutCloudGarbageCollection.md) | Amount of time to wait between successive file local garbage collection operations. | [optional] [default to null]
**SleepTimeoutReadAhead** | [***CloudSettingsSettingsSleepTimeoutCloudGarbageCollection**](CloudSettingsSettingsSleepTimeoutCloudGarbageCollection.md) | Amount of time to wait between successive file read ahead operations. | [optional] [default to null]
**SleepTimeoutRecall** | [***CloudSettingsSettingsSleepTimeoutCloudGarbageCollection**](CloudSettingsSettingsSleepTimeoutCloudGarbageCollection.md) | Amount of time to wait between successive file recall operations. | [optional] [default to null]
**SleepTimeoutWriteback** | [***CloudSettingsSettingsSleepTimeoutCloudGarbageCollection**](CloudSettingsSettingsSleepTimeoutCloudGarbageCollection.md) | Amount of time to wait between successive file writeback operations. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


