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

type CloudJobFiles struct {

	// The file filtering logic to find files for this job
	FileMatchingPattern *Empty `json:"file_matching_pattern,omitempty"`

	// A list of files to be addressed by this job.  (Note* these are only reported when audit_level is 'high'
	Names []CloudJobFilesName `json:"names,omitempty"`

	// The total number of files addressed by this job
	Total int32 `json:"total,omitempty"`

	// The number of canceled files
	TotalCanceled int32 `json:"total_canceled,omitempty"`

	// The number of files which failed
	TotalFailed int32 `json:"total_failed,omitempty"`

	// The number of files pending action
	TotalPending int32 `json:"total_pending,omitempty"`

	// The number of files currently being processed
	TotalProcessing int32 `json:"total_processing,omitempty"`

	// The total number of files successfully completed
	TotalSucceeded int32 `json:"total_succeeded,omitempty"`
}
