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

type ClusterFirmwareStatusNodePackageItem struct {

	// Valid checksum string for binary. One of the following: 'ok', 'bad', 'file_missing', or 'na'
	Checksum string `json:"checksum,omitempty"`

	// The name of the firmware binary.
	Firmware string `json:"firmware,omitempty"`

	// The version string for the binary.
	Version string `json:"version,omitempty"`
}