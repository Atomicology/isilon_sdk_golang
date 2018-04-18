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

type DrivesDriveFirmwareNodeDrive struct {

	// Numerical representation of this drive's bay.
	Baynum int32 `json:"baynum,omitempty"`

	// This drive's current firmware revision
	CurrentFirmware string `json:"current_firmware,omitempty"`

	// This drive's desired firmware revision.
	DesiredFirmware string `json:"desired_firmware,omitempty"`

	// This drive's device name.
	Devname string `json:"devname,omitempty"`

	// This drive's logical drive number in IFS.
	Lnum int32 `json:"lnum,omitempty"`

	// String representation of this drive's physical location.
	Locnstr string `json:"locnstr,omitempty"`

	// This drive's manufacturer and model.
	Model string `json:"model,omitempty"`
}