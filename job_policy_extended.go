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

type JobPolicyExtended struct {

	// A helpful human-readable description of the impact policy.
	Description string `json:"description"`

	Intervals []JobPolicyInterval `json:"intervals"`

	// The ID of the impact policy.
	Id string `json:"id"`

	// The name of the impact policy.
	Name string `json:"name"`

	// Whether or not this is a read-only system impact policy.
	System bool `json:"system,omitempty"`
}
