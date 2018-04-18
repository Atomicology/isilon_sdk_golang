# \ProtocolsApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHdfsProxyuser**](ProtocolsApi.md#CreateHdfsProxyuser) | **Post** /platform/1/protocols/hdfs/proxyusers | 
[**CreateHdfsRack**](ProtocolsApi.md#CreateHdfsRack) | **Post** /platform/1/protocols/hdfs/racks | 
[**CreateNdmpSettingsPreferredIp**](ProtocolsApi.md#CreateNdmpSettingsPreferredIp) | **Post** /platform/4/protocols/ndmp/settings/preferred-ips | 
[**CreateNdmpSettingsVariable**](ProtocolsApi.md#CreateNdmpSettingsVariable) | **Post** /platform/3/protocols/ndmp/settings/variables/{NdmpSettingsVariableId} | 
[**CreateNdmpUser**](ProtocolsApi.md#CreateNdmpUser) | **Post** /platform/3/protocols/ndmp/users | 
[**CreateNfsAlias**](ProtocolsApi.md#CreateNfsAlias) | **Post** /platform/2/protocols/nfs/aliases | 
[**CreateNfsExport**](ProtocolsApi.md#CreateNfsExport) | **Post** /platform/4/protocols/nfs/exports | 
[**CreateNfsNetgroupCheckItem**](ProtocolsApi.md#CreateNfsNetgroupCheckItem) | **Post** /platform/3/protocols/nfs/netgroup/check | 
[**CreateNfsNetgroupFlushItem**](ProtocolsApi.md#CreateNfsNetgroupFlushItem) | **Post** /platform/3/protocols/nfs/netgroup/flush | 
[**CreateNfsNlmSessionsCheckItem**](ProtocolsApi.md#CreateNfsNlmSessionsCheckItem) | **Post** /platform/3/protocols/nfs/nlm/sessions-check | 
[**CreateNfsReloadItem**](ProtocolsApi.md#CreateNfsReloadItem) | **Post** /platform/3/protocols/nfs/reload | 
[**CreateNtpServer**](ProtocolsApi.md#CreateNtpServer) | **Post** /platform/3/protocols/ntp/servers | 
[**CreateSmbLogLevelFilter**](ProtocolsApi.md#CreateSmbLogLevelFilter) | **Post** /platform/3/protocols/smb/log-level/filters | 
[**CreateSmbShare**](ProtocolsApi.md#CreateSmbShare) | **Post** /platform/4/protocols/smb/shares | 
[**CreateSwiftAccount**](ProtocolsApi.md#CreateSwiftAccount) | **Post** /platform/3/protocols/swift/accounts | 
[**DeleteHdfsProxyuser**](ProtocolsApi.md#DeleteHdfsProxyuser) | **Delete** /platform/1/protocols/hdfs/proxyusers/{HdfsProxyuserId} | 
[**DeleteHdfsRack**](ProtocolsApi.md#DeleteHdfsRack) | **Delete** /platform/1/protocols/hdfs/racks/{HdfsRackId} | 
[**DeleteNdmpContextsBackupById**](ProtocolsApi.md#DeleteNdmpContextsBackupById) | **Delete** /platform/3/protocols/ndmp/contexts/backup/{NdmpContextsBackupId} | 
[**DeleteNdmpContextsBreById**](ProtocolsApi.md#DeleteNdmpContextsBreById) | **Delete** /platform/3/protocols/ndmp/contexts/bre/{NdmpContextsBreId} | 
[**DeleteNdmpContextsRestoreById**](ProtocolsApi.md#DeleteNdmpContextsRestoreById) | **Delete** /platform/3/protocols/ndmp/contexts/restore/{NdmpContextsRestoreId} | 
[**DeleteNdmpDumpdate**](ProtocolsApi.md#DeleteNdmpDumpdate) | **Delete** /platform/3/protocols/ndmp/dumpdates/{NdmpDumpdateId} | 
[**DeleteNdmpSession**](ProtocolsApi.md#DeleteNdmpSession) | **Delete** /platform/3/protocols/ndmp/sessions/{NdmpSessionId} | 
[**DeleteNdmpSettingsPreferredIp**](ProtocolsApi.md#DeleteNdmpSettingsPreferredIp) | **Delete** /platform/4/protocols/ndmp/settings/preferred-ips/{NdmpSettingsPreferredIpId} | 
[**DeleteNdmpSettingsVariable**](ProtocolsApi.md#DeleteNdmpSettingsVariable) | **Delete** /platform/3/protocols/ndmp/settings/variables/{NdmpSettingsVariableId} | 
[**DeleteNdmpUser**](ProtocolsApi.md#DeleteNdmpUser) | **Delete** /platform/3/protocols/ndmp/users/{NdmpUserId} | 
[**DeleteNfsAlias**](ProtocolsApi.md#DeleteNfsAlias) | **Delete** /platform/2/protocols/nfs/aliases/{NfsAliasId} | 
[**DeleteNfsExport**](ProtocolsApi.md#DeleteNfsExport) | **Delete** /platform/4/protocols/nfs/exports/{NfsExportId} | 
[**DeleteNfsNlmSession**](ProtocolsApi.md#DeleteNfsNlmSession) | **Delete** /platform/3/protocols/nfs/nlm/sessions/{NfsNlmSessionId} | 
[**DeleteNtpServer**](ProtocolsApi.md#DeleteNtpServer) | **Delete** /platform/3/protocols/ntp/servers/{NtpServerId} | 
[**DeleteNtpServers**](ProtocolsApi.md#DeleteNtpServers) | **Delete** /platform/3/protocols/ntp/servers | 
[**DeleteSmbLogLevelFilter**](ProtocolsApi.md#DeleteSmbLogLevelFilter) | **Delete** /platform/3/protocols/smb/log-level/filters/{SmbLogLevelFilterId} | 
[**DeleteSmbLogLevelFilters**](ProtocolsApi.md#DeleteSmbLogLevelFilters) | **Delete** /platform/3/protocols/smb/log-level/filters | 
[**DeleteSmbOpenfile**](ProtocolsApi.md#DeleteSmbOpenfile) | **Delete** /platform/1/protocols/smb/openfiles/{SmbOpenfileId} | 
[**DeleteSmbSession**](ProtocolsApi.md#DeleteSmbSession) | **Delete** /platform/1/protocols/smb/sessions/{SmbSessionId} | 
[**DeleteSmbSessionsComputerUser**](ProtocolsApi.md#DeleteSmbSessionsComputerUser) | **Delete** /platform/1/protocols/smb/sessions/{Computer}/{SmbSessionsComputerUser} | 
[**DeleteSmbShare**](ProtocolsApi.md#DeleteSmbShare) | **Delete** /platform/4/protocols/smb/shares/{SmbShareId} | 
[**DeleteSmbShares**](ProtocolsApi.md#DeleteSmbShares) | **Delete** /platform/4/protocols/smb/shares | 
[**DeleteSwiftAccount**](ProtocolsApi.md#DeleteSwiftAccount) | **Delete** /platform/3/protocols/swift/accounts/{SwiftAccountId} | 
[**GetFtpSettings**](ProtocolsApi.md#GetFtpSettings) | **Get** /platform/3/protocols/ftp/settings | 
[**GetHdfsLogLevel**](ProtocolsApi.md#GetHdfsLogLevel) | **Get** /platform/3/protocols/hdfs/log-level | 
[**GetHdfsProxyuser**](ProtocolsApi.md#GetHdfsProxyuser) | **Get** /platform/1/protocols/hdfs/proxyusers/{HdfsProxyuserId} | 
[**GetHdfsRack**](ProtocolsApi.md#GetHdfsRack) | **Get** /platform/1/protocols/hdfs/racks/{HdfsRackId} | 
[**GetHdfsRangerPluginSettings**](ProtocolsApi.md#GetHdfsRangerPluginSettings) | **Get** /platform/4/protocols/hdfs/ranger-plugin/settings | 
[**GetHdfsSettings**](ProtocolsApi.md#GetHdfsSettings) | **Get** /platform/4/protocols/hdfs/settings | 
[**GetHttpSettings**](ProtocolsApi.md#GetHttpSettings) | **Get** /platform/3/protocols/http/settings | 
[**GetNdmpContextsBackup**](ProtocolsApi.md#GetNdmpContextsBackup) | **Get** /platform/3/protocols/ndmp/contexts/backup | 
[**GetNdmpContextsBackupById**](ProtocolsApi.md#GetNdmpContextsBackupById) | **Get** /platform/3/protocols/ndmp/contexts/backup/{NdmpContextsBackupId} | 
[**GetNdmpContextsBre**](ProtocolsApi.md#GetNdmpContextsBre) | **Get** /platform/3/protocols/ndmp/contexts/bre | 
[**GetNdmpContextsBreById**](ProtocolsApi.md#GetNdmpContextsBreById) | **Get** /platform/3/protocols/ndmp/contexts/bre/{NdmpContextsBreId} | 
[**GetNdmpContextsRestore**](ProtocolsApi.md#GetNdmpContextsRestore) | **Get** /platform/3/protocols/ndmp/contexts/restore | 
[**GetNdmpContextsRestoreById**](ProtocolsApi.md#GetNdmpContextsRestoreById) | **Get** /platform/3/protocols/ndmp/contexts/restore/{NdmpContextsRestoreId} | 
[**GetNdmpDiagnostics**](ProtocolsApi.md#GetNdmpDiagnostics) | **Get** /platform/3/protocols/ndmp/diagnostics | 
[**GetNdmpDumpdate**](ProtocolsApi.md#GetNdmpDumpdate) | **Get** /platform/3/protocols/ndmp/dumpdates/{NdmpDumpdateId} | 
[**GetNdmpLogs**](ProtocolsApi.md#GetNdmpLogs) | **Get** /platform/3/protocols/ndmp/logs | 
[**GetNdmpSession**](ProtocolsApi.md#GetNdmpSession) | **Get** /platform/3/protocols/ndmp/sessions/{NdmpSessionId} | 
[**GetNdmpSessions**](ProtocolsApi.md#GetNdmpSessions) | **Get** /platform/3/protocols/ndmp/sessions | 
[**GetNdmpSettingsDmas**](ProtocolsApi.md#GetNdmpSettingsDmas) | **Get** /platform/3/protocols/ndmp/settings/dmas | 
[**GetNdmpSettingsGlobal**](ProtocolsApi.md#GetNdmpSettingsGlobal) | **Get** /platform/3/protocols/ndmp/settings/global | 
[**GetNdmpSettingsPreferredIp**](ProtocolsApi.md#GetNdmpSettingsPreferredIp) | **Get** /platform/4/protocols/ndmp/settings/preferred-ips/{NdmpSettingsPreferredIpId} | 
[**GetNdmpSettingsVariable**](ProtocolsApi.md#GetNdmpSettingsVariable) | **Get** /platform/3/protocols/ndmp/settings/variables/{NdmpSettingsVariableId} | 
[**GetNdmpUser**](ProtocolsApi.md#GetNdmpUser) | **Get** /platform/3/protocols/ndmp/users/{NdmpUserId} | 
[**GetNfsAlias**](ProtocolsApi.md#GetNfsAlias) | **Get** /platform/2/protocols/nfs/aliases/{NfsAliasId} | 
[**GetNfsCheck**](ProtocolsApi.md#GetNfsCheck) | **Get** /platform/2/protocols/nfs/check | 
[**GetNfsExport**](ProtocolsApi.md#GetNfsExport) | **Get** /platform/4/protocols/nfs/exports/{NfsExportId} | 
[**GetNfsExportsSummary**](ProtocolsApi.md#GetNfsExportsSummary) | **Get** /platform/2/protocols/nfs/exports-summary | 
[**GetNfsLogLevel**](ProtocolsApi.md#GetNfsLogLevel) | **Get** /platform/3/protocols/nfs/log-level | 
[**GetNfsNetgroup**](ProtocolsApi.md#GetNfsNetgroup) | **Get** /platform/3/protocols/nfs/netgroup | 
[**GetNfsNlmLocks**](ProtocolsApi.md#GetNfsNlmLocks) | **Get** /platform/2/protocols/nfs/nlm/locks | 
[**GetNfsNlmSession**](ProtocolsApi.md#GetNfsNlmSession) | **Get** /platform/3/protocols/nfs/nlm/sessions/{NfsNlmSessionId} | 
[**GetNfsNlmSessions**](ProtocolsApi.md#GetNfsNlmSessions) | **Get** /platform/3/protocols/nfs/nlm/sessions | 
[**GetNfsNlmWaiters**](ProtocolsApi.md#GetNfsNlmWaiters) | **Get** /platform/2/protocols/nfs/nlm/waiters | 
[**GetNfsSettingsExport**](ProtocolsApi.md#GetNfsSettingsExport) | **Get** /platform/2/protocols/nfs/settings/export | 
[**GetNfsSettingsGlobal**](ProtocolsApi.md#GetNfsSettingsGlobal) | **Get** /platform/3/protocols/nfs/settings/global | 
[**GetNfsSettingsZone**](ProtocolsApi.md#GetNfsSettingsZone) | **Get** /platform/2/protocols/nfs/settings/zone | 
[**GetNtpServer**](ProtocolsApi.md#GetNtpServer) | **Get** /platform/3/protocols/ntp/servers/{NtpServerId} | 
[**GetNtpSettings**](ProtocolsApi.md#GetNtpSettings) | **Get** /platform/3/protocols/ntp/settings | 
[**GetSmbLogLevel**](ProtocolsApi.md#GetSmbLogLevel) | **Get** /platform/3/protocols/smb/log-level | 
[**GetSmbLogLevelFilter**](ProtocolsApi.md#GetSmbLogLevelFilter) | **Get** /platform/3/protocols/smb/log-level/filters/{SmbLogLevelFilterId} | 
[**GetSmbOpenfiles**](ProtocolsApi.md#GetSmbOpenfiles) | **Get** /platform/1/protocols/smb/openfiles | 
[**GetSmbSessions**](ProtocolsApi.md#GetSmbSessions) | **Get** /platform/1/protocols/smb/sessions | 
[**GetSmbSettingsGlobal**](ProtocolsApi.md#GetSmbSettingsGlobal) | **Get** /platform/3/protocols/smb/settings/global | 
[**GetSmbSettingsShare**](ProtocolsApi.md#GetSmbSettingsShare) | **Get** /platform/3/protocols/smb/settings/share | 
[**GetSmbShare**](ProtocolsApi.md#GetSmbShare) | **Get** /platform/4/protocols/smb/shares/{SmbShareId} | 
[**GetSmbSharesSummary**](ProtocolsApi.md#GetSmbSharesSummary) | **Get** /platform/1/protocols/smb/shares-summary | 
[**GetSnmpSettings**](ProtocolsApi.md#GetSnmpSettings) | **Get** /platform/5/protocols/snmp/settings | 
[**GetSwiftAccount**](ProtocolsApi.md#GetSwiftAccount) | **Get** /platform/3/protocols/swift/accounts/{SwiftAccountId} | 
[**ListHdfsProxyusers**](ProtocolsApi.md#ListHdfsProxyusers) | **Get** /platform/1/protocols/hdfs/proxyusers | 
[**ListHdfsRacks**](ProtocolsApi.md#ListHdfsRacks) | **Get** /platform/1/protocols/hdfs/racks | 
[**ListNdmpSettingsPreferredIps**](ProtocolsApi.md#ListNdmpSettingsPreferredIps) | **Get** /platform/4/protocols/ndmp/settings/preferred-ips | 
[**ListNdmpUsers**](ProtocolsApi.md#ListNdmpUsers) | **Get** /platform/3/protocols/ndmp/users | 
[**ListNfsAliases**](ProtocolsApi.md#ListNfsAliases) | **Get** /platform/2/protocols/nfs/aliases | 
[**ListNfsExports**](ProtocolsApi.md#ListNfsExports) | **Get** /platform/4/protocols/nfs/exports | 
[**ListNtpServers**](ProtocolsApi.md#ListNtpServers) | **Get** /platform/3/protocols/ntp/servers | 
[**ListSmbLogLevelFilters**](ProtocolsApi.md#ListSmbLogLevelFilters) | **Get** /platform/3/protocols/smb/log-level/filters | 
[**ListSmbShares**](ProtocolsApi.md#ListSmbShares) | **Get** /platform/4/protocols/smb/shares | 
[**ListSwiftAccounts**](ProtocolsApi.md#ListSwiftAccounts) | **Get** /platform/3/protocols/swift/accounts | 
[**UpdateFtpSettings**](ProtocolsApi.md#UpdateFtpSettings) | **Put** /platform/3/protocols/ftp/settings | 
[**UpdateHdfsLogLevel**](ProtocolsApi.md#UpdateHdfsLogLevel) | **Put** /platform/3/protocols/hdfs/log-level | 
[**UpdateHdfsProxyuser**](ProtocolsApi.md#UpdateHdfsProxyuser) | **Put** /platform/1/protocols/hdfs/proxyusers/{HdfsProxyuserId} | 
[**UpdateHdfsRack**](ProtocolsApi.md#UpdateHdfsRack) | **Put** /platform/1/protocols/hdfs/racks/{HdfsRackId} | 
[**UpdateHdfsRangerPluginSettings**](ProtocolsApi.md#UpdateHdfsRangerPluginSettings) | **Put** /platform/4/protocols/hdfs/ranger-plugin/settings | 
[**UpdateHdfsSettings**](ProtocolsApi.md#UpdateHdfsSettings) | **Put** /platform/4/protocols/hdfs/settings | 
[**UpdateHttpSettings**](ProtocolsApi.md#UpdateHttpSettings) | **Put** /platform/3/protocols/http/settings | 
[**UpdateNdmpDiagnostics**](ProtocolsApi.md#UpdateNdmpDiagnostics) | **Put** /platform/3/protocols/ndmp/diagnostics | 
[**UpdateNdmpSettingsGlobal**](ProtocolsApi.md#UpdateNdmpSettingsGlobal) | **Put** /platform/3/protocols/ndmp/settings/global | 
[**UpdateNdmpSettingsPreferredIp**](ProtocolsApi.md#UpdateNdmpSettingsPreferredIp) | **Put** /platform/4/protocols/ndmp/settings/preferred-ips/{NdmpSettingsPreferredIpId} | 
[**UpdateNdmpSettingsVariable**](ProtocolsApi.md#UpdateNdmpSettingsVariable) | **Put** /platform/3/protocols/ndmp/settings/variables/{NdmpSettingsVariableId} | 
[**UpdateNdmpUser**](ProtocolsApi.md#UpdateNdmpUser) | **Put** /platform/3/protocols/ndmp/users/{NdmpUserId} | 
[**UpdateNfsAlias**](ProtocolsApi.md#UpdateNfsAlias) | **Put** /platform/2/protocols/nfs/aliases/{NfsAliasId} | 
[**UpdateNfsExport**](ProtocolsApi.md#UpdateNfsExport) | **Put** /platform/4/protocols/nfs/exports/{NfsExportId} | 
[**UpdateNfsLogLevel**](ProtocolsApi.md#UpdateNfsLogLevel) | **Put** /platform/3/protocols/nfs/log-level | 
[**UpdateNfsNetgroup**](ProtocolsApi.md#UpdateNfsNetgroup) | **Put** /platform/3/protocols/nfs/netgroup | 
[**UpdateNfsSettingsExport**](ProtocolsApi.md#UpdateNfsSettingsExport) | **Put** /platform/2/protocols/nfs/settings/export | 
[**UpdateNfsSettingsGlobal**](ProtocolsApi.md#UpdateNfsSettingsGlobal) | **Put** /platform/3/protocols/nfs/settings/global | 
[**UpdateNfsSettingsZone**](ProtocolsApi.md#UpdateNfsSettingsZone) | **Put** /platform/2/protocols/nfs/settings/zone | 
[**UpdateNtpServer**](ProtocolsApi.md#UpdateNtpServer) | **Put** /platform/3/protocols/ntp/servers/{NtpServerId} | 
[**UpdateNtpSettings**](ProtocolsApi.md#UpdateNtpSettings) | **Put** /platform/3/protocols/ntp/settings | 
[**UpdateSmbLogLevel**](ProtocolsApi.md#UpdateSmbLogLevel) | **Put** /platform/3/protocols/smb/log-level | 
[**UpdateSmbSettingsGlobal**](ProtocolsApi.md#UpdateSmbSettingsGlobal) | **Put** /platform/3/protocols/smb/settings/global | 
[**UpdateSmbSettingsShare**](ProtocolsApi.md#UpdateSmbSettingsShare) | **Put** /platform/3/protocols/smb/settings/share | 
[**UpdateSmbShare**](ProtocolsApi.md#UpdateSmbShare) | **Put** /platform/4/protocols/smb/shares/{SmbShareId} | 
[**UpdateSnmpSettings**](ProtocolsApi.md#UpdateSnmpSettings) | **Put** /platform/5/protocols/snmp/settings | 
[**UpdateSwiftAccount**](ProtocolsApi.md#UpdateSwiftAccount) | **Put** /platform/3/protocols/swift/accounts/{SwiftAccountId} | 


# **CreateHdfsProxyuser**
> CreateResponse CreateHdfsProxyuser(ctx, hdfsProxyuser, optional)


Create a new HDFS proxyuser.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hdfsProxyuser** | [**HdfsProxyuserCreateParams**](HdfsProxyuserCreateParams.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hdfsProxyuser** | [**HdfsProxyuserCreateParams**](HdfsProxyuserCreateParams.md)|  | 
 **zone** | **string**| Access zone which contains HDFS proxyuser. | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateHdfsRack**
> CreateResponse CreateHdfsRack(ctx, hdfsRack, optional)


Create a new HDFS rack.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hdfsRack** | [**HdfsRackCreateParams**](HdfsRackCreateParams.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hdfsRack** | [**HdfsRackCreateParams**](HdfsRackCreateParams.md)|  | 
 **zone** | **string**| Access zone which contains HDFS rack. | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNdmpSettingsPreferredIp**
> Empty CreateNdmpSettingsPreferredIp(ctx, ndmpSettingsPreferredIp)


Create a preferred ip preference.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ndmpSettingsPreferredIp** | [**NdmpSettingsPreferredIpCreateParams**](NdmpSettingsPreferredIpCreateParams.md)|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNdmpSettingsVariable**
> Empty CreateNdmpSettingsVariable(ctx, ndmpSettingsVariable, ndmpSettingsVariableId)


Create a preferred NDMP environment variable.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ndmpSettingsVariable** | [**NdmpSettingsVariableCreateParams**](NdmpSettingsVariableCreateParams.md)|  | 
  **ndmpSettingsVariableId** | **string**| Create a preferred NDMP environment variable. | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNdmpUser**
> Empty CreateNdmpUser(ctx, ndmpUser)


Created a new user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ndmpUser** | [**NdmpUserCreateParams**](NdmpUserCreateParams.md)|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNfsAlias**
> CreateNfsAliasResponse CreateNfsAlias(ctx, nfsAlias, optional)


Create a new NFS alias.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **nfsAlias** | [**NfsAliasCreateParams**](NfsAliasCreateParams.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nfsAlias** | [**NfsAliasCreateParams**](NfsAliasCreateParams.md)|  | 
 **zone** | **string**| Access zone | 

### Return type

[**CreateNfsAliasResponse**](CreateNfsAliasResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNfsExport**
> CreateQuotaReportResponse CreateNfsExport(ctx, nfsExport, optional)


Create a new NFS export.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **nfsExport** | [**NfsExportCreateParams**](NfsExportCreateParams.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nfsExport** | [**NfsExportCreateParams**](NfsExportCreateParams.md)|  | 
 **force** | **bool**| If true, the export will be created even if it conflicts with another export. | 
 **ignoreUnresolvableHosts** | **bool**| Ignore unresolvable hosts. | 
 **zone** | **string**| Access zone | 
 **ignoreConflicts** | **bool**| Ignore conflicts with existing exports. | 
 **ignoreBadPaths** | **bool**| Ignore nonexistent or otherwise bad paths. | 
 **ignoreBadAuth** | **bool**| Ignore invalid users. | 

### Return type

[**CreateQuotaReportResponse**](CreateQuotaReportResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNfsNetgroupCheckItem**
> Empty CreateNfsNetgroupCheckItem(ctx, nfsNetgroupCheckItem, optional)


Update the NFS netgroups in the cache.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **nfsNetgroupCheckItem** | [**Empty**](Empty.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nfsNetgroupCheckItem** | [**Empty**](Empty.md)|  | 
 **host** | **string**| IP address of node to update. If unspecified, the local nodes cache is updated. | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNfsNetgroupFlushItem**
> Empty CreateNfsNetgroupFlushItem(ctx, nfsNetgroupFlushItem, optional)


Flush the NFS netgroups in the cache.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **nfsNetgroupFlushItem** | [**Empty**](Empty.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nfsNetgroupFlushItem** | [**Empty**](Empty.md)|  | 
 **host** | **string**| IP address of node to flush. If unspecified, all nodes on the cluster are flushed. | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNfsNlmSessionsCheckItem**
> CreateNfsNlmSessionsCheckItemResponse CreateNfsNlmSessionsCheckItem(ctx, nfsNlmSessionsCheckItem, optional)


Perform an active scan for lost NFSv3 locks.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **nfsNlmSessionsCheckItem** | [**Empty**](Empty.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nfsNlmSessionsCheckItem** | [**Empty**](Empty.md)|  | 
 **clusterIp** | **string**| An IP address for which NSM has client records | 
 **zone** | **string**| Represents an extant auth zone | 

### Return type

[**CreateNfsNlmSessionsCheckItemResponse**](CreateNfsNlmSessionsCheckItemResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNfsReloadItem**
> Empty CreateNfsReloadItem(ctx, nfsReloadItem, optional)


Reload default NFS export configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **nfsReloadItem** | [**Empty**](Empty.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nfsReloadItem** | [**Empty**](Empty.md)|  | 
 **zone** | **string**| Access zone | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNtpServer**
> CreateResponse CreateNtpServer(ctx, ntpServer)


Create an NTP server entry.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ntpServer** | [**NtpServerCreateParams**](NtpServerCreateParams.md)|  | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSmbLogLevelFilter**
> CreateAuthRefreshItemResponse CreateSmbLogLevelFilter(ctx, smbLogLevelFilter)


Add an SMB log filter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **smbLogLevelFilter** | [**SmbLogLevelFilter**](SmbLogLevelFilter.md)|  | 

### Return type

[**CreateAuthRefreshItemResponse**](CreateAuthRefreshItemResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSmbShare**
> CreateResponse CreateSmbShare(ctx, smbShare, optional)


Create a new share.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **smbShare** | [**SmbShareCreateParams**](SmbShareCreateParams.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smbShare** | [**SmbShareCreateParams**](SmbShareCreateParams.md)|  | 
 **zone** | **string**| Zone which contains this share. | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSwiftAccount**
> CreateResponse CreateSwiftAccount(ctx, swiftAccount, optional)


Create a new Swift account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **swiftAccount** | [**SwiftAccount**](SwiftAccount.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **swiftAccount** | [**SwiftAccount**](SwiftAccount.md)|  | 
 **zone** | **string**| Access zone which contains Swift account. | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHdfsProxyuser**
> DeleteHdfsProxyuser(ctx, hdfsProxyuserId, optional)


Delete an HDFS proxyuser.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hdfsProxyuserId** | **string**| Delete an HDFS proxyuser. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hdfsProxyuserId** | **string**| Delete an HDFS proxyuser. | 
 **zone** | **string**| Access zone which contains HDFS proxyuser. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHdfsRack**
> DeleteHdfsRack(ctx, hdfsRackId, optional)


Delete the HDFS rack.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hdfsRackId** | **string**| Delete the HDFS rack. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hdfsRackId** | **string**| Delete the HDFS rack. | 
 **zone** | **string**| Access zone which contains HDFS rack. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNdmpContextsBackupById**
> DeleteNdmpContextsBackupById(ctx, ndmpContextsBackupId)


Delete a backup context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ndmpContextsBackupId** | **string**| Delete a backup context | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNdmpContextsBreById**
> DeleteNdmpContextsBreById(ctx, ndmpContextsBreId)


Delete a NDMP BRE context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ndmpContextsBreId** | **string**| Delete a NDMP BRE context | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNdmpContextsRestoreById**
> DeleteNdmpContextsRestoreById(ctx, ndmpContextsRestoreId)


Delete a restore context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ndmpContextsRestoreId** | **string**| Delete a restore context | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNdmpDumpdate**
> DeleteNdmpDumpdate(ctx, ndmpDumpdateId, optional)


Delete dumpdates entries of a specified path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ndmpDumpdateId** | **string**| Delete dumpdates entries of a specified path. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ndmpDumpdateId** | **string**| Delete dumpdates entries of a specified path. | 
 **level** | **int32**| Level is an input from 0 to 10. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNdmpSession**
> DeleteNdmpSession(ctx, ndmpSessionId, optional)


Delete the ndmp session.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ndmpSessionId** | **string**| Delete the ndmp session. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ndmpSessionId** | **string**| Delete the ndmp session. | 
 **lnn** | **string**| Logical node number. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNdmpSettingsPreferredIp**
> DeleteNdmpSettingsPreferredIp(ctx, ndmpSettingsPreferredIpId)


Delete a preferred ip preference.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ndmpSettingsPreferredIpId** | **string**| Delete a preferred ip preference. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNdmpSettingsVariable**
> DeleteNdmpSettingsVariable(ctx, ndmpSettingsVariableId, optional)


Delete preferred environment variable entries

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ndmpSettingsVariableId** | **string**| Delete preferred environment variable entries | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ndmpSettingsVariableId** | **string**| Delete preferred environment variable entries | 
 **name** | **string**| Name of the variable to delete. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNdmpUser**
> DeleteNdmpUser(ctx, ndmpUserId)


Delete the user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ndmpUserId** | **string**| Delete the user. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNfsAlias**
> DeleteNfsAlias(ctx, nfsAliasId, optional)


Delete the export.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **nfsAliasId** | **string**| Delete the export. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nfsAliasId** | **string**| Delete the export. | 
 **zone** | **string**| Access zone | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNfsExport**
> DeleteNfsExport(ctx, nfsExportId, optional)


Delete the export.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **nfsExportId** | **string**| Delete the export. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nfsExportId** | **string**| Delete the export. | 
 **zone** | **string**| Access zone | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNfsNlmSession**
> DeleteNfsNlmSession(ctx, nfsNlmSessionId, optional)


Delete all lock state for this host.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **nfsNlmSessionId** | **string**| Delete all lock state for this host. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nfsNlmSessionId** | **string**| Delete all lock state for this host. | 
 **clusterIp** | **string**| An IP address for which NSM has client records | 
 **zone** | **string**| Represents an extant auth zone | 
 **refresh** | **bool**| if set to true, the client will be given a chance to reclaim its locks before they are destroyed | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNtpServer**
> DeleteNtpServer(ctx, ntpServerId)


Delete an NTP server entry.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ntpServerId** | **string**| Delete an NTP server entry. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNtpServers**
> DeleteNtpServers(ctx, )


Delete all NTP server entries.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSmbLogLevelFilter**
> DeleteSmbLogLevelFilter(ctx, smbLogLevelFilterId)


Delete log filter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **smbLogLevelFilterId** | **string**| Delete log filter. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSmbLogLevelFilters**
> DeleteSmbLogLevelFilters(ctx, optional)


Delete existing SMB log filters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **level** | **string**| Valid SMB logging levels | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSmbOpenfile**
> DeleteSmbOpenfile(ctx, smbOpenfileId)


Close the file in the SMB server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **smbOpenfileId** | **string**| Close the file in the SMB server. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSmbSession**
> DeleteSmbSession(ctx, smbSessionId)


Close the SMB session.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **smbSessionId** | **string**| Close the SMB session. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSmbSessionsComputerUser**
> DeleteSmbSessionsComputerUser(ctx, smbSessionsComputerUser, computer)


Close the SMB session.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **smbSessionsComputerUser** | **string**| Close the SMB session. | 
  **computer** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSmbShare**
> DeleteSmbShare(ctx, smbShareId, optional)


Delete the share.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **smbShareId** | **string**| Delete the share. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smbShareId** | **string**| Delete the share. | 
 **zone** | **string**| Zone which contains this share. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSmbShares**
> DeleteSmbShares(ctx, )


Delete multiple smb shares.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSwiftAccount**
> DeleteSwiftAccount(ctx, swiftAccountId, optional)


Delete a Swift account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **swiftAccountId** | **string**| Delete a Swift account. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **swiftAccountId** | **string**| Delete a Swift account. | 
 **zone** | **string**| Access zone which contains Swift account. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFtpSettings**
> FtpSettings GetFtpSettings(ctx, )


Retrieve the FTP settings.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**FtpSettings**](FtpSettings.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHdfsLogLevel**
> HdfsLogLevel GetHdfsLogLevel(ctx, )


Retrieve the HDFS service log-level.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**HdfsLogLevel**](HdfsLogLevel.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHdfsProxyuser**
> HdfsProxyusers GetHdfsProxyuser(ctx, hdfsProxyuserId, optional)


View the proxyuser.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hdfsProxyuserId** | **string**| View the proxyuser. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hdfsProxyuserId** | **string**| View the proxyuser. | 
 **zone** | **string**| Access zone which contains HDFS proxyuser. | 

### Return type

[**HdfsProxyusers**](HdfsProxyusers.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHdfsRack**
> HdfsRacks GetHdfsRack(ctx, hdfsRackId, optional)


Retrieve the HDFS rack.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hdfsRackId** | **string**| Retrieve the HDFS rack. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hdfsRackId** | **string**| Retrieve the HDFS rack. | 
 **zone** | **string**| Access zone which contains HDFS rack. | 

### Return type

[**HdfsRacks**](HdfsRacks.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHdfsRangerPluginSettings**
> HdfsRangerPluginSettings GetHdfsRangerPluginSettings(ctx, optional)


Retrieve HDFS ranger-plugin properties.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone** | **string**| Access zone which contains HDFS ranger-plugin settings. | 

### Return type

[**HdfsRangerPluginSettings**](HdfsRangerPluginSettings.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHdfsSettings**
> HdfsSettings GetHdfsSettings(ctx, optional)


Retrieve HDFS properties.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone** | **string**| Access zone which contains HDFS settings. | 

### Return type

[**HdfsSettings**](HdfsSettings.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHttpSettings**
> HttpSettings GetHttpSettings(ctx, )


Retrieve HTTP properties.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**HttpSettings**](HttpSettings.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNdmpContextsBackup**
> NdmpContextsBackupExtended GetNdmpContextsBackup(ctx, optional)


Get List of NDMP Backup Contexts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 

### Return type

[**NdmpContextsBackupExtended**](NdmpContextsBackupExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNdmpContextsBackupById**
> NdmpContextsBackup GetNdmpContextsBackupById(ctx, ndmpContextsBackupId)


View a NDMP backup context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ndmpContextsBackupId** | **string**| View a NDMP backup context | 

### Return type

[**NdmpContextsBackup**](NdmpContextsBackup.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNdmpContextsBre**
> NdmpContextsBreExtended GetNdmpContextsBre(ctx, optional)


Get list of NDMP BRE Contexts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 

### Return type

[**NdmpContextsBreExtended**](NdmpContextsBreExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNdmpContextsBreById**
> NdmpContextsBre GetNdmpContextsBreById(ctx, ndmpContextsBreId)


View a NDMP BRE context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ndmpContextsBreId** | **string**| View a NDMP BRE context | 

### Return type

[**NdmpContextsBre**](NdmpContextsBre.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNdmpContextsRestore**
> NdmpContextsBackupExtended GetNdmpContextsRestore(ctx, optional)


Get List of NDMP Restore Contexts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 

### Return type

[**NdmpContextsBackupExtended**](NdmpContextsBackupExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNdmpContextsRestoreById**
> NdmpContextsBackup GetNdmpContextsRestoreById(ctx, ndmpContextsRestoreId)


View a NDMP restore context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ndmpContextsRestoreId** | **string**| View a NDMP restore context | 

### Return type

[**NdmpContextsBackup**](NdmpContextsBackup.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNdmpDiagnostics**
> NdmpDiagnostics GetNdmpDiagnostics(ctx, )


List ndmp diagnostics settings.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NdmpDiagnostics**](NdmpDiagnostics.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNdmpDumpdate**
> NdmpDumpdates GetNdmpDumpdate(ctx, ndmpDumpdateId, optional)


List of dumpdates entries.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ndmpDumpdateId** | **string**| List of dumpdates entries. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ndmpDumpdateId** | **string**| List of dumpdates entries. | 
 **sort** | **string**| The field that will be used for sorting. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 
 **level** | **int32**| Filter by dumpdate level. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **path** | **string**| Filter by file path. | 
 **dir** | **string**| The direction of the sort. | 

### Return type

[**NdmpDumpdates**](NdmpDumpdates.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNdmpLogs**
> NdmpLogs GetNdmpLogs(ctx, optional)


Get NDMP logs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lnn** | **string**| Logical node number. | 
 **page** | **int32**| The page number of the NDMP logs file. | 
 **pagesize** | **int32**| The page size of each page of the NDMP log file. | 

### Return type

[**NdmpLogs**](NdmpLogs.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNdmpSession**
> NdmpSessions GetNdmpSession(ctx, ndmpSessionId, optional)


Retrieve the ndmp session.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ndmpSessionId** | **string**| Retrieve the ndmp session. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ndmpSessionId** | **string**| Retrieve the ndmp session. | 
 **lnn** | **string**| Logical node number. | 

### Return type

[**NdmpSessions**](NdmpSessions.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNdmpSessions**
> NdmpSessionsExtended GetNdmpSessions(ctx, optional)


List all ndmp sessions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consolidate** | **bool**| Combine sessions in the same context. | 
 **node** | **string**| Only return sessions of the node. | 
 **session** | **string**| Only return the specified session. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 

### Return type

[**NdmpSessionsExtended**](NdmpSessionsExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNdmpSettingsDmas**
> NdmpSettingsDmas GetNdmpSettingsDmas(ctx, )


List of supported dma vendors.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NdmpSettingsDmas**](NdmpSettingsDmas.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNdmpSettingsGlobal**
> NdmpSettingsGlobal GetNdmpSettingsGlobal(ctx, )


List global ndmp settings.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NdmpSettingsGlobal**](NdmpSettingsGlobal.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNdmpSettingsPreferredIp**
> NdmpSettingsPreferredIps GetNdmpSettingsPreferredIp(ctx, ndmpSettingsPreferredIpId)


Get one preference by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ndmpSettingsPreferredIpId** | **string**| Get one preference by id. | 

### Return type

[**NdmpSettingsPreferredIps**](NdmpSettingsPreferredIps.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNdmpSettingsVariable**
> NdmpSettingsVariables GetNdmpSettingsVariable(ctx, ndmpSettingsVariableId, optional)


List of preferred environment variables.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ndmpSettingsVariableId** | **string**| List of preferred environment variables. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ndmpSettingsVariableId** | **string**| List of preferred environment variables. | 
 **path** | **string**| Return variables of the path. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 

### Return type

[**NdmpSettingsVariables**](NdmpSettingsVariables.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNdmpUser**
> NdmpUsers GetNdmpUser(ctx, ndmpUserId)


Retrieve the user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ndmpUserId** | **string**| Retrieve the user. | 

### Return type

[**NdmpUsers**](NdmpUsers.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNfsAlias**
> NfsAliases GetNfsAlias(ctx, nfsAliasId, optional)


Retrieve export information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **nfsAliasId** | **string**| Retrieve export information. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nfsAliasId** | **string**| Retrieve export information. | 
 **scope** | **string**| When specified as &#39;effective&#39;, or not specified, all fields are returned. When specified as &#39;user&#39;, only fields with non-default values are shown. When specified as &#39;default&#39;, the original values are returned. | 
 **check** | **bool**| Check for conflicts when viewing alias. | 
 **zone** | **string**| Access zone | 

### Return type

[**NfsAliases**](NfsAliases.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNfsCheck**
> NfsCheckExtended GetNfsCheck(ctx, optional)


Retrieve NFS export validation information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ignoreBadPaths** | **bool**| Ignore nonexistent or otherwise bad paths. | 
 **ignoreBadAuth** | **bool**| Ignore invalid users. | 
 **zone** | **string**| Access zone | 
 **ignoreUnresolvableHosts** | **bool**| Ignore unresolvable hosts. | 

### Return type

[**NfsCheckExtended**](NfsCheckExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNfsExport**
> NfsExports GetNfsExport(ctx, nfsExportId, optional)


Retrieve export information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **nfsExportId** | **string**| Retrieve export information. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nfsExportId** | **string**| Retrieve export information. | 
 **scope** | **string**| When specified as &#39;effective&#39;, or not specified, all fields are returned. When specified as &#39;user&#39;, only fields with non-default values are shown. When specified as &#39;default&#39;, the original values are returned. | 
 **zone** | **string**| Access zone | 

### Return type

[**NfsExports**](NfsExports.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNfsExportsSummary**
> NfsExportsSummary GetNfsExportsSummary(ctx, optional)


Retrieve NFS export summary information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone** | **string**| Access zone | 

### Return type

[**NfsExportsSummary**](NfsExportsSummary.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNfsLogLevel**
> NfsLogLevel GetNfsLogLevel(ctx, )


Get the current NFS service logging level.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NfsLogLevel**](NfsLogLevel.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNfsNetgroup**
> NfsNetgroup GetNfsNetgroup(ctx, optional)


Get the current NFS netgroup cache settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **host** | **string**| Host to retrieve netgroup cache settings from. | 

### Return type

[**NfsNetgroup**](NfsNetgroup.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNfsNlmLocks**
> NfsNlmLocks GetNfsNlmLocks(ctx, optional)


List all NLM locks.

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
 **created** | **string**| Return locks created after the specified unix epoch time. | 
 **lin** | **string**| Filter locks by the specified LIN in /ifs that is locked. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 
 **client** | **string**| Filter locks by the specified client host name and IP address. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **clientId** | **string**| Filter locks by the specified client ID. | 
 **path** | **string**| Filter locks by the specified path under /ifs. | 
 **dir** | **string**| The direction of the sort. | 

### Return type

[**NfsNlmLocks**](NfsNlmLocks.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNfsNlmSession**
> NfsNlmSessions GetNfsNlmSession(ctx, nfsNlmSessionId, optional)


Retrieve all lock state for a single client.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **nfsNlmSessionId** | **string**| Retrieve all lock state for a single client. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nfsNlmSessionId** | **string**| Retrieve all lock state for a single client. | 
 **clusterIp** | **string**| An IP address for which NSM has client records | 
 **zone** | **string**| Represents an extant auth zone | 

### Return type

[**NfsNlmSessions**](NfsNlmSessions.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNfsNlmSessions**
> NfsNlmSessionsExtended GetNfsNlmSessions(ctx, optional)


List all NSM clients (optionally filtered by either zone or IP)

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
 **ip** | **string**| An IP address for which NSM has client records | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **zone** | **string**| Represents an extant auth zone | 
 **dir** | **string**| The direction of the sort. | 

### Return type

[**NfsNlmSessionsExtended**](NfsNlmSessionsExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNfsNlmWaiters**
> NfsNlmWaiters GetNfsNlmWaiters(ctx, optional)


List all NLM lock waiters.

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
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **dir** | **string**| The direction of the sort. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 

### Return type

[**NfsNlmWaiters**](NfsNlmWaiters.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNfsSettingsExport**
> NfsSettingsExport GetNfsSettingsExport(ctx, optional)


Retrieve export information.

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
 **zone** | **string**| Access zone | 

### Return type

[**NfsSettingsExport**](NfsSettingsExport.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNfsSettingsGlobal**
> NfsSettingsGlobal GetNfsSettingsGlobal(ctx, optional)


Retrieve the NFS configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **string**| When specified as &#39;effective&#39;, or not specified, all fields are returned. When specified as &#39;user&#39;, only fields with non-default values are shown. When specified as &#39;default&#39;, the original values are returned. | 

### Return type

[**NfsSettingsGlobal**](NfsSettingsGlobal.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNfsSettingsZone**
> NfsSettingsZone GetNfsSettingsZone(ctx, optional)


Retrieve the NFS server settings for this zone.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone** | **string**| Access zone | 

### Return type

[**NfsSettingsZone**](NfsSettingsZone.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNtpServer**
> NtpServers GetNtpServer(ctx, ntpServerId)


Retrieve one NTP server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ntpServerId** | **string**| Retrieve one NTP server. | 

### Return type

[**NtpServers**](NtpServers.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNtpSettings**
> NtpSettings GetNtpSettings(ctx, )


Retrieve the NTP settings.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NtpSettings**](NtpSettings.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSmbLogLevel**
> SmbLogLevel GetSmbLogLevel(ctx, )


Get the current SMB logging level.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SmbLogLevel**](SmbLogLevel.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSmbLogLevelFilter**
> SmbLogLevelFilters GetSmbLogLevelFilter(ctx, smbLogLevelFilterId)


View log filter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **smbLogLevelFilterId** | **string**| View log filter. | 

### Return type

[**SmbLogLevelFilters**](SmbLogLevelFilters.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSmbOpenfiles**
> SmbOpenfiles GetSmbOpenfiles(ctx, optional)


List open files.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string**| Order results by this field. Default is id. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **dir** | **string**| The direction of the sort. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 

### Return type

[**SmbOpenfiles**](SmbOpenfiles.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSmbSessions**
> SmbSessions GetSmbSessions(ctx, optional)


List open sessions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string**| Order results by this field. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **dir** | **string**| The direction of the sort. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 

### Return type

[**SmbSessions**](SmbSessions.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSmbSettingsGlobal**
> SmbSettingsGlobal GetSmbSettingsGlobal(ctx, optional)


List all settings.

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

[**SmbSettingsGlobal**](SmbSettingsGlobal.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSmbSettingsShare**
> SmbSettingsShare GetSmbSettingsShare(ctx, optional)


List all settings.

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
 **zone** | **string**| Zone which contains these share settings. | 

### Return type

[**SmbSettingsShare**](SmbSettingsShare.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSmbShare**
> SmbShares GetSmbShare(ctx, smbShareId, optional)


Retrieve share.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **smbShareId** | **string**| Retrieve share. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smbShareId** | **string**| Retrieve share. | 
 **scope** | **string**| If specified as \&quot;effective\&quot; or not specified, all fields are returned.  If specified as \&quot;user\&quot;, only fields with non-default values are shown.  If specified as \&quot;default\&quot;, the original values are returned. | 
 **resolveNames** | **bool**| If true, resolve group and user names in personas. | 
 **zone** | **string**| Zone which contains this share. | 

### Return type

[**SmbShares**](SmbShares.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSmbSharesSummary**
> SmbSharesSummary GetSmbSharesSummary(ctx, optional)


Return summary information about shares.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone** | **string**| Specifies which access zone or zones to use. | 

### Return type

[**SmbSharesSummary**](SmbSharesSummary.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSnmpSettings**
> SnmpSettings GetSnmpSettings(ctx, )


Retrieve the SNMP settings.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SnmpSettings**](SnmpSettings.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSwiftAccount**
> SwiftAccounts GetSwiftAccount(ctx, swiftAccountId, optional)


List a swift account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **swiftAccountId** | **string**| List a swift account. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **swiftAccountId** | **string**| List a swift account. | 
 **zone** | **string**| Access zone which contains Swift account. | 

### Return type

[**SwiftAccounts**](SwiftAccounts.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListHdfsProxyusers**
> HdfsProxyusers ListHdfsProxyusers(ctx, optional)


List all proxyusers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone** | **string**| Access zone which contains HDFS proxyusers. | 

### Return type

[**HdfsProxyusers**](HdfsProxyusers.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListHdfsRacks**
> HdfsRacksExtended ListHdfsRacks(ctx, optional)


List all racks.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone** | **string**| Access zone which contains HDFS racks. | 

### Return type

[**HdfsRacksExtended**](HdfsRacksExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNdmpSettingsPreferredIps**
> NdmpSettingsPreferredIps ListNdmpSettingsPreferredIps(ctx, optional)


Get list of preferences.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 

### Return type

[**NdmpSettingsPreferredIps**](NdmpSettingsPreferredIps.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNdmpUsers**
> NdmpUsersExtended ListNdmpUsers(ctx, )


List all ndmp administrators.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NdmpUsersExtended**](NdmpUsersExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNfsAliases**
> NfsAliasesExtended ListNfsAliases(ctx, optional)


List all NFS aliases.

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
 **zone** | **string**| Access zone | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **check** | **bool**| Check for conflicts when listing aliases. | 
 **dir** | **string**| The direction of the sort. | 

### Return type

[**NfsAliasesExtended**](NfsAliasesExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNfsExports**
> NfsExportsExtended ListNfsExports(ctx, optional)


List all NFS exports.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string**| The field that will be used for sorting. Default is id. | 
 **zone** | **string**| Access zone | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 
 **scope** | **string**| When specified as &#39;effective&#39;, or not specified, all fields are returned. When specified as &#39;user&#39;, only fields with non-default values are shown. When specified as &#39;default&#39;, the original values are returned. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **offset** | **int32**| The position of the first item returned for a paginated query within the full result set. | 
 **path** | **string**| If specified, only exports that explicitly reference at least one of the given paths will be returned. | 
 **check** | **bool**| Check for conflicts when listing exports. | 
 **dir** | **string**| The direction of the sort. | 

### Return type

[**NfsExportsExtended**](NfsExportsExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNtpServers**
> NtpServersExtended ListNtpServers(ctx, optional)


List all NTP servers.

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
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **dir** | **string**| The direction of the sort. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 

### Return type

[**NtpServersExtended**](NtpServersExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSmbLogLevelFilters**
> SmbLogLevelFilters ListSmbLogLevelFilters(ctx, optional)


Get the current SMB log filters.

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
 **dir** | **string**| The direction of the sort. | 
 **level** | **string**| Valid SMB logging levels | 

### Return type

[**SmbLogLevelFilters**](SmbLogLevelFilters.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSmbShares**
> SmbSharesExtended ListSmbShares(ctx, optional)


List all shares.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string**| Order results by this field. Default is id. | 
 **zone** | **string**| Zone which contains this share. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 
 **resolveNames** | **bool**| If true, resolve group and user names in personas. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **offset** | **int32**| The position of the first item returned for a paginated query within the full result set. | 
 **scope** | **string**| If specified as \&quot;effective\&quot; or not specified, all fields are returned.  If specified as \&quot;user\&quot;, only fields with non-default values are shown.  If specified as \&quot;default\&quot;, the original values are returned. | 
 **dir** | **string**| The direction of the sort. | 

### Return type

[**SmbSharesExtended**](SmbSharesExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSwiftAccounts**
> SwiftAccountsExtended ListSwiftAccounts(ctx, optional)


List all swift accounts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone** | **string**| Access zone which contains Swift accounts. | 

### Return type

[**SwiftAccountsExtended**](SwiftAccountsExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFtpSettings**
> UpdateFtpSettings(ctx, ftpSettings)


Modify the FTP settings. All input fields are optional, but one or more must be supplied.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ftpSettings** | [**FtpSettingsExtended**](FtpSettingsExtended.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateHdfsLogLevel**
> UpdateHdfsLogLevel(ctx, hdfsLogLevel)


Modify the HDFS service log-level.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hdfsLogLevel** | [**HdfsLogLevel**](HdfsLogLevel.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateHdfsProxyuser**
> UpdateHdfsProxyuser(ctx, hdfsProxyuser, hdfsProxyuserId, optional)


Modify an HDFS proxyuser.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hdfsProxyuser** | [**Empty**](Empty.md)|  | 
  **hdfsProxyuserId** | **string**| Modify an HDFS proxyuser. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hdfsProxyuser** | [**Empty**](Empty.md)|  | 
 **hdfsProxyuserId** | **string**| Modify an HDFS proxyuser. | 
 **zone** | **string**| Access zone which contains HDFS proxyuser. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateHdfsRack**
> UpdateHdfsRack(ctx, hdfsRack, hdfsRackId, optional)


Modify the HDFS rack

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hdfsRack** | [**HdfsRack**](HdfsRack.md)|  | 
  **hdfsRackId** | **string**| Modify the HDFS rack | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hdfsRack** | [**HdfsRack**](HdfsRack.md)|  | 
 **hdfsRackId** | **string**| Modify the HDFS rack | 
 **zone** | **string**| Access zone which contains HDFS rack. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateHdfsRangerPluginSettings**
> UpdateHdfsRangerPluginSettings(ctx, hdfsRangerPluginSettings, optional)


Modify HDFS ranger-plugin properties.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hdfsRangerPluginSettings** | [**HdfsRangerPluginSettingsSettings**](HdfsRangerPluginSettingsSettings.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hdfsRangerPluginSettings** | [**HdfsRangerPluginSettingsSettings**](HdfsRangerPluginSettingsSettings.md)|  | 
 **zone** | **string**| Access zone which contains HDFS ranger-plugin settings. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateHdfsSettings**
> UpdateHdfsSettings(ctx, hdfsSettings, optional)


Modify HDFS properties.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hdfsSettings** | [**HdfsSettingsSettings**](HdfsSettingsSettings.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hdfsSettings** | [**HdfsSettingsSettings**](HdfsSettingsSettings.md)|  | 
 **zone** | **string**| Access zone which contains HDFS settings. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateHttpSettings**
> UpdateHttpSettings(ctx, httpSettings)


Modify HTTP properties.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **httpSettings** | [**HttpSettingsSettings**](HttpSettingsSettings.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNdmpDiagnostics**
> UpdateNdmpDiagnostics(ctx, ndmpDiagnostics)


Modify ndmp diagnostics settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ndmpDiagnostics** | [**NdmpDiagnosticsDiagnostics**](NdmpDiagnosticsDiagnostics.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNdmpSettingsGlobal**
> UpdateNdmpSettingsGlobal(ctx, ndmpSettingsGlobal)


Modify one or more settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ndmpSettingsGlobal** | [**NdmpSettingsGlobalGlobal**](NdmpSettingsGlobalGlobal.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNdmpSettingsPreferredIp**
> UpdateNdmpSettingsPreferredIp(ctx, ndmpSettingsPreferredIp, ndmpSettingsPreferredIpId)


Modify a preferred ip preference.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ndmpSettingsPreferredIp** | [**NdmpSettingsPreferredIp**](NdmpSettingsPreferredIp.md)|  | 
  **ndmpSettingsPreferredIpId** | **string**| Modify a preferred ip preference. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNdmpSettingsVariable**
> UpdateNdmpSettingsVariable(ctx, ndmpSettingsVariable, ndmpSettingsVariableId, optional)


Modify or create a NDMP preferred environment variable.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ndmpSettingsVariable** | [**NdmpSettingsVariable**](NdmpSettingsVariable.md)|  | 
  **ndmpSettingsVariableId** | **string**| Modify or create a NDMP preferred environment variable. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ndmpSettingsVariable** | [**NdmpSettingsVariable**](NdmpSettingsVariable.md)|  | 
 **ndmpSettingsVariableId** | **string**| Modify or create a NDMP preferred environment variable. | 
 **name** | **string**| Name of the variable to modify. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNdmpUser**
> UpdateNdmpUser(ctx, ndmpUser, ndmpUserId)


Modify the user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ndmpUser** | [**NdmpUser**](NdmpUser.md)|  | 
  **ndmpUserId** | **string**| Modify the user | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNfsAlias**
> UpdateNfsAlias(ctx, nfsAlias, nfsAliasId, optional)


Modify the alias. All input fields are optional, but one or more must be supplied.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **nfsAlias** | [**NfsAlias**](NfsAlias.md)|  | 
  **nfsAliasId** | **string**| Modify the alias. All input fields are optional, but one or more must be supplied. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nfsAlias** | [**NfsAlias**](NfsAlias.md)|  | 
 **nfsAliasId** | **string**| Modify the alias. All input fields are optional, but one or more must be supplied. | 
 **zone** | **string**| Access zone | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNfsExport**
> UpdateNfsExport(ctx, nfsExport, nfsExportId, optional)


Modify the export. All input fields are optional, but one or more must be supplied.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **nfsExport** | [**NfsExport**](NfsExport.md)|  | 
  **nfsExportId** | **string**| Modify the export. All input fields are optional, but one or more must be supplied. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nfsExport** | [**NfsExport**](NfsExport.md)|  | 
 **nfsExportId** | **string**| Modify the export. All input fields are optional, but one or more must be supplied. | 
 **force** | **bool**| If true, the export will be updated even if that change conflicts with another export. | 
 **ignoreUnresolvableHosts** | **bool**| Ignore unresolvable hosts. | 
 **zone** | **string**| Access zone | 
 **ignoreConflicts** | **bool**| Ignore conflicts with existing exports. | 
 **ignoreBadPaths** | **bool**| Ignore nonexistent or otherwise bad paths. | 
 **ignoreBadAuth** | **bool**| Ignore invalid users. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNfsLogLevel**
> UpdateNfsLogLevel(ctx, nfsLogLevel)


Set the current NFS service logging level.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **nfsLogLevel** | [**NfsLogLevel**](NfsLogLevel.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNfsNetgroup**
> UpdateNfsNetgroup(ctx, nfsNetgroup, optional)


Modify the current NFS netgroup settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **nfsNetgroup** | [**NfsNetgroup**](NfsNetgroup.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nfsNetgroup** | [**NfsNetgroup**](NfsNetgroup.md)|  | 
 **host** | **string**| Host to retrieve netgroup cache settings for. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNfsSettingsExport**
> UpdateNfsSettingsExport(ctx, nfsSettingsExport, optional)


Modify the default values for NFS exports. All input fields are optional, but one or more must be supplied.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **nfsSettingsExport** | [**NfsSettingsExportSettings**](NfsSettingsExportSettings.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nfsSettingsExport** | [**NfsSettingsExportSettings**](NfsSettingsExportSettings.md)|  | 
 **zone** | **string**| Access zone | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNfsSettingsGlobal**
> UpdateNfsSettingsGlobal(ctx, nfsSettingsGlobal, optional)


Modify the default values for NFS exports. All input fields are optional, but one or more must be supplied.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **nfsSettingsGlobal** | [**NfsSettingsGlobalSettings**](NfsSettingsGlobalSettings.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nfsSettingsGlobal** | [**NfsSettingsGlobalSettings**](NfsSettingsGlobalSettings.md)|  | 
 **scope** | **string**| When specified as &#39;effective&#39;, or not specified, all fields are returned. When specified as &#39;user&#39;, only fields with non-default values are shown. When specified as &#39;default&#39;, the original values are returned. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNfsSettingsZone**
> UpdateNfsSettingsZone(ctx, nfsSettingsZone, optional)


Modify the NFS server settings for this zone.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **nfsSettingsZone** | [**NfsSettingsZoneSettings**](NfsSettingsZoneSettings.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nfsSettingsZone** | [**NfsSettingsZoneSettings**](NfsSettingsZoneSettings.md)|  | 
 **zone** | **string**| Access zone | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNtpServer**
> UpdateNtpServer(ctx, ntpServer, ntpServerId)


Modify the key value for an NTP server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ntpServer** | [**NtpServer**](NtpServer.md)|  | 
  **ntpServerId** | **string**| Modify the key value for an NTP server. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNtpSettings**
> UpdateNtpSettings(ctx, ntpSettings)


Modify the NTP settings. All input fields are optional, but one or more must be supplied.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ntpSettings** | [**NtpSettingsSettings**](NtpSettingsSettings.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSmbLogLevel**
> UpdateSmbLogLevel(ctx, smbLogLevel)


Set the current SMB logging level.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **smbLogLevel** | [**SmbLogLevel**](SmbLogLevel.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSmbSettingsGlobal**
> UpdateSmbSettingsGlobal(ctx, smbSettingsGlobal)


Modify one or more settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **smbSettingsGlobal** | [**SmbSettingsGlobalExtended**](SmbSettingsGlobalExtended.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSmbSettingsShare**
> UpdateSmbSettingsShare(ctx, smbSettingsShare, optional)


Modify one or more settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **smbSettingsShare** | [**SmbSettingsShareExtended**](SmbSettingsShareExtended.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smbSettingsShare** | [**SmbSettingsShareExtended**](SmbSettingsShareExtended.md)|  | 
 **zone** | **string**| Zone which contains these share settings. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSmbShare**
> UpdateSmbShare(ctx, smbShare, smbShareId, optional)


Modify share. All input fields are optional, but one or more must be supplied.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **smbShare** | [**SmbShare**](SmbShare.md)|  | 
  **smbShareId** | **string**| Modify share. All input fields are optional, but one or more must be supplied. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smbShare** | [**SmbShare**](SmbShare.md)|  | 
 **smbShareId** | **string**| Modify share. All input fields are optional, but one or more must be supplied. | 
 **zone** | **string**| Zone which contains this share. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSnmpSettings**
> UpdateSnmpSettings(ctx, snmpSettings)


Modify the SNMP settings. All input fields are optional, but one or more must be supplied.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **snmpSettings** | [**SnmpSettingsExtended**](SnmpSettingsExtended.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSwiftAccount**
> UpdateSwiftAccount(ctx, swiftAccount, swiftAccountId, optional)


Modify a Swift account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **swiftAccount** | [**SwiftAccount**](SwiftAccount.md)|  | 
  **swiftAccountId** | **string**| Modify a Swift account | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **swiftAccount** | [**SwiftAccount**](SwiftAccount.md)|  | 
 **swiftAccountId** | **string**| Modify a Swift account | 
 **zone** | **string**| Access zone which contains Swift account. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

