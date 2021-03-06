# ProvidersAdsAdsItemExtended

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllocateGids** | **bool** | Allocates an ID for an unmapped Active Directory (ADS) group. ADS groups without GIDs can be proactively assigned a GID by the ID mapper. If the ID mapper option is disabled, GIDs are not proactively assigned, and when a primary group for a user does not include a GID, the system may allocate one.  | [optional] [default to null]
**AllocateUids** | **bool** | Allocates a user ID for an unmapped Active Directory (ADS) user. ADS users without UIDs can be proactively assigned a UID by the ID mapper. IF the ID mapper option is disabled, UIDs are not proactively assigned, and when an identify for a user does not include a UID, the system may allocate one. | [optional] [default to null]
**AssumeDefaultDomain** | **bool** | Enables lookup of unqualified user names in the primary domain. | [optional] [default to null]
**Authentication** | **bool** | Enables authentication and identity management through the authentication provider. | [optional] [default to null]
**CheckOnlineInterval** | **int32** | Specifies the time in seconds between provider online checks. | [optional] [default to null]
**ControllerTime** | **int32** | Specifies the current time for the domain controllers. | [optional] [default to null]
**CreateHomeDirectory** | **bool** | Automatically creates a home directory on the first login. | [optional] [default to null]
**DomainOfflineAlerts** | **bool** | Sends an alert if the domain goes offline. | [optional] [default to null]
**FindableGroups** | **[]string** | Sets list of groups that can be resolved. | [optional] [default to null]
**FindableUsers** | **[]string** | Sets list of users that can be resolved. | [optional] [default to null]
**Forest** | **string** | Specifies the Active Directory forest. | [optional] [default to null]
**Groupnet** | **string** | Groupnet identifier. | [optional] [default to null]
**HomeDirectoryTemplate** | **string** | Specifies the path to the home directory template. | [optional] [default to null]
**Hostname** | **string** | Specifies the fully qualified hostname stored in the machine account. | [optional] [default to null]
**Id** | **string** | Specifies the ID of the Active Directory provider instance. | [optional] [default to null]
**IgnoreAllTrusts** | **bool** | If set to true, ignores all trusted domains. | [optional] [default to null]
**IgnoredTrustedDomains** | **[]string** | Includes trusted domains when &#39;ignore_all_trusts&#39; is set to false. | [optional] [default to null]
**IncludeTrustedDomains** | **[]string** | Includes trusted domains when &#39;ignore_all_trusts&#39; is set to true. | [optional] [default to null]
**Instance** | **string** | Specifies Active Directory provider instance. | [optional] [default to null]
**LdapSignAndSeal** | **bool** | Enables encryption and signing on LDAP requests. | [optional] [default to null]
**LoginShell** | **string** | Specifies the login shell path. | [optional] [default to null]
**LookupDomains** | **[]string** | Limits user and group lookups to the specified domains. | [optional] [default to null]
**LookupGroups** | **bool** | Looks up AD groups in other providers before allocating a group ID. | [optional] [default to null]
**LookupNormalizeGroups** | **bool** | Normalizes AD group names to lowercase before look up. | [optional] [default to null]
**LookupNormalizeUsers** | **bool** | Normalize AD user names to lowercase before look up. | [optional] [default to null]
**LookupUsers** | **bool** | Looks up AD users in other providers before allocating a user ID. | [optional] [default to null]
**MachineAccount** | **string** | Specifies the SAM account name of the machine account. | [optional] [default to null]
**MachineName** | **string** | Specifies name to join AD as. | [optional] [default to null]
**MachinePasswordChanges** | **bool** | Enables periodic changes of the machine password for security. | [optional] [default to null]
**MachinePasswordLifespan** | **int32** | Sets maximum age of a password in seconds. | [optional] [default to null]
**Name** | **string** | Specifies the Active Directory provider name. | [optional] [default to null]
**NetbiosDomain** | **string** | Specifies the NetBIOS domain name associated with the machine account. | [optional] [default to null]
**NodeDcAffinity** | **string** | Specifies the domain controller for which the node has affinity. | [optional] [default to null]
**NodeDcAffinityTimeout** | **int32** | Specifies the timeout for the domain controller for which the local node has affinity. | [optional] [default to null]
**NssEnumeration** | **bool** | Enables the Active Directory provider to respond to &#39;getpwent&#39; and &#39;getgrent&#39; requests. | [optional] [default to null]
**PrimaryDomain** | **string** | Specifies the AD domain to which the provider is joined. | [optional] [default to null]
**RestrictFindable** | **bool** | Check the provider for filtered lists of findable and unfindable users and groups. | [optional] [default to null]
**SfuSupport** | **string** | Specifies whether to support RFC 2307 attributes on ADS domain controllers. | [optional] [default to null]
**Site** | **string** | Specifies the site for the Active Directory. | [optional] [default to null]
**Status** | **string** | Specifies the status of the provider. | [optional] [default to null]
**StoreSfuMappings** | **bool** | Stores SFU mappings permanently in the ID mapper. | [optional] [default to null]
**System** | **bool** | If set to true, indicates that this provider instance was created by OneFS and cannot be removed. | [optional] [default to null]
**UnfindableGroups** | **[]string** | Specifies groups that cannot be resolved by the provider. | [optional] [default to null]
**UnfindableUsers** | **[]string** | Specifies users that cannot be resolved by the provider. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


