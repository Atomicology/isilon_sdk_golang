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

// Node smartfail state.
type NodeStateSmartfailExtended struct {

	// This node is dead (dead_devs).
	Dead bool `json:"dead,omitempty"`

	// This node is down (down_devs).
	Down bool `json:"down,omitempty"`

	// This node is in the cluster (all_devs).
	InCluster bool `json:"in_cluster,omitempty"`

	// This node is readonly (ro_devs).
	Readonly bool `json:"readonly,omitempty"`

	// This node is shutdown readonly (down_devs).
	ShutdownReadonly bool `json:"shutdown_readonly,omitempty"`

	// This node is smartfailed (soft_devs).
	Smartfailed bool `json:"smartfailed,omitempty"`
}
