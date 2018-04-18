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

type NodeDriveconfigNodeAlert struct {

	// Send alerts for unknown drive firmware.
	UnknownFirmware bool `json:"unknown_firmware,omitempty"`

	// Send alerts for unknown drive model.
	UnknownModel bool `json:"unknown_model,omitempty"`
}