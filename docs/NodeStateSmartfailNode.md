# NodeStateSmartfailNode

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dead** | **bool** | This node is dead (dead_devs). | [optional] [default to null]
**Down** | **bool** | This node is down (down_devs). | [optional] [default to null]
**Error_** | **string** | Error message, if the HTTP status returned from this node was not 200. | [optional] [default to null]
**Id** | **int32** | Node ID of the node reporting this information. | [optional] [default to null]
**InCluster** | **bool** | This node is in the cluster (all_devs). | [optional] [default to null]
**Lnn** | **int32** | Logical node number of the node reporting this information. | [optional] [default to null]
**Readonly** | **bool** | This node is readonly (ro_devs). | [optional] [default to null]
**ShutdownReadonly** | **bool** | This node is shutdown readonly (down_devs). | [optional] [default to null]
**Smartfailed** | **bool** | This node is smartfailed (soft_devs). | [optional] [default to null]
**Status** | **int32** | Status of the HTTP response from this node if not 200.  If 200, this field does not appear. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


