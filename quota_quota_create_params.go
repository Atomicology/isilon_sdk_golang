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

type QuotaQuotaCreateParams struct {

	// If true, SMB shares using the quota directory see the quota thresholds as share size.
	Container bool `json:"container,omitempty"`

	// True if the quota provides enforcement, otherwise a accounting quota.
	Enforced bool `json:"enforced"`

	// Force creation of quotas on the root of /ifs.
	Force bool `json:"force,omitempty"`

	// If true, quota governs snapshot data as well as head data.
	IncludeSnapshots bool `json:"include_snapshots"`

	// The /ifs path governed.
	Path string `json:"path"`

	// Specifies properties for a persona, which consists of either a 'type' and a 'name' or an 'ID'.
	Persona *AuthAccessAccessItemFileGroup `json:"persona,omitempty"`

	// 
	Thresholds *QuotaQuotaThresholds `json:"thresholds,omitempty"`

	// If true, thresholds apply to data plus filesystem overhead required to store the data (i.e. 'physical' usage).
	ThresholdsIncludeOverhead bool `json:"thresholds_include_overhead"`

	// The type of quota.
	Type_ string `json:"type"`
}
