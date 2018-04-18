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

type JobType struct {

	// Whether the job type is enabled and able to run.
	Enabled bool `json:"enabled,omitempty"`

	// Default impact policy of this job type.
	Policy string `json:"policy,omitempty"`

	// Default priority of this job type; lower numbers preempt higher numbers.
	Priority int32 `json:"priority,omitempty"`

	// The schedule on which this job type is queued, if any.
	Schedule string `json:"schedule,omitempty"`
}
