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

type DedupeReportExtended struct {

	// The amount of space the directory trees under this job's paths now take up, compared to what they would take up if not deduplicated (0 ~ 100).
	DedupePercent string `json:"dedupe_percent,omitempty"`

	// The amount of time in seconds it took to run this job.
	ElapsedTime int32 `json:"elapsed_time,omitempty"`

	// An unique identifier for this report.
	Id int32 `json:"id,omitempty"`

	// The job id this report refers to.
	JobId int32 `json:"job_id,omitempty"`

	// The type of dedupe job this report refers to.
	JobType string `json:"job_type,omitempty"`

	// A list of report entries for this dedupe job.
	Reports []DedupeReport `json:"reports,omitempty"`
}
