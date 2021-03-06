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

type SmbSettingsGlobalExtended struct {

	// Only enumerate files and folders the requesting user has access to.
	AccessBasedShareEnum bool `json:"access_based_share_enum,omitempty"`

	// Specify level of file share audit events to log.
	AuditFileshare string `json:"audit_fileshare,omitempty"`

	// Specifies a list of permissions to audit.
	AuditGlobalSacl []SmbSettingsGlobalSettingsAuditGlobalSaclItem `json:"audit_global_sacl,omitempty"`

	// Specify the level of logon audit events to log.
	AuditLogon string `json:"audit_logon,omitempty"`

	// Allow access to .snapshot directories in share subdirectories.
	DotSnapAccessibleChild bool `json:"dot_snap_accessible_child,omitempty"`

	// Allow access to the .snapshot directory in the root of the share.
	DotSnapAccessibleRoot bool `json:"dot_snap_accessible_root,omitempty"`

	// Show .snapshot directories in share subdirectories.
	DotSnapVisibleChild bool `json:"dot_snap_visible_child,omitempty"`

	// Show the .snapshot directory in the root of a share.
	DotSnapVisibleRoot bool `json:"dot_snap_visible_root,omitempty"`

	// Indicates whether the server supports signed SMB packets.
	EnableSecuritySignatures bool `json:"enable_security_signatures,omitempty"`

	// Specifies the fully-qualified user to use for guest access.
	GuestUser string `json:"guest_user,omitempty"`

	// Specify whether to ignore EAs on files.
	IgnoreEas bool `json:"ignore_eas,omitempty"`

	// Specify the number of OneFS driver worker threads per CPU.
	OnefsCpuMultiplier int32 `json:"onefs_cpu_multiplier,omitempty"`

	// Set the maximum number of OneFS driver worker threads.
	OnefsNumWorkers int32 `json:"onefs_num_workers,omitempty"`

	// Indicates whether the server requires signed SMB packets.
	RequireSecuritySignatures bool `json:"require_security_signatures,omitempty"`

	// Enable Server Side Copy.
	ServerSideCopy bool `json:"server_side_copy,omitempty"`

	// Provides a description of the server.
	ServerString string `json:"server_string,omitempty"`

	// Specify whether service is enabled.
	Service bool `json:"service,omitempty"`

	// Specify the number of SRV service worker threads per CPU.
	SrvCpuMultiplier int32 `json:"srv_cpu_multiplier,omitempty"`

	// Set the maximum number of SRV service worker threads.
	SrvNumWorkers int32 `json:"srv_num_workers,omitempty"`

	// Support multichannel.
	SupportMultichannel bool `json:"support_multichannel,omitempty"`

	// Support NetBIOS.
	SupportNetbios bool `json:"support_netbios,omitempty"`

	// Support the SMB2 protocol on the server.
	SupportSmb2 bool `json:"support_smb2,omitempty"`
}
