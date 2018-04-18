# ProvidersFileIdParams

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authentication** | **bool** | Enables authentication and identity mapping through the authentication provider. | [optional] [default to null]
**CreateHomeDirectory** | **bool** | Automatically creates a home directory on the first login. | [optional] [default to null]
**Enabled** | **bool** | Enables the file provider. | [optional] [default to null]
**EnumerateGroups** | **bool** | Enables the provider to enumerate groups. | [optional] [default to null]
**EnumerateUsers** | **bool** | Enables the provider to enumerate users. | [optional] [default to null]
**FindableGroups** | **[]string** | Specifies the list of groups that can be resolved. | [optional] [default to null]
**FindableUsers** | **[]string** | Specifies the list of users that can be resolved. | [optional] [default to null]
**GroupDomain** | **string** | Specifies the domain for this provider through which domains are qualified. | [optional] [default to null]
**GroupFile** | **string** | Specifies the location of the file that contains information about the group. | [optional] [default to null]
**HomeDirectoryTemplate** | **string** | Specifies the path to the home directory template. | [optional] [default to null]
**ListableGroups** | **[]string** | Specifies the groups that can be viewed in the provider. | [optional] [default to null]
**ListableUsers** | **[]string** | Specifies the users that can be viewed in the provider. | [optional] [default to null]
**LoginShell** | **string** | Specifies the login shell path. | [optional] [default to null]
**ModifiableGroups** | **[]string** | Specifies the groups that can be modified in the provider. | [optional] [default to null]
**ModifiableUsers** | **[]string** | Specifies the users that can be modified in the provider. | [optional] [default to null]
**Name** | **string** | Specifies the name of the file provider. | [optional] [default to null]
**NetgroupFile** | **string** | Specifies the path to a netgroups replacement file. | [optional] [default to null]
**NormalizeGroups** | **bool** | Normalizes group names to lowercase before look up. | [optional] [default to null]
**NormalizeUsers** | **bool** | Normalizes user names to lowercase before look up. | [optional] [default to null]
**NtlmSupport** | **string** | Specifies which NTLM versions to support for users with NTLM-compatible credentials. | [optional] [default to null]
**PasswordFile** | **string** | Specifies the location of the file containing information about users. | [optional] [default to null]
**ProviderDomain** | **string** | Specifies the domain for the provider. | [optional] [default to null]
**RestrictFindable** | **bool** | If true, checks the provider for filtered lists of findable and unfindable users and groups. | [optional] [default to null]
**RestrictListable** | **bool** | If true, checks the provider for filtered lists of listable and unlistable users and groups. | [optional] [default to null]
**RestrictModifiable** | **bool** | If true, checks the provider for filtered lists of modifiable and unmodifiable users and groups. | [optional] [default to null]
**UnfindableGroups** | **[]string** | Specifies groups that cannot be resolved by the provider. | [optional] [default to null]
**UnfindableUsers** | **[]string** | Specifies users that cannot be resolved by the provider. | [optional] [default to null]
**UnlistableGroups** | **[]string** | Specifies a group that cannot be listed by the provider. | [optional] [default to null]
**UnlistableUsers** | **[]string** | Specifies a user that cannot be listed by the provider. | [optional] [default to null]
**UnmodifiableGroups** | **[]string** | Specifies a group that cannot be modified by the provider. | [optional] [default to null]
**UnmodifiableUsers** | **[]string** | Specifies a user that cannot be modified by the provider. | [optional] [default to null]
**UserDomain** | **string** | Specifies the domain for this provider through which users are qualified. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


