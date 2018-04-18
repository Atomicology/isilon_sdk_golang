# QuotaQuota

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Container** | **bool** | If true, SMB shares using the quota directory see the quota thresholds as share size. | [optional] [default to null]
**Enforced** | **bool** | True if the quota provides enforcement, otherwise a accounting quota. | [optional] [default to null]
**Linked** | **bool** | If false and the quota is linked, attempt to unlink. | [optional] [default to null]
**Thresholds** | [***QuotaQuotaThresholds**](QuotaQuotaThresholds.md) |  | [optional] [default to null]
**ThresholdsIncludeOverhead** | **bool** | If true, thresholds apply to data plus filesystem overhead required to store the data (i.e. &#39;physical&#39; usage). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


