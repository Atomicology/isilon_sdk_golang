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

type SettingsReportsSettings struct {

	// The directory on /ifs where manual or live reports will be placed.
	LiveDir string `json:"live_dir"`

	// The number of manual reports to keep.
	LiveRetain int32 `json:"live_retain"`

	// The isidate schedule used to generate reports.
	Schedule string `json:"schedule"`

	// The directory on /ifs where schedule reports will be placed.
	ScheduledDir string `json:"scheduled_dir"`

	// The number of scheduled reports to keep.
	ScheduledRetain int32 `json:"scheduled_retain"`
}
