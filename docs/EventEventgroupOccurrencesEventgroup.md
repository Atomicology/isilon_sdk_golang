# EventEventgroupOccurrencesEventgroup

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Causes** | [**[][]string**](array.md) | List of eventgroup IDs that may be causes of this occurrence. | [optional] [default to null]
**Channels** | **[]string** | List of channels to which alerts are configured for this eventgroup | [optional] [default to null]
**EventCount** | **int32** | Number of events linked to this eventgroup. | [optional] [default to null]
**EventgroupInstance** | **string** | Unique identifier of eventgroup instance. | [optional] [default to null]
**Id** | **string** | Same as eventgroup_instance. | [optional] [default to null]
**Ignore** | **bool** | True if eventgroup is marked as &#39;ignore&#39;. | [optional] [default to null]
**IgnoreTime** | **int32** | Time eventgroup was marked as &#39;ignore&#39;. | [optional] [default to null]
**LastEvent** | **int32** | Time the most recent event was added to this eventgroup. | [optional] [default to null]
**ResolveTime** | **int32** | When the eventgroup became resolved. | [optional] [default to null]
**Resolved** | **bool** | True if eventgroup is resolved. | [optional] [default to null]
**Resolver** | **string** | &#39;USER&#39; if the eventgroup was marked resolved via PAPI &lt;event_instance&gt; if eventgroup was marked resolved as a result of an event. | [optional] [default to null]
**Sequence** | **int32** | XXX description needed. | [optional] [default to null]
**Severity** | **string** | Event Group severity. | [optional] [default to null]
**Specifier** | [***Empty**](Empty.md) | A collection of parameters defined per eventgroup_id. | [optional] [default to null]
**TimeNoticed** | **int32** | Time of first event linked to this eventgroup. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


