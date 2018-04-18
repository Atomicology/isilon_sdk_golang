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

// Apply hardening on the cluster.
type HardeningApplyItem struct {

	// Hardening profile.
	Profile string `json:"profile,omitempty"`

	// Option to only generate and display a report on current cluster configuration with respect to the expected configuation required to apply hardening. If his option is set to true, hardening is not applied after the report is displayed. By default, this option is false.
	Report bool `json:"report,omitempty"`
}