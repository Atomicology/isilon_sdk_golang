# NodeStatusBatterystatusNode

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error_** | **string** | Error message, if the HTTP status returned from this node was not 200. | [optional] [default to null]
**Id** | **int32** | Node ID of the node reporting this information. | [optional] [default to null]
**LastTestTime1** | **string** | The last battery test time for battery 1. | [optional] [default to null]
**LastTestTime2** | **string** | The last battery test time for battery 2. | [optional] [default to null]
**Lnn** | **int32** | Logical node number of the node reporting this information. | [optional] [default to null]
**NextTestTime1** | **string** | The next checkup for battery 1. | [optional] [default to null]
**NextTestTime2** | **string** | The next checkup for battery 2. | [optional] [default to null]
**Present** | **bool** | Node has battery status. | [optional] [default to null]
**Result1** | **string** | The result of the last battery test for battery 1. | [optional] [default to null]
**Result2** | **string** | The result of the last battery test for battery 2. | [optional] [default to null]
**Status** | **int32** | Status of the HTTP response from this node if not 200.  If 200, this field does not appear. | [optional] [default to null]
**Status1** | **string** | The status of battery 1. | [optional] [default to null]
**Status2** | **string** | The status of battery 2. | [optional] [default to null]
**Supported** | **bool** | Node supports battery status. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


