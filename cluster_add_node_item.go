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

// Add Node information.
type ClusterAddNodeItem struct {

	// Allow down nodes (Default false).
	AllowDown bool `json:"allow_down,omitempty"`

	// Serial number of this node.
	SerialNumber string `json:"serial_number"`

	// Bypass hardware version checks (Default false).
	SkipHardwareVersionCheck bool `json:"skip_hardware_version_check,omitempty"`
}
