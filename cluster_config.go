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

// General cluster information.
type ClusterConfig struct {

	// Customer configurable description.
	Description string `json:"description"`

	Devices []ClusterConfigDevice `json:"devices"`

	// Default encoding.
	Encoding string `json:"encoding"`

	// Cluster GUID.
	Guid string `json:"guid"`

	// If true, the local node is in a group with quorum: It is communicating, storing, and protecting data normally with other nodes in its group.  Under normal circumstances, every node in the cluster is part of one group.
	HasQuorum bool `json:"has_quorum"`

	// If true, the cluster is in compliance mode.  Compliance mode clusters do not allow root access and have stricter WORM (Write Once Read Many) features for retaining data in compliance with U.S. Securities and Exchange Commission rule 17a-4.
	IsCompliance bool `json:"is_compliance"`

	// true if the cluster is deployed on a desktop VMWorkstation
	IsVirtual bool `json:"is_virtual"`

	// true if this is a vOneFS cluster on an ESXi
	IsVonefs bool `json:"is_vonefs"`

	// Node join mode: 'manual' or 'secure'.
	JoinMode string `json:"join_mode"`

	// Device ID of the queried node.
	LocalDevid int32 `json:"local_devid"`

	// Device logical node number of the queried node.
	LocalLnn int32 `json:"local_lnn"`

	// Device serial number of the queried node.
	LocalSerial string `json:"local_serial"`

	// Cluster name.
	Name string `json:"name"`

	// 
	OnefsVersion *ClusterConfigOnefsVersion `json:"onefs_version,omitempty"`

	// The cluster timezone settings.
	Timezone *ClusterConfigTimezone `json:"timezone,omitempty"`

	UpgradeType string `json:"upgrade_type,omitempty"`
}
