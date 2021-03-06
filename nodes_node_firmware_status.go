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

// The firmware status for the node.
type NodesNodeFirmwareStatus struct {

	// List of the firmware status for hardware components on the node.
	Devices []ClusterFirmwareStatusNodeDevice `json:"devices,omitempty"`

	// Node is unavailable.
	NodeUnavailable bool `json:"node_unavailable,omitempty"`

	// List of the firmware binary information for the installed firmware package.
	Package_ []ClusterFirmwareStatusNodePackageItem `json:"package,omitempty"`
}
