# NdmpLogsNode

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | **string** | The page on this node&#39;s NDMP log file being returned. | [optional] [default to null]
**Error_** | **string** | Error message, if the HTTP status returned from this node was not 200. | [optional] [default to null]
**Id** | **int32** | Node ID of the node reporting this information. | [optional] [default to null]
**Lnn** | **int32** | Logical node number of the node reporting this information. | [optional] [default to null]
**Logs** | **string** | The log file entries from the current page on this node. | [optional] [default to null]
**MaxPage** | **int32** | The highest page number in this node&#39;s NDMP log file. | [optional] [default to null]
**Status** | **int32** | Status of the HTTP response from this node if not 200.  If 200, this field does not appear. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


