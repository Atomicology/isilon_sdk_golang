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

type NodeStateSmartfailNode struct {

	// This node is dead (dead_devs).
	Dead bool `json:"dead,omitempty"`

	// This node is down (down_devs).
	Down bool `json:"down,omitempty"`

	// Error message, if the HTTP status returned from this node was not 200.
	Error_ string `json:"error,omitempty"`

	// Node ID of the node reporting this information.
	Id int32 `json:"id,omitempty"`

	// This node is in the cluster (all_devs).
	InCluster bool `json:"in_cluster,omitempty"`

	// Logical node number of the node reporting this information.
	Lnn int32 `json:"lnn,omitempty"`

	// This node is readonly (ro_devs).
	Readonly bool `json:"readonly,omitempty"`

	// This node is shutdown readonly (down_devs).
	ShutdownReadonly bool `json:"shutdown_readonly,omitempty"`

	// This node is smartfailed (soft_devs).
	Smartfailed bool `json:"smartfailed,omitempty"`

	// Status of the HTTP response from this node if not 200.  If 200, this field does not appear.
	Status int32 `json:"status,omitempty"`
}
