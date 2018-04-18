# NodeStateNode

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error_** | **string** | Error message, if the HTTP status returned from this node was not 200. | [optional] [default to null]
**Id** | **int32** | Node ID of the node reporting this information. | [optional] [default to null]
**Lnn** | **int32** | Logical node number of the node reporting this information. | [optional] [default to null]
**Readonly** | [***NodeStateReadonlyExtended**](NodeStateReadonlyExtended.md) | Node readonly state. | [optional] [default to null]
**Servicelight** | [***NodeStateNodeServicelight**](NodeStateNodeServicelight.md) | Node service light state. | [optional] [default to null]
**Smartfail** | [***NodeStateSmartfailExtended**](NodeStateSmartfailExtended.md) | Node smartfail state. | [optional] [default to null]
**Status** | **int32** | Status of the HTTP response from this node if not 200.  If 200, this field does not appear. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


