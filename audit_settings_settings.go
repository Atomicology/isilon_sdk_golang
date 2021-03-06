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

type AuditSettingsSettings struct {

	// Filter of protocol operations to Audit when they fail.
	AuditFailure []string `json:"audit_failure,omitempty"`

	// Filter of protocol operations to Audit when they succeed.
	AuditSuccess []string `json:"audit_success,omitempty"`

	// Filter of Audit operations to forward to syslog.
	SyslogAuditEvents []string `json:"syslog_audit_events,omitempty"`

	// Enables forwarding of events to syslog.
	SyslogForwardingEnabled bool `json:"syslog_forwarding_enabled,omitempty"`
}
