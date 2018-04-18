# SummaryClientClientItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | **string** | The class of the operation. | [default to null]
**In** | **float32** | Rate of input (in bytes/second) for an operation since the last time isi statistics collected the data. | [default to null]
**InAvg** | **float32** | Average input (received) bytes for an operation, in bytes. | [default to null]
**InMax** | **float32** | Maximum input (received) bytes for an operation, in bytes. | [default to null]
**InMin** | **float32** | Minimum input (received) bytes for an operation, in bytes. | [default to null]
**LocalAddr** | **string** | The IP address (in dotted-quad form) of the host receiving the operation request. | [default to null]
**LocalName** | **string** | The resolved text name of the LocalAddr, if resolution can be performed. | [default to null]
**Node** | **int32** | The node on which the operation was performed. | [optional] [default to null]
**NumOperations** | **int32** | The number of times an operation has been performed. | [default to null]
**OperationRate** | **float32** | The rate (in ops/second) at which an operation has been performed. | [default to null]
**Out** | **float32** | Rate of output (in bytes/second) for an operation since the last time isi statistics collected the data. | [default to null]
**OutAvg** | **float32** | Average output (sent) bytes for an operation, in bytes. | [default to null]
**OutMax** | **float32** | Maximum output (sent) bytes for an operation, in bytes. | [default to null]
**OutMin** | **float32** | Minimum output (sent) bytes for an operation, in bytes. | [default to null]
**Protocol** | **string** | The protocol of the operation. | [default to null]
**RemoteAddr** | **string** | The IP address (in dotted-quad form) of the host sending the operation request. | [default to null]
**RemoteName** | **string** | The resolved text name of the RemoteAddr, if resolution can be performed. | [default to null]
**Time** | **int32** | Unix Epoch time in seconds of the request. | [default to null]
**TimeAvg** | **float32** | The average elapsed time (in microseconds) taken to complete an operation. | [default to null]
**TimeMax** | **float32** | The maximum elapsed time (in microseconds) taken to complete an operation. | [default to null]
**TimeMin** | **float32** | The minimum elapsed time (in microseconds) taken to complete an operation. | [default to null]
**User** | [***AuthAccessAccessItemFileGroup**](AuthAccessAccessItemFileGroup.md) | Specifies properties for a persona, which consists of either a &#39;type&#39; and a &#39;name&#39; or an &#39;ID&#39;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


