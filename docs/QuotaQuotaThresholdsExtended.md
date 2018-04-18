# QuotaQuotaThresholdsExtended

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Advisory** | **int32** | Usage bytes at which notifications will be sent but writes will not be denied. | [optional] [default to null]
**Hard** | **int32** | Usage bytes at which further writes will be denied. | [optional] [default to null]
**Soft** | **int32** | Usage bytes at which notifications will be sent and soft grace time will be started. | [optional] [default to null]
**SoftGrace** | **int32** | Time in seconds after which the soft threshold has been hit before writes will be denied. | [optional] [default to null]
**AdvisoryExceeded** | **bool** | True if the advisory threshold has been hit. | [optional] [default to null]
**AdvisoryLastExceeded** | **int32** | Time at which advisory threshold was hit. | [optional] [default to null]
**HardExceeded** | **bool** | True if the hard threshold has been hit. | [optional] [default to null]
**HardLastExceeded** | **int32** | Time at which hard threshold was hit. | [optional] [default to null]
**SoftExceeded** | **bool** | True if the soft threshold has been hit. | [optional] [default to null]
**SoftLastExceeded** | **int32** | Time at which soft threshold was hit | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


