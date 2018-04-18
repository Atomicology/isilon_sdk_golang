# NfsNlmSessionsSession

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delegates** | **[]int32** |  | [optional] [default to null]
**HostType** | **string** | The sort of host that this entry represents | [optional] [default to null]
**Hostname** | **string** | The host being monitored | [optional] [default to null]
**IsActive** | **bool** | Whether or not the client is actively being monitored | [optional] [default to null]
**LastModified** | **int32** | Unix time in seconds that the client was last modified (monitored or unmonitored) | [optional] [default to null]
**NodeIp** | **string** | An IP address for which NSM has client records | [optional] [default to null]
**NotifyAttemptsRemaining** | **int32** | Number of times we will attempt to notify this client before giving up | [optional] [default to null]
**NotifyError** | **string** | Last error recieved attempting to notify this client | [optional] [default to null]
**NotifyLastAttempt** | **int32** | Unix time in seconds when we last attempted to notify this clients | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


