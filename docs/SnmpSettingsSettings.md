# SnmpSettingsSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReadOnlyCommunity** | **string** | The read-only community name.  @DEFAULT reverts this field to its default value. | [optional] [default to null]
**Service** | **bool** | Whether the SNMP service is enabled. | [optional] [default to null]
**SnmpV1V2cAccess** | **bool** | Whether SNMP v1 and v2c protocols are enabled.  @DEFAULT reverts this field to its default value. | [optional] [default to null]
**SnmpV3Access** | **bool** | Whether SNMP v3 is enabled.  @DEFAULT reverts this field to its default value. | [optional] [default to null]
**SnmpV3AuthProtocol** | **string** | SNMPv3 authentication protocol. May only be SHA or MD5.  @DEFAULT reverts this field to its default value. | [optional] [default to null]
**SnmpV3PrivProtocol** | **string** | SNMPv3 privacy protocol. May only be AES or DES. @DEFAULT reverts this field to its default value. | [optional] [default to null]
**SnmpV3ReadOnlyUser** | **string** | The read-only user for SNMP v3 read requests.  @DEFAULT reverts this field to its default value. | [optional] [default to null]
**SnmpV3SecurityLevel** | **string** | SNMPv3 privacy protocol. May only be AES or DES. @DEFAULT reverts this field to its default value. | [optional] [default to null]
**SystemContact** | **string** | Contact information for the system owner.  This must be a valid email address.  @DEFAULT reverts this field to its default value. | [optional] [default to null]
**SystemLocation** | **string** | A location name for the SNMP system.  @DEFAULT reverts this field to its default value. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


