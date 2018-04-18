# NodeStatusNode

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Batterystatus** | [***NodeStatusNodeBatterystatus**](NodeStatusNodeBatterystatus.md) | Battery status information. | [optional] [default to null]
**Capacity** | [**[]NodeStatusNodeCapacityItem**](NodeStatusNodeCapacityItem.md) | Storage capacity of this node. | [optional] [default to null]
**Cpu** | [***NodeStatusNodeCpu**](NodeStatusNodeCpu.md) | CPU status information for this node. | [optional] [default to null]
**Error_** | **string** | Error message, if the HTTP status returned from this node was not 200. | [optional] [default to null]
**Id** | **int32** | Node ID of the node reporting this information. | [optional] [default to null]
**Lnn** | **int32** | Logical node number of the node reporting this information. | [optional] [default to null]
**Nvram** | [***NodeStatusNodeNvram**](NodeStatusNodeNvram.md) | Node NVRAM information. | [optional] [default to null]
**Powersupplies** | [***NodeStatusNodePowersupplies**](NodeStatusNodePowersupplies.md) | Information about this node&#39;s power supplies. | [optional] [default to null]
**Release** | **string** | OneFS release. | [optional] [default to null]
**Status** | **int32** | Status of the HTTP response from this node if not 200.  If 200, this field does not appear. | [optional] [default to null]
**Uptime** | **int32** | Seconds this node has been online. | [optional] [default to null]
**Version** | **string** | OneFS version. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


