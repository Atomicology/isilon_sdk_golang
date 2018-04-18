# NetworkDnscacheSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheEntryLimit** | **int32** | DNS cache entry limit | [default to null]
**ClusterTimeout** | **int32** | Timeout value for calls made to other nodes in the cluster | [default to null]
**DnsTimeout** | **int32** | Timeout value for calls made to the dns resolvers | [default to null]
**EagerRefresh** | **int32** | Lead time to refresh cache entries nearing expiration | [default to null]
**TestpingDelta** | **int32** | Deltas for checking cbind cluster health | [default to null]
**TtlMaxNoerror** | **int32** | Upper bound on ttl for cache hits | [default to null]
**TtlMaxNxdomain** | **int32** | Upper bound on ttl for nxdomain | [default to null]
**TtlMaxOther** | **int32** | Upper bound on ttl for non-nxdomain failures | [default to null]
**TtlMaxServfail** | **int32** | Upper bound on ttl for server failures | [default to null]
**TtlMinNoerror** | **int32** | Lower bound on ttl for cache hits | [default to null]
**TtlMinNxdomain** | **int32** | Lower bound on ttl for nxdomain | [default to null]
**TtlMinOther** | **int32** | Lower bound on ttl for non-nxdomain failures | [default to null]
**TtlMinServfail** | **int32** | Lower bound on ttl for server failures | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


