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

type AntivirusPolicyExtended struct {

	// A description for the policy.
	Description string `json:"description,omitempty"`

	// Whether the policy is enabled.
	Enabled bool `json:"enabled,omitempty"`

	// Forces the scan to run regardless of whether the files were recently scanned.
	ForceRun bool `json:"force_run,omitempty"`

	Impact string `json:"impact,omitempty"`

	// The name of the policy.
	Name string `json:"name,omitempty"`

	// Paths to include in the scan.
	Paths []string `json:"paths,omitempty"`

	// The depth to recurse in directories.  The default of -1 gives unlimited recursion.
	RecursionDepth int32 `json:"recursion_depth,omitempty"`

	Schedule string `json:"schedule,omitempty"`

	// A unique identifier for the policy.
	Id string `json:"id,omitempty"`

	// The epoch time of the last run of this policy.
	LastRun int32 `json:"last_run,omitempty"`
}
