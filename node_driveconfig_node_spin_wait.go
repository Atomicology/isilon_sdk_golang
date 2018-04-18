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

type NodeDriveconfigNodeSpinWait struct {

	// Seconds to wait between enabling a bay and checking for an inserted drive.
	CheckDrive int32 `json:"check_drive,omitempty"`

	// Seconds to wait between enabling a bay and enabling another bay.
	Stagger int32 `json:"stagger,omitempty"`
}
