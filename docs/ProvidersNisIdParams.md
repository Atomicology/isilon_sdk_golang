# ProvidersNisIdParams

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authentication** | **bool** | If true, enables authentication and identity management through the authentication provider. | [optional] [default to null]
**BalanceServers** | **bool** | If true, connects the provider to a random server. | [optional] [default to null]
**CheckOnlineInterval** | **int32** | Specifies the time in seconds between provider online checks. | [optional] [default to null]
**CreateHomeDirectory** | **bool** | Automatically creates the home directory on the first login. | [optional] [default to null]
**Enabled** | **bool** | If true, enables the NIS provider. | [optional] [default to null]
**EnumerateGroups** | **bool** | If true, allows the provider to enumerate groups. | [optional] [default to null]
**EnumerateUsers** | **bool** | If true, allows the provider to enumerate users. | [optional] [default to null]
**FindableGroups** | **[]string** | Specifies the list of groups that can be resolved. | [optional] [default to null]
**FindableUsers** | **[]string** | Specifies the list of users that can be resolved. | [optional] [default to null]
**GroupDomain** | **string** | Specifies the domain for this provider through which groups are qualified. | [optional] [default to null]
**HomeDirectoryTemplate** | **string** | Specifies the path to the home directory template. | [optional] [default to null]
**HostnameLookup** | **bool** | If true, enables host name look ups. | [optional] [default to null]
**ListableGroups** | **[]string** | Specifies the groups that can be viewed in the provider. | [optional] [default to null]
**ListableUsers** | **[]string** | Specifies the users that can be viewed in the provider. | [optional] [default to null]
**LoginShell** | **string** | Specifies the login shell path. | [optional] [default to null]
**Name** | **string** | Specifies the NIS provider name. | [optional] [default to null]
**NisDomain** | **string** | Specifies the NIS domain name. | [optional] [default to null]
**NormalizeGroups** | **bool** | Normalizes group names to lowercase before look up. | [optional] [default to null]
**NormalizeUsers** | **bool** | Normalizes user names to lowercase before look up. | [optional] [default to null]
**NtlmSupport** | **string** | Specifies which NTLM versions to support for users with NTLM-compatible credentials. | [optional] [default to null]
**ProviderDomain** | **string** | Specifies the domain for the provider. | [optional] [default to null]
**RequestTimeout** | **int32** | Specifies the request timeout interval in seconds. | [optional] [default to null]
**RestrictFindable** | **bool** | If true, checks the provider for filtered lists of findable and unfindable users and groups. | [optional] [default to null]
**RestrictListable** | **bool** | If true, checks the provider for filtered lists of listable and unlistable users and groups. | [optional] [default to null]
**RetryTime** | **int32** | Specifies the timeout period in seconds after which a request will be retried. | [optional] [default to null]
**Servers** | **[]string** | Adds an NIS server for this provider. | [optional] [default to null]
**UnfindableGroups** | **[]string** | Specifies groups that cannot be resolved by the provider. | [optional] [default to null]
**UnfindableUsers** | **[]string** | Specifies users that cannot be resolved by the provider. | [optional] [default to null]
**UnlistableGroups** | **[]string** | Specifies a group that cannot be listed by the provider. | [optional] [default to null]
**UnlistableUsers** | **[]string** | Specifies a user that cannot be listed by the provider. | [optional] [default to null]
**UserDomain** | **string** | Specifies the domain for this provider through which users are qualified. | [optional] [default to null]
**YpmatchUsingTcp** | **bool** | If true, specifies TCP for YP Match operations. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


