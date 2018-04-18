# \HardwareApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHardwareTapeName**](HardwareApi.md#CreateHardwareTapeName) | **Post** /platform/3/hardware/tape/{HardwareTapeName} | 
[**DeleteHardwareTapeName**](HardwareApi.md#DeleteHardwareTapeName) | **Delete** /platform/3/hardware/tape/{HardwareTapeName} | 
[**GetHardwareFcport**](HardwareApi.md#GetHardwareFcport) | **Get** /platform/3/hardware/fcports/{HardwareFcportId} | 
[**GetHardwareFcports**](HardwareApi.md#GetHardwareFcports) | **Get** /platform/3/hardware/fcports | 
[**GetHardwareTapes**](HardwareApi.md#GetHardwareTapes) | **Get** /platform/3/hardware/tapes | 
[**UpdateHardwareFcport**](HardwareApi.md#UpdateHardwareFcport) | **Put** /platform/3/hardware/fcports/{HardwareFcportId} | 
[**UpdateHardwareTapeName**](HardwareApi.md#UpdateHardwareTapeName) | **Put** /platform/3/hardware/tape/{HardwareTapeName} | 


# **CreateHardwareTapeName**
> CreateHardwareTapeNameResponse CreateHardwareTapeName(ctx, hardwareTapeName, hardwareTapeName2, optional)


Tape/Changer devices rescan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hardwareTapeName** | [**Empty**](Empty.md)|  | 
  **hardwareTapeName2** | **string**| Tape/Changer devices rescan | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hardwareTapeName** | [**Empty**](Empty.md)|  | 
 **hardwareTapeName2** | **string**| Tape/Changer devices rescan | 
 **lnn** | **string**| Logical node number. | 
 **port** | **int32**| Scan only specified port. | 
 **timeout** | **float32**| Timeout for request. | 
 **reconcile** | **bool**| Remove entries for devices or paths that have become inaccessible. | 

### Return type

[**CreateHardwareTapeNameResponse**](CreateHardwareTapeNameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHardwareTapeName**
> DeleteHardwareTapeName(ctx, hardwareTapeName)


Tape/Changer devices remove

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hardwareTapeName** | **string**| Tape/Changer devices remove | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHardwareFcport**
> HardwareFcports GetHardwareFcport(ctx, hardwareFcportId, optional)


Get one fibre-channel port

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hardwareFcportId** | **int32**| Get one fibre-channel port | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hardwareFcportId** | **int32**| Get one fibre-channel port | 
 **lnn** | **string**| Logical node number. | 

### Return type

[**HardwareFcports**](HardwareFcports.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHardwareFcports**
> HardwareFcports GetHardwareFcports(ctx, optional)


Get list of fibre-channel ports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lnn** | **string**| Logical node number. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 

### Return type

[**HardwareFcports**](HardwareFcports.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHardwareTapes**
> HardwareTapes GetHardwareTapes(ctx, optional)


Get list Tape and Changer devices

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **node** | **string**| List only devices on the specified node. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 
 **devname** | **string**| List only the named device. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **activepath** | **string**| List devices with only active paths. | 
 **type_** | **string**| Restrict to list on tape or mc device. | 

### Return type

[**HardwareTapes**](HardwareTapes.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateHardwareFcport**
> UpdateHardwareFcport(ctx, hardwareFcport, hardwareFcportId, optional)


Change wwnn, wwpn, state, topology, or rate of a FC port

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hardwareFcport** | [**HardwareFcport**](HardwareFcport.md)|  | 
  **hardwareFcportId** | **int32**| Change wwnn, wwpn, state, topology, or rate of a FC port | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hardwareFcport** | [**HardwareFcport**](HardwareFcport.md)|  | 
 **hardwareFcportId** | **int32**| Change wwnn, wwpn, state, topology, or rate of a FC port | 
 **lnn** | **string**| Logical node number. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateHardwareTapeName**
> UpdateHardwareTapeName(ctx, hardwareTapeNameParams, hardwareTapeName)


Tape/Changer device modify

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hardwareTapeNameParams** | [**HardwareTapeNameParams**](HardwareTapeNameParams.md)|  | 
  **hardwareTapeName** | **string**| Tape/Changer device modify | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

