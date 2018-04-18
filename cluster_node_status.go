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

type ClusterNodeStatus struct {

	// Battery status information.
	Batterystatus *NodeStatusNodeBatterystatus `json:"batterystatus,omitempty"`

	// Storage capacity of this node.
	Capacity []NodeStatusNodeCapacityItem `json:"capacity,omitempty"`

	// CPU status information for this node.
	Cpu *NodeStatusNodeCpu `json:"cpu,omitempty"`

	// Node NVRAM information.
	Nvram *NodeStatusNodeNvram `json:"nvram,omitempty"`

	// Information about this node's power supplies.
	Powersupplies *NodeStatusNodePowersupplies `json:"powersupplies,omitempty"`

	// OneFS release.
	Release string `json:"release,omitempty"`

	// Seconds this node has been online.
	Uptime int32 `json:"uptime,omitempty"`

	// OneFS version.
	Version string `json:"version,omitempty"`
}
