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

type ProvidersFileFileItem struct {

	// Enables authentication and identity mapping through the authentication provider.
	Authentication bool `json:"authentication,omitempty"`

	// Automatically creates a home directory on the first login.
	CreateHomeDirectory bool `json:"create_home_directory,omitempty"`

	// Enables the file provider.
	Enabled bool `json:"enabled,omitempty"`

	// Enables the provider to enumerate groups.
	EnumerateGroups bool `json:"enumerate_groups,omitempty"`

	// Enables the provider to enumerate users.
	EnumerateUsers bool `json:"enumerate_users,omitempty"`

	// Specifies the list of groups that can be resolved.
	FindableGroups []string `json:"findable_groups,omitempty"`

	// Specifies the list of users that can be resolved.
	FindableUsers []string `json:"findable_users,omitempty"`

	// Specifies the domain for this provider through which domains are qualified.
	GroupDomain string `json:"group_domain,omitempty"`

	// Specifies the location of the file that contains information about the group.
	GroupFile string `json:"group_file,omitempty"`

	// Specifies the path to the home directory template.
	HomeDirectoryTemplate string `json:"home_directory_template,omitempty"`

	// Specifies the file provider ID.
	Id string `json:"id,omitempty"`

	// Specifies the groups that can be viewed in the provider.
	ListableGroups []string `json:"listable_groups,omitempty"`

	// Specifies the users that can be viewed in the provider.
	ListableUsers []string `json:"listable_users,omitempty"`

	// Specifies the login shell path.
	LoginShell string `json:"login_shell,omitempty"`

	// Specifies the groups that can be modified in the provider.
	ModifiableGroups []string `json:"modifiable_groups,omitempty"`

	// Specifies the users that can be modified in the provider.
	ModifiableUsers []string `json:"modifiable_users,omitempty"`

	// Specifies the name of the file provider.
	Name string `json:"name,omitempty"`

	// Specifies the path to a netgroups replacement file.
	NetgroupFile string `json:"netgroup_file,omitempty"`

	// Normalizes group names to lowercase before look up.
	NormalizeGroups bool `json:"normalize_groups,omitempty"`

	// Normalizes user names to lowercase before look up.
	NormalizeUsers bool `json:"normalize_users,omitempty"`

	// Specifies which NTLM versions to support for users with NTLM-compatible credentials.
	NtlmSupport string `json:"ntlm_support,omitempty"`

	// Specifies the location of the file containing information about users.
	PasswordFile string `json:"password_file,omitempty"`

	// Specifies the domain for the provider.
	ProviderDomain string `json:"provider_domain,omitempty"`

	// If true, checks the provider for filtered lists of findable and unfindable users and groups.
	RestrictFindable bool `json:"restrict_findable,omitempty"`

	// If true, checks the provider for filtered lists of listable and unlistable users and groups.
	RestrictListable bool `json:"restrict_listable,omitempty"`

	// If true, checks the provider for filtered lists of modifiable and unmodifiable users and groups.
	RestrictModifiable bool `json:"restrict_modifiable,omitempty"`

	// Specifies the status of the provider.
	Status string `json:"status,omitempty"`

	// If true, indicates that this provider instance was created by OneFS and cannot be removed.
	System bool `json:"system,omitempty"`

	// Specifies groups that cannot be resolved by the provider.
	UnfindableGroups []string `json:"unfindable_groups,omitempty"`

	// Specifies users that cannot be resolved by the provider.
	UnfindableUsers []string `json:"unfindable_users,omitempty"`

	// Specifies a group that cannot be listed by the provider.
	UnlistableGroups []string `json:"unlistable_groups,omitempty"`

	// Specifies a user that cannot be listed by the provider.
	UnlistableUsers []string `json:"unlistable_users,omitempty"`

	// Specifies a group that cannot be modified by the provider.
	UnmodifiableGroups []string `json:"unmodifiable_groups,omitempty"`

	// Specifies a user that cannot be modified by the provider.
	UnmodifiableUsers []string `json:"unmodifiable_users,omitempty"`

	// Specifies the domain for this provider through which users are qualified.
	UserDomain string `json:"user_domain,omitempty"`
}
