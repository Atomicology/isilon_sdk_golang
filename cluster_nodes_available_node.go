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

type ClusterNodesAvailableNode struct {

	// Node configuration ID.
	ConfigurationId string `json:"configuration_id,omitempty"`

	// Human-readable description giving further detail on status (may be empty)
	Description string `json:"description,omitempty"`

	// Isilon product name.
	Product string `json:"product,omitempty"`

	// Serial number of this node.
	SerialNumber string `json:"serial_number,omitempty"`

	// Availability of the node.
	Status string `json:"status,omitempty"`

	// OneFS build version running on the node.
	Version string `json:"version,omitempty"`
}
