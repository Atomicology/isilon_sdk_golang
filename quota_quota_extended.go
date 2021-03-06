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

type QuotaQuotaExtended struct {

	// If true, SMB shares using the quota directory see the quota thresholds as share size.
	Container bool `json:"container"`

	// True if the quota provides enforcement, otherwise a accounting quota.
	Enforced bool `json:"enforced"`

	// The system ID given to the quota.
	Id string `json:"id"`

	// If true, quota governs snapshot data as well as head data.
	IncludeSnapshots bool `json:"include_snapshots"`

	// For user and group quotas, true if the quota is linked and controlled by a parent default-* quota. Linked quotas cannot be modified until they are unlinked.
	Linked bool `json:"linked,omitempty"`

	// Summary of notifications: 'custom' indicates one or more notification rules available from the notifications sub-resource; 'default' indicates system default rules are used; 'disabled' indicates that no notifications will be used for this quota.
	Notifications string `json:"notifications"`

	// The /ifs path governed.
	Path string `json:"path"`

	// Specifies properties for a persona, which consists of either a 'type' and a 'name' or an 'ID'.
	Persona *AuthAccessAccessItemFileGroup `json:"persona,omitempty"`

	// True if the accounting is accurate on the quota.  If false, this quota is waiting on completion of a QuotaScan job.
	Ready bool `json:"ready"`

	// 
	Thresholds *QuotaQuotaThresholdsExtended `json:"thresholds"`

	// If true, thresholds apply to data plus filesystem overhead required to store the data (i.e. 'physical' usage).
	ThresholdsIncludeOverhead bool `json:"thresholds_include_overhead"`

	// The type of quota.
	Type_ string `json:"type"`

	// 
	Usage *QuotaQuotaUsage `json:"usage"`
}
