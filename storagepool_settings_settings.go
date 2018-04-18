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

type StoragepoolSettingsSettings struct {

	// Automatically manage IO optimization settings on files.
	AutomaticallyManageIoOptimization string `json:"automatically_manage_io_optimization"`

	// Automatically manage protection settings on files.
	AutomaticallyManageProtection string `json:"automatically_manage_protection"`

	// Optimize namespace operations by storing metadata on SSDs.
	GlobalNamespaceAccelerationEnabled bool `json:"global_namespace_acceleration_enabled"`

	// Whether or not namespace operation optimizations are currently in effect.
	GlobalNamespaceAccelerationState string `json:"global_namespace_acceleration_state"`

	// Automatically add additional protection level to all directories.
	ProtectDirectoriesOneLevelHigher bool `json:"protect_directories_one_level_higher"`

	// Spill writes into other pools as needed.
	SpilloverEnabled bool `json:"spillover_enabled"`

	// Target pool for spilled writes.
	SpilloverTarget *StoragepoolSettingsSettingsSpilloverTarget `json:"spillover_target"`

	// The L3 Cache default enabled state. This specifies whether L3 Cache should be enabled on new node pools.
	SsdL3CacheDefaultEnabled bool `json:"ssd_l3_cache_default_enabled"`

	// Controls number of mirrors of QAB blocks to place on SSDs.
	SsdQabMirrors string `json:"ssd_qab_mirrors"`

	// Controls number of mirrors of system B-tree blocks to place on SSDs.
	SsdSystemBtreeMirrors string `json:"ssd_system_btree_mirrors"`

	// Controls number of mirrors of system delta blocks to place on SSDs.
	SsdSystemDeltaMirrors string `json:"ssd_system_delta_mirrors"`

	// Deny writes into reserved virtual hot spare space.
	VirtualHotSpareDenyWrites bool `json:"virtual_hot_spare_deny_writes"`

	// Hide reserved virtual hot spare space from free space counts.
	VirtualHotSpareHideSpare bool `json:"virtual_hot_spare_hide_spare"`

	// The number of drives to reserve for the virtual hot spare, from 0-4.
	VirtualHotSpareLimitDrives int32 `json:"virtual_hot_spare_limit_drives"`

	// The percent space to reserve for the virtual hot spare, from 0-20.
	VirtualHotSpareLimitPercent int32 `json:"virtual_hot_spare_limit_percent"`
}
