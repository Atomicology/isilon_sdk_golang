# \AuthApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAuthCacheItem**](AuthApi.md#CreateAuthCacheItem) | **Post** /platform/4/auth/cache | 
[**CreateAuthGroup**](AuthApi.md#CreateAuthGroup) | **Post** /platform/1/auth/groups | 
[**CreateAuthRefreshItem**](AuthApi.md#CreateAuthRefreshItem) | **Post** /platform/3/auth/refresh | 
[**CreateAuthRole**](AuthApi.md#CreateAuthRole) | **Post** /platform/1/auth/roles | 
[**CreateAuthUser**](AuthApi.md#CreateAuthUser) | **Post** /platform/1/auth/users | 
[**CreateMappingIdentity**](AuthApi.md#CreateMappingIdentity) | **Post** /platform/1/auth/mapping/identities | 
[**CreateMappingIdentity_0**](AuthApi.md#CreateMappingIdentity_0) | **Post** /platform/1/auth/mapping/identities/{MappingIdentityId} | 
[**CreateProvidersAdsItem**](AuthApi.md#CreateProvidersAdsItem) | **Post** /platform/3/auth/providers/ads | 
[**CreateProvidersFileItem**](AuthApi.md#CreateProvidersFileItem) | **Post** /platform/1/auth/providers/file | 
[**CreateProvidersKrb5Item**](AuthApi.md#CreateProvidersKrb5Item) | **Post** /platform/3/auth/providers/krb5 | 
[**CreateProvidersLdapItem**](AuthApi.md#CreateProvidersLdapItem) | **Post** /platform/4/auth/providers/ldap | 
[**CreateProvidersNisItem**](AuthApi.md#CreateProvidersNisItem) | **Post** /platform/3/auth/providers/nis | 
[**CreateSettingsKrb5Domain**](AuthApi.md#CreateSettingsKrb5Domain) | **Post** /platform/1/auth/settings/krb5/domains | 
[**CreateSettingsKrb5Realm**](AuthApi.md#CreateSettingsKrb5Realm) | **Post** /platform/1/auth/settings/krb5/realms | 
[**DeleteAuthGroup**](AuthApi.md#DeleteAuthGroup) | **Delete** /platform/1/auth/groups/{AuthGroupId} | 
[**DeleteAuthGroups**](AuthApi.md#DeleteAuthGroups) | **Delete** /platform/1/auth/groups | 
[**DeleteAuthRole**](AuthApi.md#DeleteAuthRole) | **Delete** /platform/1/auth/roles/{AuthRoleId} | 
[**DeleteAuthUser**](AuthApi.md#DeleteAuthUser) | **Delete** /platform/1/auth/users/{AuthUserId} | 
[**DeleteAuthUsers**](AuthApi.md#DeleteAuthUsers) | **Delete** /platform/1/auth/users | 
[**DeleteMappingIdentities**](AuthApi.md#DeleteMappingIdentities) | **Delete** /platform/1/auth/mapping/identities | 
[**DeleteMappingIdentity**](AuthApi.md#DeleteMappingIdentity) | **Delete** /platform/1/auth/mapping/identities/{MappingIdentityId} | 
[**DeleteProvidersAdsById**](AuthApi.md#DeleteProvidersAdsById) | **Delete** /platform/3/auth/providers/ads/{ProvidersAdsId} | 
[**DeleteProvidersFileById**](AuthApi.md#DeleteProvidersFileById) | **Delete** /platform/1/auth/providers/file/{ProvidersFileId} | 
[**DeleteProvidersKrb5ById**](AuthApi.md#DeleteProvidersKrb5ById) | **Delete** /platform/3/auth/providers/krb5/{ProvidersKrb5Id} | 
[**DeleteProvidersLdapById**](AuthApi.md#DeleteProvidersLdapById) | **Delete** /platform/4/auth/providers/ldap/{ProvidersLdapId} | 
[**DeleteProvidersLocalById**](AuthApi.md#DeleteProvidersLocalById) | **Delete** /platform/1/auth/providers/local/{ProvidersLocalId} | 
[**DeleteProvidersNisById**](AuthApi.md#DeleteProvidersNisById) | **Delete** /platform/3/auth/providers/nis/{ProvidersNisId} | 
[**DeleteSettingsKrb5Domain**](AuthApi.md#DeleteSettingsKrb5Domain) | **Delete** /platform/1/auth/settings/krb5/domains/{SettingsKrb5DomainId} | 
[**DeleteSettingsKrb5Realm**](AuthApi.md#DeleteSettingsKrb5Realm) | **Delete** /platform/1/auth/settings/krb5/realms/{SettingsKrb5RealmId} | 
[**GetAuthAccessUser**](AuthApi.md#GetAuthAccessUser) | **Get** /platform/1/auth/access/{AuthAccessUser} | 
[**GetAuthGroup**](AuthApi.md#GetAuthGroup) | **Get** /platform/1/auth/groups/{AuthGroupId} | 
[**GetAuthId**](AuthApi.md#GetAuthId) | **Get** /platform/1/auth/id | 
[**GetAuthLdapTemplate**](AuthApi.md#GetAuthLdapTemplate) | **Get** /platform/4/auth/ldap-templates/{AuthLdapTemplateId} | 
[**GetAuthLdapTemplates**](AuthApi.md#GetAuthLdapTemplates) | **Get** /platform/4/auth/ldap-templates | 
[**GetAuthLogLevel**](AuthApi.md#GetAuthLogLevel) | **Get** /platform/3/auth/log-level | 
[**GetAuthNetgroup**](AuthApi.md#GetAuthNetgroup) | **Get** /platform/1/auth/netgroups/{AuthNetgroupId} | 
[**GetAuthPrivileges**](AuthApi.md#GetAuthPrivileges) | **Get** /platform/1/auth/privileges | 
[**GetAuthRole**](AuthApi.md#GetAuthRole) | **Get** /platform/1/auth/roles/{AuthRoleId} | 
[**GetAuthShells**](AuthApi.md#GetAuthShells) | **Get** /platform/1/auth/shells | 
[**GetAuthUser**](AuthApi.md#GetAuthUser) | **Get** /platform/1/auth/users/{AuthUserId} | 
[**GetAuthWellknown**](AuthApi.md#GetAuthWellknown) | **Get** /platform/1/auth/wellknowns/{AuthWellknownId} | 
[**GetAuthWellknowns**](AuthApi.md#GetAuthWellknowns) | **Get** /platform/1/auth/wellknowns | 
[**GetMappingDump**](AuthApi.md#GetMappingDump) | **Get** /platform/3/auth/mapping/dump | 
[**GetMappingIdentity**](AuthApi.md#GetMappingIdentity) | **Get** /platform/1/auth/mapping/identities/{MappingIdentityId} | 
[**GetMappingUsersLookup**](AuthApi.md#GetMappingUsersLookup) | **Get** /platform/1/auth/mapping/users/lookup | 
[**GetMappingUsersRules**](AuthApi.md#GetMappingUsersRules) | **Get** /platform/1/auth/mapping/users/rules | 
[**GetProvidersAdsById**](AuthApi.md#GetProvidersAdsById) | **Get** /platform/3/auth/providers/ads/{ProvidersAdsId} | 
[**GetProvidersFileById**](AuthApi.md#GetProvidersFileById) | **Get** /platform/1/auth/providers/file/{ProvidersFileId} | 
[**GetProvidersKrb5ById**](AuthApi.md#GetProvidersKrb5ById) | **Get** /platform/3/auth/providers/krb5/{ProvidersKrb5Id} | 
[**GetProvidersLdapById**](AuthApi.md#GetProvidersLdapById) | **Get** /platform/4/auth/providers/ldap/{ProvidersLdapId} | 
[**GetProvidersLocal**](AuthApi.md#GetProvidersLocal) | **Get** /platform/1/auth/providers/local | 
[**GetProvidersLocalById**](AuthApi.md#GetProvidersLocalById) | **Get** /platform/1/auth/providers/local/{ProvidersLocalId} | 
[**GetProvidersNisById**](AuthApi.md#GetProvidersNisById) | **Get** /platform/3/auth/providers/nis/{ProvidersNisId} | 
[**GetProvidersSummary**](AuthApi.md#GetProvidersSummary) | **Get** /platform/3/auth/providers/summary | 
[**GetSettingsAcls**](AuthApi.md#GetSettingsAcls) | **Get** /platform/3/auth/settings/acls | 
[**GetSettingsGlobal**](AuthApi.md#GetSettingsGlobal) | **Get** /platform/1/auth/settings/global | 
[**GetSettingsKrb5Defaults**](AuthApi.md#GetSettingsKrb5Defaults) | **Get** /platform/1/auth/settings/krb5/defaults | 
[**GetSettingsKrb5Domain**](AuthApi.md#GetSettingsKrb5Domain) | **Get** /platform/1/auth/settings/krb5/domains/{SettingsKrb5DomainId} | 
[**GetSettingsKrb5Realm**](AuthApi.md#GetSettingsKrb5Realm) | **Get** /platform/1/auth/settings/krb5/realms/{SettingsKrb5RealmId} | 
[**GetSettingsMapping**](AuthApi.md#GetSettingsMapping) | **Get** /platform/1/auth/settings/mapping | 
[**ListAuthGroups**](AuthApi.md#ListAuthGroups) | **Get** /platform/1/auth/groups | 
[**ListAuthRoles**](AuthApi.md#ListAuthRoles) | **Get** /platform/1/auth/roles | 
[**ListAuthUsers**](AuthApi.md#ListAuthUsers) | **Get** /platform/1/auth/users | 
[**ListProvidersAds**](AuthApi.md#ListProvidersAds) | **Get** /platform/3/auth/providers/ads | 
[**ListProvidersFile**](AuthApi.md#ListProvidersFile) | **Get** /platform/1/auth/providers/file | 
[**ListProvidersKrb5**](AuthApi.md#ListProvidersKrb5) | **Get** /platform/3/auth/providers/krb5 | 
[**ListProvidersLdap**](AuthApi.md#ListProvidersLdap) | **Get** /platform/4/auth/providers/ldap | 
[**ListProvidersNis**](AuthApi.md#ListProvidersNis) | **Get** /platform/3/auth/providers/nis | 
[**ListSettingsKrb5Domains**](AuthApi.md#ListSettingsKrb5Domains) | **Get** /platform/1/auth/settings/krb5/domains | 
[**ListSettingsKrb5Realms**](AuthApi.md#ListSettingsKrb5Realms) | **Get** /platform/1/auth/settings/krb5/realms | 
[**UpdateAuthGroup**](AuthApi.md#UpdateAuthGroup) | **Put** /platform/1/auth/groups/{AuthGroupId} | 
[**UpdateAuthLogLevel**](AuthApi.md#UpdateAuthLogLevel) | **Put** /platform/3/auth/log-level | 
[**UpdateAuthRole**](AuthApi.md#UpdateAuthRole) | **Put** /platform/1/auth/roles/{AuthRoleId} | 
[**UpdateAuthUser**](AuthApi.md#UpdateAuthUser) | **Put** /platform/1/auth/users/{AuthUserId} | 
[**UpdateMappingImport**](AuthApi.md#UpdateMappingImport) | **Put** /platform/3/auth/mapping/import | 
[**UpdateMappingUsersRules**](AuthApi.md#UpdateMappingUsersRules) | **Put** /platform/1/auth/mapping/users/rules | 
[**UpdateProvidersAdsById**](AuthApi.md#UpdateProvidersAdsById) | **Put** /platform/3/auth/providers/ads/{ProvidersAdsId} | 
[**UpdateProvidersFileById**](AuthApi.md#UpdateProvidersFileById) | **Put** /platform/1/auth/providers/file/{ProvidersFileId} | 
[**UpdateProvidersKrb5ById**](AuthApi.md#UpdateProvidersKrb5ById) | **Put** /platform/3/auth/providers/krb5/{ProvidersKrb5Id} | 
[**UpdateProvidersLdapById**](AuthApi.md#UpdateProvidersLdapById) | **Put** /platform/4/auth/providers/ldap/{ProvidersLdapId} | 
[**UpdateProvidersLocalById**](AuthApi.md#UpdateProvidersLocalById) | **Put** /platform/1/auth/providers/local/{ProvidersLocalId} | 
[**UpdateProvidersNisById**](AuthApi.md#UpdateProvidersNisById) | **Put** /platform/3/auth/providers/nis/{ProvidersNisId} | 
[**UpdateSettingsAcls**](AuthApi.md#UpdateSettingsAcls) | **Put** /platform/3/auth/settings/acls | 
[**UpdateSettingsGlobal**](AuthApi.md#UpdateSettingsGlobal) | **Put** /platform/1/auth/settings/global | 
[**UpdateSettingsKrb5Defaults**](AuthApi.md#UpdateSettingsKrb5Defaults) | **Put** /platform/1/auth/settings/krb5/defaults | 
[**UpdateSettingsKrb5Domain**](AuthApi.md#UpdateSettingsKrb5Domain) | **Put** /platform/1/auth/settings/krb5/domains/{SettingsKrb5DomainId} | 
[**UpdateSettingsKrb5Realm**](AuthApi.md#UpdateSettingsKrb5Realm) | **Put** /platform/1/auth/settings/krb5/realms/{SettingsKrb5RealmId} | 
[**UpdateSettingsMapping**](AuthApi.md#UpdateSettingsMapping) | **Put** /platform/1/auth/settings/mapping | 


# **CreateAuthCacheItem**
> CreateResponse CreateAuthCacheItem(ctx, authCacheItem, optional)


Flush the Security Objects Cache.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **authCacheItem** | [**AuthCacheItem**](AuthCacheItem.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authCacheItem** | [**AuthCacheItem**](AuthCacheItem.md)|  | 
 **zone** | **string**| Specifies access zone from which to flush objects. | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAuthGroup**
> CreateResponse CreateAuthGroup(ctx, authGroup, optional)


Create a new group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **authGroup** | [**AuthGroupCreateParams**](AuthGroupCreateParams.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authGroup** | [**AuthGroupCreateParams**](AuthGroupCreateParams.md)|  | 
 **force** | **bool**| Skip validation checks when creating a group. | 
 **zone** | **string**| Optional zone. | 
 **provider** | **string**| Optional provider type. | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAuthRefreshItem**
> CreateAuthRefreshItemResponse CreateAuthRefreshItem(ctx, authRefreshItem)


Refresh the authentication service configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **authRefreshItem** | [**Empty**](Empty.md)|  | 

### Return type

[**CreateAuthRefreshItemResponse**](CreateAuthRefreshItemResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAuthRole**
> CreateResponse CreateAuthRole(ctx, authRole)


Create a new role.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **authRole** | [**AuthRoleCreateParams**](AuthRoleCreateParams.md)|  | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAuthUser**
> CreateResponse CreateAuthUser(ctx, authUser, optional)


Create a new user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **authUser** | [**AuthUserCreateParams**](AuthUserCreateParams.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authUser** | [**AuthUserCreateParams**](AuthUserCreateParams.md)|  | 
 **force** | **bool**| Skip validation checks when creating user. | 
 **zone** | **string**| Optional zone. | 
 **provider** | **string**| Optional provider type. | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMappingIdentity**
> Empty CreateMappingIdentity(ctx, mappingIdentity, optional)


Manually set or modify a mapping between two personae.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **mappingIdentity** | [**MappingIdentityCreateParams**](MappingIdentityCreateParams.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mappingIdentity** | [**MappingIdentityCreateParams**](MappingIdentityCreateParams.md)|  | 
 **var2way** | **bool**| Create a bi-directional mapping from source to target and target to source. | 
 **zone** | **string**| Optional zone. | 
 **replace** | **bool**| Replace existing mappings. | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMappingIdentity_0**
> MappingIdentities CreateMappingIdentity_0(ctx, mappingIdentity, mappingIdentityId, optional)


Manually set or modify a mapping between two personae.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **mappingIdentity** | [**Empty**](Empty.md)|  | 
  **mappingIdentityId** | **string**| Manually set or modify a mapping between two personae. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mappingIdentity** | [**Empty**](Empty.md)|  | 
 **mappingIdentityId** | **string**| Manually set or modify a mapping between two personae. | 
 **type_** | **string**| Desired mapping target to fetch/generate. | 
 **zone** | **string**| Optional zone. | 

### Return type

[**MappingIdentities**](MappingIdentities.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateProvidersAdsItem**
> CreateResponse CreateProvidersAdsItem(ctx, providersAdsItem)


Create a new ADS provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providersAdsItem** | [**ProvidersAdsItem**](ProvidersAdsItem.md)|  | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateProvidersFileItem**
> CreateResponse CreateProvidersFileItem(ctx, providersFileItem)


Create a new file provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providersFileItem** | [**ProvidersFileItem**](ProvidersFileItem.md)|  | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateProvidersKrb5Item**
> CreateResponse CreateProvidersKrb5Item(ctx, providersKrb5Item)


Create a new KRB5 provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providersKrb5Item** | [**ProvidersKrb5Item**](ProvidersKrb5Item.md)|  | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateProvidersLdapItem**
> CreateResponse CreateProvidersLdapItem(ctx, providersLdapItem, optional)


Create a new LDAP provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providersLdapItem** | [**ProvidersLdapItem**](ProvidersLdapItem.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providersLdapItem** | [**ProvidersLdapItem**](ProvidersLdapItem.md)|  | 
 **force** | **bool**| Ignore unresolvable server URIs. | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateProvidersNisItem**
> CreateResponse CreateProvidersNisItem(ctx, providersNisItem)


Create a new NIS provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providersNisItem** | [**ProvidersNisItem**](ProvidersNisItem.md)|  | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSettingsKrb5Domain**
> CreateResponse CreateSettingsKrb5Domain(ctx, settingsKrb5Domain)


Create a new krb5 domain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **settingsKrb5Domain** | [**SettingsKrb5DomainCreateParams**](SettingsKrb5DomainCreateParams.md)|  | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSettingsKrb5Realm**
> CreateResponse CreateSettingsKrb5Realm(ctx, settingsKrb5Realm)


Create a new krb5 realm.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **settingsKrb5Realm** | [**SettingsKrb5RealmCreateParams**](SettingsKrb5RealmCreateParams.md)|  | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAuthGroup**
> DeleteAuthGroup(ctx, authGroupId, optional)


Delete the group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **authGroupId** | **string**| Delete the group. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authGroupId** | **string**| Delete the group. | 
 **cached** | **bool**| If true, flush the group from the cache. | 
 **zone** | **string**| Filter groups by zone. | 
 **provider** | **string**| Filter groups by provider. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAuthGroups**
> DeleteAuthGroups(ctx, optional)


Flush the groups cache.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cached** | **bool**| If true, only flush cached objects. | 
 **zone** | **string**| Filter groups by zone. | 
 **provider** | **string**| Filter groups by provider. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAuthRole**
> DeleteAuthRole(ctx, authRoleId)


Delete the role.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **authRoleId** | **string**| Delete the role. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAuthUser**
> DeleteAuthUser(ctx, authUserId, optional)


Delete the user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **authUserId** | **string**| Delete the user. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authUserId** | **string**| Delete the user. | 
 **cached** | **bool**| If true, flush the user from the cache. | 
 **zone** | **string**| Filter users by zone. | 
 **provider** | **string**| Filter users by provider. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAuthUsers**
> DeleteAuthUsers(ctx, optional)


Flush the users cache.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cached** | **bool**| If true, only flush cached objects. | 
 **zone** | **string**| Filter users by zone. | 
 **provider** | **string**| Filter users by provider. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMappingIdentities**
> DeleteMappingIdentities(ctx, optional)


Flush the entire idmap cache.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string**| Filter to apply when deleting identity mappings. | 
 **zone** | **string**| Optional zone. | 
 **remove** | **bool**| Delete mapping instead of flush mapping cache. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMappingIdentity**
> DeleteMappingIdentity(ctx, mappingIdentityId, optional)


Flush the entire idmap cache.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **mappingIdentityId** | **string**| Flush the entire idmap cache. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mappingIdentityId** | **string**| Flush the entire idmap cache. | 
 **zone** | **string**| Optional zone. | 
 **var2way** | **bool**| Delete the bi-directional mapping from source to target and target to source. | 
 **target** | **string**| Target identity persona. | 
 **remove** | **bool**| Delete mapping instead of flush mapping from cache. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProvidersAdsById**
> DeleteProvidersAdsById(ctx, providersAdsId)


Delete the ADS provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providersAdsId** | **string**| Delete the ADS provider. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProvidersFileById**
> DeleteProvidersFileById(ctx, providersFileId)


Delete the file provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providersFileId** | **string**| Delete the file provider. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProvidersKrb5ById**
> DeleteProvidersKrb5ById(ctx, providersKrb5Id)


Delete the KRB5 provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providersKrb5Id** | **string**| Delete the KRB5 provider. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProvidersLdapById**
> DeleteProvidersLdapById(ctx, providersLdapId)


Delete the LDAP provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providersLdapId** | **string**| Delete the LDAP provider. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProvidersLocalById**
> DeleteProvidersLocalById(ctx, providersLocalId)


Delete the local provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providersLocalId** | **string**| Delete the local provider. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProvidersNisById**
> DeleteProvidersNisById(ctx, providersNisId)


Delete the NIS provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providersNisId** | **string**| Delete the NIS provider. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSettingsKrb5Domain**
> DeleteSettingsKrb5Domain(ctx, settingsKrb5DomainId)


Remove a krb5 domain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **settingsKrb5DomainId** | **string**| Remove a krb5 domain. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSettingsKrb5Realm**
> DeleteSettingsKrb5Realm(ctx, settingsKrb5RealmId)


Remove a realm.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **settingsKrb5RealmId** | **string**| Remove a realm. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuthAccessUser**
> AuthAccess GetAuthAccessUser(ctx, authAccessUser, optional)


Determine user's access rights to a file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **authAccessUser** | **string**| Determine user&#39;s access rights to a file | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authAccessUser** | **string**| Determine user&#39;s access rights to a file | 
 **path** | **string**| Path to the file. Must be within /ifs. | 
 **share** | **string**| SMB share name | 
 **zone** | **string**| Access zone the user is in. | 
 **numeric** | **bool**| Show the user&#39;s numeric identifier. | 

### Return type

[**AuthAccess**](AuthAccess.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuthGroup**
> AuthGroups GetAuthGroup(ctx, authGroupId, optional)


Retrieve the group information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **authGroupId** | **string**| Retrieve the group information. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authGroupId** | **string**| Retrieve the group information. | 
 **cached** | **bool**| If true, only return cached objects. | 
 **resolveNames** | **bool**| Resolve names of personas. | 
 **queryMemberOf** | **bool**| Enumerate all groups that a group is a member of. | 
 **zone** | **string**| Filter groups by zone. | 
 **provider** | **string**| Filter groups by provider. | 

### Return type

[**AuthGroups**](AuthGroups.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuthId**
> AuthId GetAuthId(ctx, )


Retrieve the current security token.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AuthId**](AuthId.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuthLdapTemplate**
> AuthLdapTemplates GetAuthLdapTemplate(ctx, authLdapTemplateId)


Retrieve the LDAP provider template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **authLdapTemplateId** | **string**| Retrieve the LDAP provider template. | 

### Return type

[**AuthLdapTemplates**](AuthLdapTemplates.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuthLdapTemplates**
> AuthLdapTemplatesExtended GetAuthLdapTemplates(ctx, )


List all LDAP provider templates.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AuthLdapTemplatesExtended**](AuthLdapTemplatesExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuthLogLevel**
> AuthLogLevel GetAuthLogLevel(ctx, )


Get the current authentications service and netlogon logging level.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AuthLogLevel**](AuthLogLevel.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuthNetgroup**
> AuthNetgroups GetAuthNetgroup(ctx, authNetgroupId, optional)


Retrieve the user information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **authNetgroupId** | **string**| Retrieve the user information. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authNetgroupId** | **string**| Retrieve the user information. | 
 **ignoreErrors** | **bool**| Ignore netgroup errors. | 
 **recursive** | **bool**| Perform recursive search. | 
 **zone** | **string**| Filter users by zone. | 
 **provider** | **string**| Filter users by provider. | 

### Return type

[**AuthNetgroups**](AuthNetgroups.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuthPrivileges**
> AuthPrivileges GetAuthPrivileges(ctx, )


List all privileges.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AuthPrivileges**](AuthPrivileges.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuthRole**
> AuthRoles GetAuthRole(ctx, authRoleId, optional)


Retrieve the role information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **authRoleId** | **string**| Retrieve the role information. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authRoleId** | **string**| Retrieve the role information. | 
 **resolveNames** | **bool**| Resolve names of personas. | 

### Return type

[**AuthRoles**](AuthRoles.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuthShells**
> AuthShells GetAuthShells(ctx, )


List all shells.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AuthShells**](AuthShells.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuthUser**
> AuthUsers GetAuthUser(ctx, authUserId, optional)


Retrieve the user information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **authUserId** | **string**| Retrieve the user information. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authUserId** | **string**| Retrieve the user information. | 
 **cached** | **bool**| If true, only return cached objects. | 
 **resolveNames** | **bool**| Resolve names of personas. | 
 **queryMemberOf** | **bool**| Enumerate all users that a group is a member of. | 
 **zone** | **string**| Filter users by zone. | 
 **provider** | **string**| Filter users by provider. | 

### Return type

[**AuthUsers**](AuthUsers.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuthWellknown**
> AuthWellknowns GetAuthWellknown(ctx, authWellknownId, optional)


Retrieve the wellknown persona.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **authWellknownId** | **string**| Retrieve the wellknown persona. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authWellknownId** | **string**| Retrieve the wellknown persona. | 
 **scope** | **string**| If specified as \&quot;effective\&quot; or not specified, all fields are returned.  If specified as \&quot;user\&quot;, only fields with non-default values are shown.  If specified as \&quot;default\&quot;, the original values are returned. | 

### Return type

[**AuthWellknowns**](AuthWellknowns.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuthWellknowns**
> AuthWellknowns GetAuthWellknowns(ctx, )


List all wellknown personas.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AuthWellknowns**](AuthWellknowns.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMappingDump**
> MappingDump GetMappingDump(ctx, optional)


Retrieve all identity mappings (uid, gid, sid, and on-disk) for the supplied source persona.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nocreate** | **bool**| Idmap should attempt to create missing identity mappings. | 
 **zone** | **string**| Optional zone. | 

### Return type

[**MappingDump**](MappingDump.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMappingIdentity**
> MappingIdentities GetMappingIdentity(ctx, mappingIdentityId, optional)


Retrieve all identity mappings (uid, gid, sid, and on-disk) for the supplied source persona.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **mappingIdentityId** | **string**| Retrieve all identity mappings (uid, gid, sid, and on-disk) for the supplied source persona. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mappingIdentityId** | **string**| Retrieve all identity mappings (uid, gid, sid, and on-disk) for the supplied source persona. | 
 **nocreate** | **bool**| Idmap should attempt to create missing identity mappings. | 
 **zone** | **string**| Optional zone. | 

### Return type

[**MappingIdentities**](MappingIdentities.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMappingUsersLookup**
> MappingUsersLookup GetMappingUsersLookup(ctx, optional)


Retrieve the user information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **primaryGid** | **int32**| The user&#39;s primary group ID. | 
 **uid** | **int32**| The user ID. | 
 **zone** | **string**| The zone the user belongs to. | 
 **gid** | [**[]int32**](int32.md)| The IDs of the groups that the user belongs to. | 
 **user** | **string**| The user name. | 
 **kerberosPrincipal** | **string**| The Kerberos principal name, of the form user@realm. | 

### Return type

[**MappingUsersLookup**](MappingUsersLookup.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMappingUsersRules**
> MappingUsersRules GetMappingUsersRules(ctx, optional)


Retrieve the user mapping rules.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone** | **string**| The zone to which the user mapping applies. | 

### Return type

[**MappingUsersRules**](MappingUsersRules.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProvidersAdsById**
> ProvidersAds GetProvidersAdsById(ctx, providersAdsId, optional)


Retrieve the ADS provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providersAdsId** | **string**| Retrieve the ADS provider. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providersAdsId** | **string**| Retrieve the ADS provider. | 
 **scope** | **string**| If specified as \&quot;effective\&quot; or not specified, all fields are returned.  If specified as \&quot;user\&quot;, only fields with non-default values are shown.  If specified as \&quot;default\&quot;, the original values are returned. | 

### Return type

[**ProvidersAds**](ProvidersAds.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProvidersFileById**
> ProvidersFile GetProvidersFileById(ctx, providersFileId, optional)


Retrieve the file provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providersFileId** | **string**| Retrieve the file provider. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providersFileId** | **string**| Retrieve the file provider. | 
 **scope** | **string**| If specified as \&quot;effective\&quot; or not specified, all fields are returned.  If specified as \&quot;user\&quot;, only fields with non-default values are shown.  If specified as \&quot;default\&quot;, the original values are returned. | 

### Return type

[**ProvidersFile**](ProvidersFile.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProvidersKrb5ById**
> ProvidersKrb5 GetProvidersKrb5ById(ctx, providersKrb5Id, optional)


Retrieve the KRB5 provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providersKrb5Id** | **string**| Retrieve the KRB5 provider. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providersKrb5Id** | **string**| Retrieve the KRB5 provider. | 
 **scope** | **string**| If specified as \&quot;effective\&quot; or not specified, all fields are returned.  If specified as \&quot;user\&quot;, only fields with non-default values are shown.  If specified as \&quot;default\&quot;, the original values are returned. | 

### Return type

[**ProvidersKrb5**](ProvidersKrb5.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProvidersLdapById**
> ProvidersLdap GetProvidersLdapById(ctx, providersLdapId, optional)


Retrieve the LDAP provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providersLdapId** | **string**| Retrieve the LDAP provider. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providersLdapId** | **string**| Retrieve the LDAP provider. | 
 **scope** | **string**| If specified as \&quot;effective\&quot; or not specified, all fields are returned.  If specified as \&quot;user\&quot;, only fields with non-default values are shown.  If specified as \&quot;default\&quot;, the original values are returned. | 

### Return type

[**ProvidersLdap**](ProvidersLdap.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProvidersLocal**
> ProvidersLocal GetProvidersLocal(ctx, optional)


List all local providers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **string**| If specified as \&quot;effective\&quot; or not specified, all fields are returned.  If specified as \&quot;user\&quot;, only fields with non-default values are shown.  If specified as \&quot;default\&quot;, the original values are returned. | 

### Return type

[**ProvidersLocal**](ProvidersLocal.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProvidersLocalById**
> ProvidersLocal GetProvidersLocalById(ctx, providersLocalId, optional)


Retrieve the local provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providersLocalId** | **string**| Retrieve the local provider. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providersLocalId** | **string**| Retrieve the local provider. | 
 **scope** | **string**| If specified as \&quot;effective\&quot; or not specified, all fields are returned.  If specified as \&quot;user\&quot;, only fields with non-default values are shown.  If specified as \&quot;default\&quot;, the original values are returned. | 

### Return type

[**ProvidersLocal**](ProvidersLocal.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProvidersNisById**
> ProvidersNis GetProvidersNisById(ctx, providersNisId, optional)


Retrieve the NIS provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providersNisId** | **string**| Retrieve the NIS provider. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providersNisId** | **string**| Retrieve the NIS provider. | 
 **scope** | **string**| If specified as \&quot;effective\&quot; or not specified, all fields are returned.  If specified as \&quot;user\&quot;, only fields with non-default values are shown.  If specified as \&quot;default\&quot;, the original values are returned. | 

### Return type

[**ProvidersNis**](ProvidersNis.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProvidersSummary**
> ProvidersSummary GetProvidersSummary(ctx, optional)


Retrieve the summary information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupnet** | **string**| Filter providers by groupnet. | 
 **zone** | **string**| Filter providers by zone. | 

### Return type

[**ProvidersSummary**](ProvidersSummary.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSettingsAcls**
> SettingsAcls GetSettingsAcls(ctx, optional)


Retrieve the ACL policy settings and preset configurations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **preset** | **string**| If specified the preset configuration values for all applicable ACL policies are returned. | 

### Return type

[**SettingsAcls**](SettingsAcls.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSettingsGlobal**
> SettingsGlobal GetSettingsGlobal(ctx, optional)


Retrieve the global settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **string**| If specified as \&quot;effective\&quot; or not specified, all fields are returned.  If specified as \&quot;user\&quot;, only fields with non-default values are shown.  If specified as \&quot;default\&quot;, the original values are returned. | 
 **zone** | **string**| Zone which contains any per-zone settings. | 

### Return type

[**SettingsGlobal**](SettingsGlobal.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSettingsKrb5Defaults**
> SettingsKrb5Defaults GetSettingsKrb5Defaults(ctx, )


Retrieve the krb5 settings.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SettingsKrb5Defaults**](SettingsKrb5Defaults.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSettingsKrb5Domain**
> SettingsKrb5Domains GetSettingsKrb5Domain(ctx, settingsKrb5DomainId)


View the krb5 domain settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **settingsKrb5DomainId** | **string**| View the krb5 domain settings. | 

### Return type

[**SettingsKrb5Domains**](SettingsKrb5Domains.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSettingsKrb5Realm**
> SettingsKrb5Realms GetSettingsKrb5Realm(ctx, settingsKrb5RealmId)


Retrieve the krb5 settings for realms.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **settingsKrb5RealmId** | **string**| Retrieve the krb5 settings for realms. | 

### Return type

[**SettingsKrb5Realms**](SettingsKrb5Realms.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSettingsMapping**
> SettingsMapping GetSettingsMapping(ctx, optional)


Retrieve the mapping settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **string**| If specified as \&quot;effective\&quot; or not specified, all fields are returned.  If specified as \&quot;user\&quot;, only fields with non-default values are shown.  If specified as \&quot;default\&quot;, the original values are returned. | 
 **zone** | **string**| Access zone which contains mapping settings. | 

### Return type

[**SettingsMapping**](SettingsMapping.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAuthGroups**
> AuthGroupsExtended ListAuthGroups(ctx, optional)


List all groups.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string**| Filter groups by domain. | 
 **zone** | **string**| Filter groups by zone. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 
 **cached** | **bool**| If true, only return cached objects. | 
 **resolveNames** | **bool**| Resolve names of personas. | 
 **filter** | **string**| Filter groups by name prefix. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **provider** | **string**| Filter groups by provider. | 
 **queryMemberOf** | **bool**| Enumerate all groups that a group is a member of. | 

### Return type

[**AuthGroupsExtended**](AuthGroupsExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAuthRoles**
> AuthRolesExtended ListAuthRoles(ctx, optional)


List all roles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string**| The field that will be used for sorting. | 
 **resolveNames** | **bool**| Filter users by zone. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **dir** | **string**| The direction of the sort. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 

### Return type

[**AuthRolesExtended**](AuthRolesExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAuthUsers**
> AuthUsersExtended ListAuthUsers(ctx, optional)


List all users.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string**| Filter users by domain. | 
 **zone** | **string**| Filter users by zone. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 
 **cached** | **bool**| If true, only return cached objects. | 
 **resolveNames** | **bool**| Resolve names of personas. | 
 **filter** | **string**| Filter users by name prefix. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **provider** | **string**| Filter users by provider. | 
 **queryMemberOf** | **bool**| Enumerate all users that a group is a member of. | 

### Return type

[**AuthUsersExtended**](AuthUsersExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProvidersAds**
> ProvidersAdsExtended ListProvidersAds(ctx, optional)


List all ADS providers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **string**| If specified as \&quot;effective\&quot; or not specified, all fields are returned.  If specified as \&quot;user\&quot;, only fields with non-default values are shown.  If specified as \&quot;default\&quot;, the original values are returned. | 

### Return type

[**ProvidersAdsExtended**](ProvidersAdsExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProvidersFile**
> ProvidersFile ListProvidersFile(ctx, optional)


List all file providers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **string**| If specified as \&quot;effective\&quot; or not specified, all fields are returned.  If specified as \&quot;user\&quot;, only fields with non-default values are shown.  If specified as \&quot;default\&quot;, the original values are returned. | 

### Return type

[**ProvidersFile**](ProvidersFile.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProvidersKrb5**
> ProvidersKrb5Extended ListProvidersKrb5(ctx, optional)


List all KRB5 providers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **string**| If specified as \&quot;effective\&quot; or not specified, all fields are returned.  If specified as \&quot;user\&quot;, only fields with non-default values are shown.  If specified as \&quot;default\&quot;, the original values are returned. | 

### Return type

[**ProvidersKrb5Extended**](ProvidersKrb5Extended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProvidersLdap**
> ProvidersLdap ListProvidersLdap(ctx, optional)


List all LDAP providers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **string**| If specified as \&quot;effective\&quot; or not specified, all fields are returned.  If specified as \&quot;user\&quot;, only fields with non-default values are shown.  If specified as \&quot;default\&quot;, the original values are returned. | 

### Return type

[**ProvidersLdap**](ProvidersLdap.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProvidersNis**
> ProvidersNisExtended ListProvidersNis(ctx, optional)


List all NIS providers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **string**| If specified as \&quot;effective\&quot; or not specified, all fields are returned.  If specified as \&quot;user\&quot;, only fields with non-default values are shown.  If specified as \&quot;default\&quot;, the original values are returned. | 

### Return type

[**ProvidersNisExtended**](ProvidersNisExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSettingsKrb5Domains**
> SettingsKrb5Domains ListSettingsKrb5Domains(ctx, )


Retrieve the krb5 settings for domains.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SettingsKrb5Domains**](SettingsKrb5Domains.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSettingsKrb5Realms**
> SettingsKrb5Realms ListSettingsKrb5Realms(ctx, )


Retrieve the krb5 settings for realms.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SettingsKrb5Realms**](SettingsKrb5Realms.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAuthGroup**
> UpdateAuthGroup(ctx, authGroup, authGroupId, optional)


Modify the group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **authGroup** | [**AuthGroup**](AuthGroup.md)|  | 
  **authGroupId** | **string**| Modify the group. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authGroup** | [**AuthGroup**](AuthGroup.md)|  | 
 **authGroupId** | **string**| Modify the group. | 
 **force** | **bool**| Changes to the group ID can result in loss of access to the file system. To mitigate this risk of lost access, the force option is required for group ID changes. | 
 **zone** | **string**| Optional zone. | 
 **provider** | **string**| Optional provider type. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAuthLogLevel**
> UpdateAuthLogLevel(ctx, authLogLevel)


Set the current authentication service and netlogon logging level.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **authLogLevel** | [**AuthLogLevelExtended**](AuthLogLevelExtended.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAuthRole**
> UpdateAuthRole(ctx, authRole, authRoleId)


Modify the role.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **authRole** | [**AuthRole**](AuthRole.md)|  | 
  **authRoleId** | **string**| Modify the role. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAuthUser**
> UpdateAuthUser(ctx, authUser, authUserId, optional)


Modify the user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **authUser** | [**AuthUser**](AuthUser.md)|  | 
  **authUserId** | **string**| Modify the user. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authUser** | [**AuthUser**](AuthUser.md)|  | 
 **authUserId** | **string**| Modify the user. | 
 **force** | **bool**| Changes to the user ID can result in loss of access to the file system. To mitigate this risk of lost access, the force option is required for user ID changes. | 
 **zone** | **string**| Optional zone. | 
 **provider** | **string**| Optional provider type. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMappingImport**
> UpdateMappingImport(ctx, mappingImport, optional)


Set or update a list of mappings between two personae.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **mappingImport** | [**MappingImport**](MappingImport.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mappingImport** | [**MappingImport**](MappingImport.md)|  | 
 **zone** | **string**| Optional zone. | 
 **replace** | **bool**| Specify whether existing mappings should be replaced. The default behavior is to leave existing mappings intact and return an error. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMappingUsersRules**
> UpdateMappingUsersRules(ctx, mappingUsersRules, optional)


Modify the user mapping rules.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **mappingUsersRules** | [**MappingUsersRulesExtended**](MappingUsersRulesExtended.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mappingUsersRules** | [**MappingUsersRulesExtended**](MappingUsersRulesExtended.md)|  | 
 **zone** | **string**| The zone to which the user mapping applies. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProvidersAdsById**
> UpdateProvidersAdsById(ctx, providersAdsIdParams, providersAdsId)


Modify the ADS provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providersAdsIdParams** | [**ProvidersAdsIdParams**](ProvidersAdsIdParams.md)|  | 
  **providersAdsId** | **string**| Modify the ADS provider. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProvidersFileById**
> UpdateProvidersFileById(ctx, providersFileIdParams, providersFileId)


Modify the file provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providersFileIdParams** | [**ProvidersFileIdParams**](ProvidersFileIdParams.md)|  | 
  **providersFileId** | **string**| Modify the file provider. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProvidersKrb5ById**
> UpdateProvidersKrb5ById(ctx, providersKrb5IdParams, providersKrb5Id)


Modify the KRB5 provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providersKrb5IdParams** | [**ProvidersKrb5IdParams**](ProvidersKrb5IdParams.md)|  | 
  **providersKrb5Id** | **string**| Modify the KRB5 provider. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProvidersLdapById**
> UpdateProvidersLdapById(ctx, providersLdapIdParams, providersLdapId, optional)


Modify the LDAP provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providersLdapIdParams** | [**ProvidersLdapIdParams**](ProvidersLdapIdParams.md)|  | 
  **providersLdapId** | **string**| Modify the LDAP provider. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providersLdapIdParams** | [**ProvidersLdapIdParams**](ProvidersLdapIdParams.md)|  | 
 **providersLdapId** | **string**| Modify the LDAP provider. | 
 **force** | **bool**| Ignore unresolvable server URIs. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProvidersLocalById**
> UpdateProvidersLocalById(ctx, providersLocalIdParams, providersLocalId)


Modify the local provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providersLocalIdParams** | [**ProvidersLocalIdParams**](ProvidersLocalIdParams.md)|  | 
  **providersLocalId** | **string**| Modify the local provider. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProvidersNisById**
> UpdateProvidersNisById(ctx, providersNisIdParams, providersNisId)


Modify the NIS provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **providersNisIdParams** | [**ProvidersNisIdParams**](ProvidersNisIdParams.md)|  | 
  **providersNisId** | **string**| Modify the NIS provider. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSettingsAcls**
> UpdateSettingsAcls(ctx, settingsAcls)


Modify cluster ACL policy settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **settingsAcls** | [**SettingsAclsExtended**](SettingsAclsExtended.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSettingsGlobal**
> UpdateSettingsGlobal(ctx, settingsGlobal, optional)


Modify the global settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **settingsGlobal** | [**SettingsGlobalGlobalSettings**](SettingsGlobalGlobalSettings.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **settingsGlobal** | [**SettingsGlobalGlobalSettings**](SettingsGlobalGlobalSettings.md)|  | 
 **zone** | **string**| Zone which contains any per-zone settings. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSettingsKrb5Defaults**
> UpdateSettingsKrb5Defaults(ctx, settingsKrb5Defaults)


Modify the krb5 settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **settingsKrb5Defaults** | [**SettingsKrb5DefaultsKrb5Settings**](SettingsKrb5DefaultsKrb5Settings.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSettingsKrb5Domain**
> UpdateSettingsKrb5Domain(ctx, settingsKrb5Domain, settingsKrb5DomainId)


Modify the krb5 domain settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **settingsKrb5Domain** | [**SettingsKrb5Domain**](SettingsKrb5Domain.md)|  | 
  **settingsKrb5DomainId** | **string**| Modify the krb5 domain settings. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSettingsKrb5Realm**
> UpdateSettingsKrb5Realm(ctx, settingsKrb5Realm, settingsKrb5RealmId)


Modify the krb5 realm settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **settingsKrb5Realm** | [**SettingsKrb5Realm**](SettingsKrb5Realm.md)|  | 
  **settingsKrb5RealmId** | **string**| Modify the krb5 realm settings. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSettingsMapping**
> UpdateSettingsMapping(ctx, settingsMapping, optional)


Modify the mapping settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **settingsMapping** | [**SettingsMappingMappingSettings**](SettingsMappingMappingSettings.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **settingsMapping** | [**SettingsMappingMappingSettings**](SettingsMappingMappingSettings.md)|  | 
 **zone** | **string**| Access zone which contains mapping settings. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

