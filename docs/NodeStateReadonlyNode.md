# NodeStateReadonlyNode

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowed** | **bool** | The current read-only mode allowed status for the node. | [optional] [default to null]
**Enabled** | **bool** | The current read-only user mode status for the node. NOTE: If read-only mode is currently disallowed for this node, it will remain read/write until read-only mode is allowed again. This value only sets or clears any user-specified requests for read-only mode. If the node has been placed into read-only mode by the system, it will remain in read-only mode until the system conditions which triggered read-only mode have cleared. | [optional] [default to null]
**Error_** | **string** | Error message, if the HTTP status returned from this node was not 200. | [optional] [default to null]
**Id** | **int32** | Node ID of the node reporting this information. | [optional] [default to null]
**Lnn** | **int32** | Logical node number of the node reporting this information. | [optional] [default to null]
**Mode** | **bool** | The current read-only mode status for the node. | [optional] [default to null]
**Status** | **string** | The current read-only mode status description for the node. | [optional] [default to null]
**Valid** | **bool** | The read-only state values are valid (False &#x3D; Error). | [optional] [default to null]
**Value** | **int32** | The current read-only value (enumerated bitfield) for the node. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


