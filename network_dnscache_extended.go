/*
 * Isilon SDK
 *
 * Isilon SDK - Language bindings for the OneFS API
 *
 * API version: 5
 * Contact: sdk@isilon.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package isi_sdk_8_1_0

type NetworkDnscacheExtended struct {

	// DNS cache entry limit
	CacheEntryLimit int32 `json:"cache_entry_limit,omitempty"`

	// Timeout value for calls made to other nodes in the cluster
	ClusterTimeout int32 `json:"cluster_timeout,omitempty"`

	// Timeout value for calls made to the dns resolvers
	DnsTimeout int32 `json:"dns_timeout,omitempty"`

	// Lead time to refresh cache entries nearing expiration
	EagerRefresh int32 `json:"eager_refresh,omitempty"`

	// Deltas for checking cbind cluster health
	TestpingDelta int32 `json:"testping_delta,omitempty"`

	// Upper bound on ttl for cache hits
	TtlMaxNoerror int32 `json:"ttl_max_noerror,omitempty"`

	// Upper bound on ttl for nxdomain
	TtlMaxNxdomain int32 `json:"ttl_max_nxdomain,omitempty"`

	// Upper bound on ttl for non-nxdomain failures
	TtlMaxOther int32 `json:"ttl_max_other,omitempty"`

	// Upper bound on ttl for server failures
	TtlMaxServfail int32 `json:"ttl_max_servfail,omitempty"`

	// Lower bound on ttl for cache hits
	TtlMinNoerror int32 `json:"ttl_min_noerror,omitempty"`

	// Lower bound on ttl for nxdomain
	TtlMinNxdomain int32 `json:"ttl_min_nxdomain,omitempty"`

	// Lower bound on ttl for non-nxdomain failures
	TtlMinOther int32 `json:"ttl_min_other,omitempty"`

	// Lower bound on ttl for server failures
	TtlMinServfail int32 `json:"ttl_min_servfail,omitempty"`
}
