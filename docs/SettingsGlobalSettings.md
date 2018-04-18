# SettingsGlobalSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditedZones** | **[]string** | Specifies zones that are audited when the protocol_auditing_enabled property is enabled. | [optional] [default to null]
**CeeLogTime** | **string** | Specifies that events past a certain date are forwarded by the audit CEE forwarder. Format these events as follows: &#39;Topic@YYYY-MM-DD HH:MM:SS&#39;. | [optional] [default to null]
**CeeServerUris** | **[]string** | Specifies a list of Common Event Enabler (CEE) server URIs. Protocol audit logs are sent to these URIs for external processing. | [optional] [default to null]
**ConfigAuditingEnabled** | **bool** | Specifies whether logging for API configuration changes are enabled. | [optional] [default to null]
**ConfigSyslogEnabled** | **bool** | Specifies whether configuration audit syslog messages are forwarded. | [optional] [default to null]
**Hostname** | **string** | Specifies the hostname that is reported in protocol events from this cluster. | [optional] [default to null]
**ProtocolAuditingEnabled** | **bool** | Specifies if logging for the I/O stream is enabled. | [optional] [default to null]
**SyslogLogTime** | **string** | Specifies that events past a specified date are forwarded by the audit syslog forwarder. Format these events as follows: &#39;Topic@YYYY-MM-DD HH:MM:SS&#39; format | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


