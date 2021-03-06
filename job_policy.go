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

type JobPolicy struct {

	// A helpful human-readable description of the impact policy.
	Description string `json:"description,omitempty"`

	Intervals []JobPolicyInterval `json:"intervals,omitempty"`
}
