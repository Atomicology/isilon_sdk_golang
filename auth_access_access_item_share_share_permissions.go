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

type AuthAccessAccessItemShareSharePermissions struct {

	// Returns Share level permissions for the user.{ 'read' , 'write' , 'full' or 'none' will be the values}
	ExpectedPermissions string `json:"expected_permissions,omitempty"`

	// Returns whether impersonate guest setting is enabled for the user on the share.
	ImpersonateGuest bool `json:"impersonate_guest,omitempty"`

	// Returns whether impersonate user setting is enabled on the share
	ImpersonateUser bool `json:"impersonate_user,omitempty"`

	// Returns whether run as root is enabled for the user on the share
	RunAsRoot bool `json:"run_as_root,omitempty"`

	// Specifies a list of the relevant Access Control Entries withrespect to the user in the share.
	ShareRelevantAces []AuthAccessAccessItemShareSharePermissionsShareRelevantAce `json:"share_relevant_aces,omitempty"`
}
