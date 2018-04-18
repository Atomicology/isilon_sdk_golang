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

type JobJobAvscanParams struct {

	// Force files to be scanned, even if excluded by the policy.
	ForceRun bool `json:"force_run,omitempty"`

	// The antivirus scan policy to run.
	Policy string `json:"policy"`

	// An optional report id for the scan.
	ReportId string `json:"report_id,omitempty"`

	// Update the last run time for the policy.
	Update bool `json:"update,omitempty"`
}