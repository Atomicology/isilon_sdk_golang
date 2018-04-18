# ProvidersKrb5Krb5ItemExtended

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groupnet** | **string** | Groupnet identifier. | [optional] [default to null]
**Id** | **string** | Specifies the Kerberos provider ID. | [optional] [default to null]
**KeytabEntries** | [**[]ProvidersKrb5IdParamsKeytabEntry**](ProvidersKrb5IdParamsKeytabEntry.md) | Specifies the key information for the Kerberos SPNs. | [optional] [default to null]
**KeytabFile** | **string** | Specifies the path to a keytab file to import. | [optional] [default to null]
**ManualKeying** | **bool** | If true, keys are managed manually. If false, keys are managed through kadmin. | [optional] [default to null]
**Name** | **string** | Specifies the Kerberos provider name. | [optional] [default to null]
**Realm** | **string** | Specifies the name of realm. | [optional] [default to null]
**RecommendedSpns** | **[]string** | Specifies the recommended SPNs. | [optional] [default to null]
**Status** | **string** | Specifies the status of the provider. | [optional] [default to null]
**System** | **bool** | If true, indicates that this provider instance was created by OneFS and cannot be removed | [optional] [default to null]
**User** | **string** | Specifies the name of the user that performs kadmin tasks. | [optional] [default to null]
**Password** | **string** | Specifies the Kerberos provider password. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


