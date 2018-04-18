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

// Specifies under what conditions and over which channel an alert should be sent
type EventAlertCondition struct {

	// Event Group categories to be alerted: all, 100000000 (SYS_DISK_EVENTS), 200000000 (NODE_STATUS_EVENTS), 300000000 (REBOOT_EVENTS), 400000000 (SW_EVENTS), 500000000 (QUOTA_EVENTS), 600000000 (SNAP_EVENTS), 700000000 (WINNET_EVENTS), 800000000 (FILESYS_EVENTS), 900000000 (HW_EVENTS), 1100000000 (CPOOL_EVENTS)
	Categories []string `json:"categories,omitempty"`

	// Channels for alert
	Channels []string `json:"channels,omitempty"`

	// Trigger condition for alert
	Condition string `json:"condition,omitempty"`

	// Event Group IDs to be alerted
	EventgroupIds []string `json:"eventgroup_ids,omitempty"`

	// Required with ONGOING condition only, period in seconds between alerts of ongoing conditions
	Interval int32 `json:"interval,omitempty"`

	// Required with NEW EVENTS condition only, limits the number of alerts sent as events are added
	Limit int32 `json:"limit,omitempty"`

	// Severities to be alerted
	Severities []string `json:"severities,omitempty"`

	// Any eventgroup lasting less than this many seconds is deemed transient and will not generate alerts under this condition.
	Transient int32 `json:"transient,omitempty"`
}
