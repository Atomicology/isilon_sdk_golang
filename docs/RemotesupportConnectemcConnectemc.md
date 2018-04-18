# RemotesupportConnectemcConnectemc

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailCustomerOnFailure** | **bool** | Email the customer if all transmission methods fail. | [optional] [default to null]
**Enabled** | **bool** | Enable ConnectEMC. | [optional] [default to null]
**GatewayAccessPools** | **[]string** | List of network pools that are able to connect to the ESRS gateway.  Necessary to enable ConnectEMC. | [optional] [default to null]
**PrimaryEsrsGateway** | **string** | Primary ESRS Gateway. Necessary to enable ConnectEMC. | [optional] [default to null]
**SecondaryEsrsGateway** | **string** | Secondary ESRS Gateway. Used if Primary is unavailable. | [optional] [default to null]
**UseSmtpFailover** | **bool** | Use SMPT if primary and secondary gateways are unavailable. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


