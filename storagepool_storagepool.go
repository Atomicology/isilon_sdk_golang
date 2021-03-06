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

type StoragepoolStoragepool struct {

	// Indicates if disabling L3 is possible.
	CanDisableL3 bool `json:"can_disable_l3,omitempty"`

	// Indicates if enabling L3 is possible. L3 cannot be enabled if there are unprovisioned drives.
	CanEnableL3 bool `json:"can_enable_l3,omitempty"`

	// The names or IDs of the pool's children.
	Children []string `json:"children,omitempty"`

	// An array of containing any health issues with this pool.  If the pool is healthy, the list is empty.  Only appears on 'nodepool' type storagepools.
	HealthFlags []string `json:"health_flags,omitempty"`

	// The system ID given to the pool.
	Id int32 `json:"id"`

	// Use SSDs in this node pool for L3 cache.
	L3 bool `json:"l3,omitempty"`

	// 'storage' if the 'l3' option is disabled. If the l3 option is enabled, 'migrating' if any SSDs in this node pool have not yet been migrated to L3. If all SSDs have been migrated, 'l3'.
	L3Status string `json:"l3_status,omitempty"`

	// The nodes that are part of this pool.
	Lnns []int32 `json:"lnns"`

	// Whether or not the node pool is manually managed.
	Manual bool `json:"manual,omitempty"`

	// The pool name.
	Name string `json:"name"`

	// The underlying protection policy.
	ProtectionPolicy string `json:"protection_policy,omitempty"`

	// The pool type.
	Type_ string `json:"type"`

	// Total pool usage.
	Usage *StoragepoolTierUsage `json:"usage"`
}
