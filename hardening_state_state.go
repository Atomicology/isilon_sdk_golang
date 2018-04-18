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

type HardeningStateState struct {

	// Full path name of issues file, or null if no issues file is configured. This file contains all issues found when the cluster configuration is checked against expected configuration.
	IssuesFile string `json:"issues_file,omitempty"`

	// This contains more information and details about the operation state.
	Message string `json:"message,omitempty"`

	// The state of the hardening operation. In case there is no operation currently going on, this will display the last state of operation.
	State string `json:"state,omitempty"`
}