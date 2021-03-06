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

// View information about nodes during an upgrade, rollback, or pre-upgrade assessment.
type ClusterNodesExtended struct {

	// List of detailed info of nodes which are part of the current upgrade
	Nodes []ClusterNodes `json:"nodes,omitempty"`

	// Total number of nodes the upgrade framework is aware of and was able to collect info for at this point.
	Total int32 `json:"total,omitempty"`
}
