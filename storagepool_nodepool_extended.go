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

type StoragepoolNodepoolExtended struct {

	// Indicates if enabling L3 is possible. L3 cannot be enabled if there are unprovisioned drives.
	CanEnableL3 bool `json:"can_enable_l3"`

	// An array of containing any health issues with this pool.  If the pool is healthy, the list is empty.
	HealthFlags []string `json:"health_flags,omitempty"`

	// The system ID given to the node pool.
	Id int32 `json:"id"`

	// Use SSDs in this node pool for L3 cache.
	L3 bool `json:"l3"`

	// 'storage' if the 'l3' option is disabled. If the l3 option is enabled, 'migrating' if any SSDs in this node pool have not yet been migrated to L3. If all SSDs have been migrated, 'l3'.
	L3Status string `json:"l3_status"`

	// The nodes that are part of this node pool.
	Lnns []int32 `json:"lnns"`

	// Whether or not the node pool is manually managed.
	Manual bool `json:"manual"`

	// The node pool name.
	Name string `json:"name"`

	// The underlying protection policy.
	ProtectionPolicy string `json:"protection_policy,omitempty"`

	// The name (if named) or system ID of the node pool's tier, if it is in a tier. Otherwise null.
	Tier string `json:"tier,omitempty"`

	// Total pool usage.
	Usage *StoragepoolTierUsage `json:"usage"`
}