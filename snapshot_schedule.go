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

type SnapshotSchedule struct {

	// Alias name to create for each snapshot.
	Alias string `json:"alias,omitempty"`

	// Time in seconds added to creation time to construction expiration time.
	Duration int32 `json:"duration,omitempty"`

	// The schedule name.
	Name string `json:"name,omitempty"`

	// The /ifs path snapshotted.
	Path string `json:"path,omitempty"`

	// Pattern expanded with strftime to create snapshot names.
	Pattern string `json:"pattern,omitempty"`

	// The isidate compatible natural language description of the schedule.
	Schedule string `json:"schedule,omitempty"`
}