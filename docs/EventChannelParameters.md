# EventChannelParameters

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **[]string** | Email addresses to send to. | [optional] [default to null]
**Batch** | **string** | Batching criterion. | [optional] [default to null]
**BatchPeriod** | **int32** | Period over which batching is to be performed. | [optional] [default to null]
**CustomTemplate** | **string** | Path to custom notification template. | [optional] [default to null]
**SendAs** | **string** | Email address to use as from. | [optional] [default to null]
**SmtpHost** | **string** | SMTP relay host. | [optional] [default to null]
**SmtpPassword** | **string** | Password for SMTP authentication - only if smtp_use_auth true. | [optional] [default to null]
**SmtpPort** | **int32** | SMTP relay port - optional defaults to 25. | [optional] [default to null]
**SmtpSecurity** | **string** | Encryption protocol to use for SMTP. | [optional] [default to null]
**SmtpUseAuth** | **bool** | Use SMTP authentication - optional defaulst to false. | [optional] [default to null]
**SmtpUsername** | **string** | Username for SMTP authentication - only if smtp_use_auth true. | [optional] [default to null]
**Subject** | **string** | Subject for emails. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


